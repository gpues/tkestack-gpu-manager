package main

import (
	"bytes"
	"io"
	"k8s.io/klog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"tkestack.io/gpu-manager/pkg/server"
	"tkestack.io/gpu-manager/pkg/types"
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

func main() {

	vs := server.VolumeManager{}
	vs.Init()
	copyBinFileMap := map[string]string{}
	copyFileMap := map[string][2]string{}
	copyLinkMap := map[string][2]string{}
	for _, v := range vs.Config {
		os.RemoveAll(filepath.Join(types.DriverDir, v.Name))
		_ = os.MkdirAll(filepath.Join(types.DriverDir, v.Name, "lib"), os.ModeDir)
		_ = os.MkdirAll(filepath.Join(types.DriverDir, v.Name, "lib64"), os.ModeDir)
		_ = os.MkdirAll(filepath.Join(types.DriverDir, v.Name, "bin"), os.ModeDir)

		if v.Name != "nvidia" {
			continue
		}
		copyFileMap[types.ControlLib] = [2]string{v.Name, filepath.Base(types.ControlLib)}
		for _, lib := range v.Components["libraries"] {
			reallyFiles := SearchFile(types.FindBase, lib+"*", "stubs")
			for linkFile, reallyFile := range reallyFiles {
				if linkFile != reallyFile {
					continue
				}
				copyFileMap[reallyFile] = [2]string{v.Name, filepath.Base(linkFile)}
			}
			for linkFile, reallyFile := range reallyFiles { // 链接文件
				if linkFile == reallyFile {
					continue
				}
				copyLinkMap[linkFile] = [2]string{v.Name, reallyFile}
			}
		}
		for _, bin := range v.Components["binaries"] {
			reallyFiles := SearchFile(types.FindBase, bin, "")
			for linkFile, reallyFile := range reallyFiles {
				copyBinFileMap[reallyFile] = filepath.Join(types.DriverDir, v.Name, "bin", filepath.Base(linkFile))
			}
		}
	}
	for s, d := range copyBinFileMap {
		err := copyFileWithModeAndOwnership(s, d)
		if err != nil {
			klog.Fatalln(err)
		}
	}
	for reallyFile, ds := range copyFileMap {
		arch := GetArchFromPath(reallyFile)
		d := filepath.Join(filepath.Join(types.DriverDir, ds[0], "lib")+arch, ds[1])
		err := copyFileWithModeAndOwnership(reallyFile, d)
		if err != nil {
			klog.Fatalln(err)
		}
	}
	for linkFile, ds := range copyLinkMap {
		reallyFile := ds[1]
		arch := GetArchFromPath(reallyFile)
		server.GenLink(filepath.Base(reallyFile), filepath.Base(linkFile), filepath.Join(types.DriverDir, ds[0], "lib")+arch)
	}
	vs.Copy()

}

var fileMap = map[string]bool{}
var walkOnce sync.Once

func SearchFile(root string, searchPattern string, skipPattern string) map[string]string {
	reallyFiles := map[string]string{}
	walkOnce.Do(func() {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				if skipPattern != "" && strings.Contains(path, skipPattern) {
					return nil
				}
				fileMap[path] = true
			}
			return nil
		})
	})

	for filePath, _ := range fileMap {
		// 检查文件名是否符合规则
		if match, _ := filepath.Match(searchPattern, filepath.Base(filePath)); match {
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

	return reallyFiles
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
