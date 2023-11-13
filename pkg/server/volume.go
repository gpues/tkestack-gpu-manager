/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package server

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"tkestack.io/gpu-manager/pkg/services/volume"
	"tkestack.io/gpu-manager/pkg/types"

	"k8s.io/klog"
)

// VolumeManager manages volumes used by containers running GPU application
type VolumeManager struct {
	Config          []Config `json:"volume,omitempty"`
	cfgPath         string
	cudaControlFile string
	CudaSoname      map[string]string `json:"cudaSoname"`
	MlSoName        map[string]string `json:"mlSoName"`
	Share           bool              `json:"share"`
}

type components map[string][]string

// Config contains volume details in config file
type Config struct {
	Name       string     `json:"name,omitempty"`
	Mode       string     `json:"mode,omitempty"`
	Components components `json:"components,omitempty"`
	BasePath   string     `json:"base,omitempty"`
}

const (
	binDir   = "bin"
	lib32Dir = "lib"
	lib64Dir = "lib64"
)

type volumeDir struct {
	name  string
	files []string
}

// Volume contains directory and file info of volume
type Volume struct {
	Path string
	dirs []volumeDir
}

// VolumeMap stores Volume for each type
type VolumeMap map[string]*Volume

// Run starts a VolumeManager
func (vm *VolumeManager) Run() (err error) {
	file, err := os.ReadFile(types.FILE)
	if err != nil {
		klog.Fatalln(err)
	}
	_ = json.Unmarshal(file, &vm)

	for _, cfg := range vm.Config {
		if cfg.Name == "nvidia" {
			types.DriverLibraryPath = filepath.Join(cfg.BasePath, cfg.Name)
		} else {
			types.DriverOriginLibraryPath = filepath.Join(cfg.BasePath, cfg.Name)
		}
	}
	return nil
}

func (vm *VolumeManager) Copy() error {

	vols := make(VolumeMap)
	for _, cfg := range vm.Config {
		vol := &Volume{
			Path: path.Join(cfg.BasePath, cfg.Name),
		}
		for t, c := range cfg.Components {
			switch t {
			case "binaries":
				bins, err := volume.Which(c...)
				if err != nil {
					return err
				}

				klog.Infof("Find binaries: %+v", bins)

				vol.dirs = append(vol.dirs, volumeDir{binDir, bins})
			case "libraries":
				var libs32 []string
				var libs64 []string
				filepath.WalkDir(filepath.Join(types.NvDir, "lib"), func(path string, d fs.DirEntry, err error) error {
					if path != filepath.Join(types.NvDir, "lib") {
						libs32 = append(libs32, path)
					}
					return nil
				})
				filepath.WalkDir(filepath.Join(types.NvDir, "lib64"), func(path string, d fs.DirEntry, err error) error {
					if path != filepath.Join(types.NvDir, "lib64") {
						libs64 = append(libs64, path)
					}
					return nil
				})
				klog.V(2).Infof("Find 32bit libraries: %+v", libs32)
				klog.V(2).Infof("Find 64bit libraries: %+v", libs64)

				vol.dirs = append(vol.dirs, volumeDir{lib32Dir, libs32}, volumeDir{lib64Dir, libs64})
			}

			vols[cfg.Name] = vol
		}
	}

	if err := vm.mirror(vols); err != nil {
		return err
	}

	klog.V(2).Infof("Volume manager is running")
	return nil
}

// #lizard forgives
func (vm *VolumeManager) mirror(vols VolumeMap) error {
	for _, vol := range vols {
		if exist, _ := vol.exist(); !exist {
			if err := os.MkdirAll(vol.Path, 0755); err != nil {
				return err
			}
		}

		for _, d := range vol.dirs {
			vpath := path.Join(vol.Path, d.name)
			if err := os.MkdirAll(vpath, 0755); err != nil {
				return err
			}

			for _, f := range d.files {
				klog.V(2).Infof("Mirror %s to %s", f, vpath)
				if strings.HasPrefix(path.Base(f), "libvgpu.so") {
					vm.cudaControlFile = f
				}
			}
			for _, f := range d.files {
				if strings.HasPrefix(path.Base(f), "libcuda.so.") {
					driverStr := strings.SplitN(strings.TrimPrefix(path.Base(f), "libcuda.so."), ".", -1)
					if len(driverStr) < 2 {
						continue
					}
					types.DriverVersionMajor, _ = strconv.Atoi(driverStr[0])
					types.DriverVersionMinor, _ = strconv.Atoi(driverStr[1])
					klog.Infof("Driver version: %d.%d", types.DriverVersionMajor, types.DriverVersionMinor)
					break
				}
			}
		}
	}

	vCudaFileFn := func(soFile string) error {
		os.Chmod(vm.cudaControlFile, os.ModePerm)
		l := filepath.Join(filepath.Dir(vm.cudaControlFile), soFile)
		os.Remove(l)
		GenLink(filepath.Base(vm.cudaControlFile), soFile, filepath.Dir(vm.cudaControlFile))
		klog.Infof("Vcuda %s to %s", vm.cudaControlFile, l)
		return nil
	}
	klog.Infoln("cudaControlFile ", vm.cudaControlFile)
	klog.Infoln("Share ", vm.Share)
	klog.Infoln("CudaSoname ", vm.CudaSoname)
	if vm.Share && len(vm.cudaControlFile) > 0 {
		if len(vm.CudaSoname) > 0 {
			for _, f := range vm.CudaSoname {
				if err := vCudaFileFn(f); err != nil {
					klog.Errorln(err)
					return err
				}
			}
		}

		if len(vm.MlSoName) > 0 {
			for _, f := range vm.MlSoName {
				if err := vCudaFileFn(f); err != nil {
					klog.Errorln(err)
					return err
				}
			}
		}
	}

	return nil
}

func (v *Volume) exist() (bool, error) {
	_, err := os.Stat(v.Path)
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func (v *Volume) remove() error {
	return os.RemoveAll(v.Path)
}
func GenLink(source, target, pwd string) {
	klog.Infoln("gen link", source, "->", target)
	cmd := exec.Command("ln", "-s", source, target) // 你可以替换成你想要执行的命令和参数
	cmd.Dir = pwd
	var out bytes.Buffer
	cmd.Stdout = &out
	klog.Infoln(source, target, pwd, out.String())
	err := cmd.Run()
	if err != nil {
		klog.Fatalln(err, out.String())
	}
}
