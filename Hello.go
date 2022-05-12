package main

import (
	"fmt"
	"net"
	"time"
)

func CheckServer() {
	timeout := time.Duration(5 * time.Second)
	_, err := net.DialTimeout("tcp", "www.baidu.com:443", timeout)
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
		return
	}
	fmt.Println("网络连接正常")
}
