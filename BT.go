package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func GetWithHttp(uri string) (code int) {
	response, _ := http.Get(uri)
	return response.StatusCode
}

func writeFile(str string) {
	filepath := "./dict.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("错误提示: ", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func main() {

	var path_list []string
	var x = ""
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		path = filepath.ToSlash(path)
		path_list = append(path_list, path)
		return nil
	})

	for _, value := range path_list {
		x += (value + "\n")
	}
	writeFile(x)

	s := GetWithHttp("https://baidu.com/")
	fmt.Println(s)
}
