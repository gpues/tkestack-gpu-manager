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

// Init starts a VolumeManager
func (vm *VolumeManager) Init() (err error) {
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

				klog.Infof("Find %s binaries: %+v", cfg.Name, bins)

				vol.dirs = append(vol.dirs, volumeDir{binDir, bins})
			case "libraries":
				var libs32 []string
				var libs64 []string
				filepath.WalkDir(filepath.Join(types.DriverDir, cfg.Name, "lib"), func(path string, d fs.DirEntry, err error) error {
					if path != filepath.Join(types.DriverDir, cfg.Name, "lib") {
						libs32 = append(libs32, path)
					}
					return nil
				})
				filepath.WalkDir(filepath.Join(types.DriverDir, cfg.Name, "lib64"), func(path string, d fs.DirEntry, err error) error {
					if path != filepath.Join(types.DriverDir, cfg.Name, "lib64") {
						libs64 = append(libs64, path)
					}
					return nil
				})
				klog.Infof("Find %s 32bit libraries: %+v", cfg.Name, libs32)
				klog.Infof("Find %s 64bit libraries: %+v", cfg.Name, libs64)

				vol.dirs = append(vol.dirs, volumeDir{lib32Dir, libs32}, volumeDir{lib64Dir, libs64})
			}

			vols[cfg.Name] = vol
		}
	}
	for _, vol := range vols {
		for _, d := range vol.dirs {
			for _, f := range d.files {
				if strings.HasPrefix(path.Base(f), "libvgpu.so") {
					vm.cudaControlFile = f
				}
				if strings.HasPrefix(path.Base(f), "libcuda.so") {
					driverStr := strings.SplitN(strings.TrimPrefix(path.Base(f), "libcuda.so."), ".", -1)
					if len(driverStr) > 2 {
						types.DriverVersionMajor, _ = strconv.Atoi(driverStr[0])
						types.DriverVersionMinor, _ = strconv.Atoi(driverStr[1])
						klog.Infof("Driver version: %d.%d", types.DriverVersionMajor, types.DriverVersionMinor)
					}
				}
			}
		}
	}

	klog.Infoln("Share", vm.Share)

	klog.V(2).Infof("Volume manager is running")
	return nil
}

// #lizard forgives

func exist(Path string) (bool, error) {
	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GenLink(source, target, pwd string) {
	os.Remove(filepath.Join(pwd, target))
	klog.Infoln("gen link", source, "->", target)
	//err := os.Symlink(filepath.Join(pwd, source), filepath.Join(pwd, target))
	//if err != nil {
	//	klog.Fatalln(err)
	//}
	cmd := exec.Command("ln", "-s", source, target) // 你可以替换成你想要执行的命令和参数
	cmd.Dir = pwd
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		klog.Fatalln(err, out.String())
	}
}
