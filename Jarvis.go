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
	fs1, _ := os.Open(fileName)
	content, _ := csv.NewReader(fs1).ReadAll()
	defer fs1.Close()

	for _, row := range content {
		j = append(j, row)
	}
}

func main() {
	Weather_url := "https://restapi.amap.com/v3/weather/weatherInfo?city=" + "" + "&key=c019f803910fb10b03043c50116b9be5"
	IP_url := "https://restapi.amap.com/v3/ip?output=xml&key=538f2f4c317fe3031b9e87bc38722a0a"
	Wr := HttpGet(Weather_url)
	Ir := HttpGet(IP_url)
	fmt.Println("Hello,sir!")
	CheckServer()
	fmt.Println(Wr, Ir)
}
