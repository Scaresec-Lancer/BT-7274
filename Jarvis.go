package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//查询天气的函数,用到了DoHttpGetRequest函数，高德API只允许Get请求
func Weather() {
	weather_result, err := DoHttpGetRequest("https://restapi.amap.com/v3/weather/weatherInfo?key=c019f803910fb10b03043c50116b9be5&city=110101")
	if err != nil {
		fmt.Println("network error")
	} else {
		fmt.Println(weather_result)
	}
}

func DoHttpGetRequest(url string) (rlt string, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		} else {
			return string(body), err
		}
	}
}

func main() {
	fmt.Println("Hello,sir!")
	Weather()
}
