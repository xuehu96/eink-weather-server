package main

import (
	"eink-cyweather/piximg"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Weather WeatherDataType

func EinkWeatherFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 检查方法 如果不是GET  打回去
	if r.Method != "GET" {
		w.WriteHeader(405)
		log.Println("Method not equal POST, return header 405")
		fmt.Fprintf(w, "405Method Not Allowed")
		return
	}
	// 2. 检查参数
	data := r.URL.Query()
	latlng := data.Get("l")
	token := data.Get("t")
	cdTxt := data.Get("txt")
	cdDate := data.Get("date")
	if latlng == "" || token == "" {
		w.WriteHeader(400)
		log.Println("GET parameter err")
		fmt.Fprintf(w, "400 parameter err")
		return
	}
	if cdTxt == "" {
		cdTxt = "跑路"
	}
	if cdDate == "" {
		cdDate = "20221231"
	}

	// 3. 获取天气信息
	var weather WeatherDataType
	// 以下为夜间节省API次数, 如果夜间需要实时更新，改为true
	nighton := false
	h, m, _ := time.Now().Clock()
	if h >= 7 || nighton {
		weather = getWeatherInfo(latlng, token)
		Weather = weather
	} else {
		if m <= 5 {
			// 夜间前5分钟刷新
			weather = getWeatherInfo(latlng, token)
			Weather = weather
		} else {
			weather = Weather
		}
	}
	log.Println(weather)

	// 4. 拼装图片
	dc := piximg.Create(200, 200)

	// 画天气图标
	piximg.DrawWeatherIcon(dc, weather.WeatherKey)
	// 画日期
	piximg.DrawDate(dc)
	// 画天气实时预报
	piximg.DrawWeatherDescribe(dc, weather.Keypoint)
	// 画天气预报
	piximg.DrawForecast(dc,
		fmt.Sprintf("今天: %s %s", weather.Weather, weather.TemperatureRange),
		fmt.Sprintf("明天: %s %s", weather.TomorrowWeather, weather.TomorrowTemperature),
	)
	// 画温度
	piximg.DrawTemp(dc, weather.Temperature, weather.Comfort)
	// 画底部
	piximg.DrawBottom(dc)
	piximg.DrawCountdown(dc, cdTxt, cdDate)

	// 5. 转为墨水屏的数据格式 generate_eink_bytes
	einkByte := piximg.GenerateEinkBytes(dc.Image())

	// 测试
	dc.SavePNG("out.png")

	// 6. 返回数据
	w.WriteHeader(200)
	w.Write(einkByte)
	log.Printf("HTTP OK")
}
