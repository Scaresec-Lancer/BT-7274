package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	k string
	j [][]string
)

//查询天气的函数,用到了DoHttpGetRequest函数，高德API只允许Get请求
func Weather() {
	weather_result := DoHttpGetRequest("https://restapi.amap.com/v3/weather/weatherInfo?key=c019f803910fb10b03043c50116b9be5&city=110101")
	fmt.Println(weather_result)

}

func DoHttpGetRequest(url string) (rlt string) {

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

func ReadCsv() {
	//准备读取文件
	fileName := "map_code\\adcode.csv"
	fs1, _ := os.Open(fileName)
	content, _ := csv.NewReader(fs1).ReadAll()
	defer fs1.Close()

	for _, row := range content {
		j = append(j, row)
	}

}

func main() {
	fmt.Println("Hello,sir!")
	ReadCsv()
}
