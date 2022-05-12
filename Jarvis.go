package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	k    string
	j    [][]string
	b    []string
	l    string
	Data map[interface{}]interface{}
)

//做GET请求的函数
func HttpGet(url string) (rlt string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			k = string(body)
		}
	}
	return k
}

//读取CSV的函数
func ReadCsv() {
	fileName := "map_code\\adcode.csv"
	// 创建句柄
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)
	Data = make(map[interface{}]interface{})
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		b = strings.Split(line, ",")
		Data[b[0]] = b[1]

	}
}

func main() {

	fmt.Println("Hello,sir!")
	CheckServer()
	ReadCsv()
	l = "中华人民共和国"
	fmt.Println(Data[l])

}
