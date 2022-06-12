package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//对应json天气数据源的结构，头字母大写
type WeatherInfoJson struct {
	Pub        string
	Name       string
	Wind       WindObject
	Astronomy  AstronomyObject
	Atmosphere AtmosphereObject
	Forecasts  []ForecastsObject
}

type WindObject struct {
	Chill     float64
	Direction float64
	Speed     float64
}

type AstronomyObject struct {
	Sunrise string
	Sunset  string
}

type AtmosphereObject struct {
	Humidity   float64
	Visibility float64
	Pressure   float64
	Rising     float64
}

type ForecastsObject struct {
	Date        string
	Day         float64
	Code        float64
	Text        string
	Low         float64
	High        float64
	Image_large string
	Image_small string
}

func GetWeather(city string) {

	str := "http://api.help.bj.cn/apis/weather/?id=" + city

	resp, err := http.Get(str) //使用get方法访问
	if err != nil {
		println(err)
		return
	}

	defer resp.Body.Close()
	input, err1 := ioutil.ReadAll(resp.Body) //读取流数据
	if err1 != nil {
		return
	}

	var jsonWeather WeatherInfoJson
	err2 := json.Unmarshal(input, &jsonWeather) //解析json数据
	if err2 != nil {
		return
	}

	if len(jsonWeather.Name) != 0 { //判断有无解析数据
		for i := 0; i < 3; i++ {
			fmt.Printf("-->地区:%s 时间:%s 温度:%d-%d 天气:%s 发布时间:%s\n", jsonWeather.Name, jsonWeather.Forecasts[i].Date, int(jsonWeather.Forecasts[i].Low), int(jsonWeather.Forecasts[i].High), jsonWeather.Forecasts[i].Text, jsonWeather.Pub)
		}
	}
}

func main() {

	GetWeather("101280601")
	// GetWeather("nanning")

}
