package runtime

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	v1 "k8s.io/api/core/v1"
	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/qos"

	"tkestack.io/gpu-manager/pkg/services/watchdog"
	"tkestack.io/gpu-manager/pkg/types"
	"tkestack.io/gpu-manager/pkg/utils"
	"tkestack.io/gpu-manager/pkg/utils/cgroup"
)

type ContainerRuntimeInterface interface {
	// Get pids in the given container id
	GetPidsInContainers(containerID string) ([]int, error)
	// InspectContainer returns the container information by the given name
	InspectContainer(containerID string) (*criapi.ContainerStatus, error)
	// RuntimeName returns the container runtime name
	RuntimeName() string
}

type containerRuntimeManager struct {
	cgroupDriver   string
	runtimeName    string
	requestTimeout time.Duration
	client         criapi.RuntimeServiceClient
}

var _ ContainerRuntimeInterface = (*containerRuntimeManager)(nil)

var (
	containerRoot = cgroup.NewCgroupName([]string{}, "kubepods.slice")
)

func (m *containerRuntimeManager) GetPidsInContainers(containerID string) ([]int, error) {
	req := &criapi.ContainerStatusRequest{
		ContainerId: containerID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.requestTimeout)
	defer cancel()

	resp, err := m.client.ContainerStatus(ctx, req)
	if err != nil {
		klog.Errorf("can't get container %s status, %v", containerID, err)
		return nil, err
	}

	ns := resp.Status.Labels[types.PodNamespaceLabelKey]
	podName := resp.Status.Labels[types.PodNameLabelKey]

	pod, err := watchdog.GetPod(ns, podName)
	if err != nil {
		klog.Errorf("can't get pod %s/%s, %v", ns, podName, err)
		return nil, err
	}

	cgroupPath, err := m.getCGroupName(pod, containerID)
	if err != nil {
		klog.Errorf("can't get cgroup parent, %v", err)
		return nil, err
	}

	pids := make([]int, 0)

	baseDir := filepath.Clean(filepath.Join(types.CGROUP_BASE, cgroupPath))

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() || info.Name() != types.CGROUP_PROCS {
			return nil
		}

		p, err := readProcsFile(path)
		if err == nil {
			pids = append(pids, p...)
		}

		return nil
	})

	return pids, nil
}

func readProcsFile(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		klog.Errorf("can't read %s, %v", file, err)
		return nil, nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pids := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if pid, err := strconv.Atoi(line); err == nil {
			pids = append(pids, pid)
		}
	}

	klog.V(4).Infof("Read from %s, pids: %v", file, pids)
	return pids, nil
}

func (m *containerRuntimeManager) getCGroupName(pod *v1.Pod, containerID string) (cGroupPath string, err error) {
	defer klog.Infoln(cGroupPath)
	podQos := pod.Status.QOSClass
	if len(podQos) == 0 {
		podQos = qos.GetPodQOS(pod)
	}
	PodCGroupNamePrefix := "kubepods-pod"
	var parentContainer cgroup.CgroupName
	switch podQos {
	case v1.PodQOSGuaranteed:
		parentContainer = cgroup.NewCgroupName(containerRoot)
	case v1.PodQOSBurstable:
		parentContainer = cgroup.NewCgroupName(containerRoot, fmt.Sprintf("kubepods-%s.slice", strings.ToLower(string(v1.PodQOSBurstable))))
		PodCGroupNamePrefix = fmt.Sprintf("kubepods-%s-pod", strings.ToLower(string(v1.PodQOSBurstable)))
	case v1.PodQOSBestEffort:
		parentContainer = cgroup.NewCgroupName(containerRoot, fmt.Sprintf("kubepods-%s.slice", strings.ToLower(string(v1.PodQOSBestEffort))))
		PodCGroupNamePrefix = fmt.Sprintf("kubepods-%s-pod", strings.ToLower(string(v1.PodQOSBestEffort)))
	}

	podContainer := PodCGroupNamePrefix + string(pod.UID)
	cGroupName := cgroup.NewCgroupName(parentContainer, podContainer)

	switch m.cgroupDriver {
	case "systemd":
		cGroupPath = fmt.Sprintf("%s/%s-%s.scope", cGroupName.ToSystemd(), "docker", containerID)
	case "cgroupfs":
		cGroupPath = fmt.Sprintf("%s/%s", cGroupName.ToCgroupfs(), containerID)
	default:
		err = fmt.Errorf("unsupported cgroup driver")
	}
	return
}

func (m *containerRuntimeManager) RuntimeName() string { return m.runtimeName }

func (m *containerRuntimeManager) InspectContainer(containerID string) (*criapi.ContainerStatus, error) {
	req := &criapi.ContainerStatusRequest{
		ContainerId: containerID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.requestTimeout)
	defer cancel()

	resp, err := m.client.ContainerStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Status, nil
}

func NewContainerRuntimeManager(cgroupDriver, endpoint string, requestTimeout time.Duration) (*containerRuntimeManager, error) {
	dialOptions := []grpc.DialOption{grpc.WithInsecure(), grpc.WithDialer(utils.UnixDial), grpc.WithBlock(), grpc.WithTimeout(time.Second * 5)}
	conn, err := grpc.Dial(endpoint, dialOptions...)
	if err != nil {
		return nil, err
	}

	client := criapi.NewRuntimeServiceClient(conn)

	m := &containerRuntimeManager{
		cgroupDriver:   cgroupDriver,
		client:         client,
		requestTimeout: requestTimeout,
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.requestTimeout)
	defer cancel()
	resp, err := client.Version(ctx, &criapi.VersionRequest{})
	if err != nil {
		return nil, err
	}

	klog.V(2).Infof("Container runtime info %s", resp)
	m.runtimeName = resp.RuntimeName

	return m, nil
}
