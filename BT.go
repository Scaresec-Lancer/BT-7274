package main

import (
	"fmt"
	"runtime"
	"time"
)

func getTime() {
	timeStr := time.Now().Format("2006-01-02 15:04")
	fmt.Println(timeStr)
}

func getSys() (os, arch string, core int) {

	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS
	sysArch := runtime.GOARCH
	sysCore := runtime.NumCPU()

	return sysType, sysArch, sysCore
}

func main() {
	os, arch, core := getSys()
	fmt.Println(os, arch, core)
	getTime()
}
