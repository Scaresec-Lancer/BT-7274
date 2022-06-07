package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

func getTime() {

	t1 := time.Now().Hour()
	if t1 >= 18 {
		fmt.Print("晚上好 ")
	} else if t1 >= 14 {
		fmt.Print("下午好 ")
	} else if t1 >= 7 {
		fmt.Print("上午好 ")
	} else {
		fmt.Print("凌晨好 ")
	}

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

func readCsv() {

	//准备读取文件
	filePath := "D:\\test.csv"
	fs, err := os.Open(filePath)

	//打开文件报错输出
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	//针对大文件，一行一行的读取文件
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()

		//如果错误不为空（有报错）而且不是因为读取到文件结尾了，直接将错误原文输出
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}

		//如果报错是因为读取文件到结尾了
		if err == io.EOF {
			break
		}

		//没有错误，正常进行读取
		fmt.Println(row, len(row))

	}
}

func Hello() {
	getTime()
	getSys()

}
