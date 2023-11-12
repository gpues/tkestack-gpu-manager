package main

import (
	"bytes"
	"encoding/json"
	"io"
	"k8s.io/klog"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"tkestack.io/gpu-manager/pkg/services/volume"
)

func ReadReallyFile(filePath string) string {
	symlinks, e1 := filepath.EvalSymlinks(filePath)
	if e1 == nil {
		return symlinks
	}
	symlinks, e2 := os.Readlink(filePath)
	if e2 == nil {
		return symlinks
	}
	lstat, e3 := os.Lstat(filePath)
	if e3 == nil {
		return filepath.Join(filepath.Dir(filePath), lstat.Name())
	}
	return filePath
}

const (
	FILE       = "/etc/gpu-manager/volume.json"
	NvDir      = "/etc/gpu-manager/vdriver/nvidia"
	FindBase   = "/usr/local/gpu/"
	controlLib = "/usr/local/gpu/libvgpu.so"
)

func main() {
	_ = os.RemoveAll(NvDir)
	_ = os.MkdirAll(NvDir, os.ModeDir)
	_ = os.MkdirAll(filepath.Join(NvDir, "lib"), os.ModeDir)
	_ = os.MkdirAll(filepath.Join(NvDir, "lib64"), os.ModeDir)
	_ = os.MkdirAll(filepath.Join(NvDir, "bin"), os.ModeDir)
	copyFileWithModeAndOwnership(controlLib, filepath.Join(NvDir, "lib64", filepath.Base(controlLib)))
	file, err := os.ReadFile(FILE)
	if err != nil {
		klog.Fatalln(err)
	}

	vs := volume.VolumeManager{}
	_ = json.Unmarshal(file, &vs)

	for _, v := range vs.Config {
		if v.Name != "nvidia" {
			continue
		}
		for _, lib := range v.Components["libraries"] {
			reallyFiles := SearchFile(FindBase, lib+"*", "stubs")
			for linkFile, reallyFile := range reallyFiles {
				if linkFile != reallyFile {
					continue
				}
				arch := GetArchFromPath(reallyFile)
				err := copyFileWithModeAndOwnership(reallyFile, filepath.Join(filepath.Join(NvDir, "lib")+arch, filepath.Base(linkFile)))
				if err != nil {
					klog.Fatalln(err)
				}
			}
			for linkFile, reallyFile := range reallyFiles { // 链接文件
				if linkFile == reallyFile {
					continue
				}
				arch := GetArchFromPath(reallyFile)
				//reallyPath := filepath.Join(filepath.Join(NvDir, "lib")+arch, filepath.Base(reallyFile))
				targetLinkPath := filepath.Join(filepath.Join(NvDir, "lib")+arch, filepath.Base(linkFile))
				_ = os.Remove(targetLinkPath)
				GenLink(filepath.Base(reallyFile), filepath.Base(linkFile), filepath.Join(NvDir, "lib")+arch)
			}
		}
		for _, bin := range v.Components["binaries"] {
			reallyFiles := SearchFile(FindBase, bin, "")
			for linkFile, reallyFile := range reallyFiles {
				err := copyFileWithModeAndOwnership(reallyFile, filepath.Join(NvDir, "bin", filepath.Base(linkFile)))
				if err != nil {
					klog.Fatalln(err)
				}
			}
		}
	}
}

func SearchFile(root string, searchPattern string, skipPattern string) map[string]string {
	reallyFiles := map[string]string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePath := path
			// 检查文件名是否符合规则
			if match, _ := filepath.Match(searchPattern, info.Name()); match {
				if skipPattern != "" && strings.Contains(filePath, skipPattern) {
					return nil
				}
				if strings.Contains(filePath, ".so") {
					reallyFiles[filePath] = ReadReallyFile(filePath)
				} else {
					reallyFiles[filePath] = filePath
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Error:", err)
	}
	return reallyFiles
}
func GenLink(source, target, pwd string) {
	klog.Infoln("gen link", source, "->", target)
	cmd := exec.Command("ln", "-s", source, target) // 你可以替换成你想要执行的命令和参数
	cmd.Dir = pwd
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		klog.Fatalln(err, out.String())
	}
}

func GetArchFromPath(libPath string) string {
	if strings.Contains(libPath, "lib64") {
		return "64"
	}
	if strings.Contains(libPath, "x86_64") {
		return "64"
	}

	cmd := exec.Command("objdump", "-p", libPath) // 你可以替换成你想要执行的命令和参数
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()
	outStr := out.String()
	if strings.Contains(outStr, "elf64-x86-64") {
		return "64"
	}
	return ""
}

func copyFileWithModeAndOwnership(src, dst string) error {
	klog.Infof("copy file from %s -> %s", src, dst)
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	os.ReadFile(dst)
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return err
	}

	if stat, ok := sourceInfo.Sys().(*syscall.Stat_t); ok {
		err = os.Chown(dst, int(stat.Uid), int(stat.Gid))
		if err != nil {
			return err
		}
	}

	return nil
}
