package main

import (
	"fmt"
	"runtime"
	"time"
)

func getTime() {
	t1 := time.Now().Hour()
	fmt.Println("现在是", t1, "时，先生")
}

func getSys() {
	sysType := runtime.GOOS

	if sysType == "linux" {
		fmt.Println("当前运行在Linux平台")
	} else if sysType == "windows" {
		fmt.Println("当前运行在Windows平台")
	}
}

func Hello() {
	getTime()
	getSys()
}
