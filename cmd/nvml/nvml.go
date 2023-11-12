package main

import (
	"fmt"
	"tkestack.io/nvml"
)

func main() {
	fmt.Println(nvml.Init())
}
