package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//对应json天气数据源的结构体
//{"status":"0","cityen":"shenzhen","city":"深圳","citycode":"101280601","temp":"6",
// "tempf":"42","wd":"东北风","wden":"NE","wdenimg":"//www.help.bj.cn/weather/images/wind/ne.png",
// "wdforce":"3级","wdspd":"12km\/h","humidity":"93%","stp":"1015","wisib":"11km","uptime":"00:10",
// "prcp":"1.5","prcp24h":"1.5","aqi":"15","pm25":"15","weather":"多云","weatheren":"Cloudy",
// "weathercode":"d01","weatherimg":"//www.help.bj.cn/weather/images/d01.png","today":"2022-02-21 周一"}

type WeatherInfoJson struct {
	Status     string `json:"status"`     //反馈代码 0成功
	Msg        string `json:"msg"`        //反馈信息
	Cityen     string `json:"cityen"`     //城市名称英文
	City       string `json:"city"`       //城市名称
	Citycode   string `json:"citycode"`   //城市编码
	Temp       string `json:"temp"`       //实时温度
	Tempf      string `json:"tempf"`      //华氏温度
	Wd         string `json:"wd"`         //风向
	Wden       string `json:"wden"`       //风向英文
	Wdforce    string `json:"wdforce"`    //风力
	Wdspd      string `json:"wdspd"`      //风速
	Uptime     string `json:"uptime"`     //更新时间
	Weather    string `json:"weather"`    //天气状况
	Weatheren  string `json:"weatheren"`  //天气状况英文
	Weatherimg string `json:"weatherimg"` //天气状况图标
	Stp        string `json:"stp"`        //气压
	Wisib      string `json:"wisib"`      //能见度
	Humidity   string `json:"humidity"`   //湿度
	Prcp       string `json:"prcp"`       //降雨
	Prcp24h    string `json:"prcp24h"`    //24小时降雨量
	Aqi        string `json:"aqi"`        //AQI
	Pm25       string `json:"pm25"`       //PM2.5
	Today      string `json:"today"`      //今天日期
}

//获取天气接口返回的json
func GetWeather(id string) {
	url := "http://api.help.bj.cn/apis/weather/?id=" + id
	resp, err := http.Get(url) //使用get方法访问
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()
	bodyContent, _ := ioutil.ReadAll(resp.Body) //读取流数据bodyContent
	// fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	// fmt.Printf("resp body data:[%s]\n", string(bodyContent))

	//解析json
	var jsonWeatherInfo WeatherInfoJson
	err2 := json.Unmarshal(bodyContent, &jsonWeatherInfo)
	if err2 != nil {
		fmt.Println("解析错误", err2)
		return
	}

	//判断有无解析json后的数据
	if len(jsonWeatherInfo.City) != 0 {
		println("------------------------------------")
		fmt.Printf("%v天气预报： 今天是%v\n", jsonWeatherInfo.City, jsonWeatherInfo.Today)
		fmt.Printf("\t\t气温%v度/%v\n", jsonWeatherInfo.Temp, jsonWeatherInfo.Weather)
		fmt.Printf("\t\t%v/%v\n", jsonWeatherInfo.Wd, jsonWeatherInfo.Wdforce)
		println("------------------------------------")
	}
	return
}

func main() {
	id := "101280601" //深圳编号
	GetWeather(id)
}
