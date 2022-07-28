package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var weatherType = map[string]string{
	"CLEAR_DAY":           "晴",
	"CLEAR_NIGHT":         "晴",
	"PARTLY_CLOUDY_DAY":   "多云",
	"PARTLY_CLOUDY_NIGHT": "多云",
	"CLOUDY":              "阴",
	"LIGHT_HAZE":          "轻霾",
	"MODERATE_HAZE":       "中霾",
	"HEAVY_HAZE":          "重霾",
	"LIGHT_RAIN":          "小雨",
	"MODERATE_RAIN":       "中雨",
	"HEAVY_RAIN":          "大雨",
	"STORM_RAIN":          "暴雨",
	"FOG":                 "雾",
	"LIGHT_SNOW":          "小雪",
	"MODERATE_SNOW":       "中雪",
	"HEAVY_SNOW":          "大雪",
	"STORM_SNOW":          "暴雪",
	"DUST":                "浮尘",
	"SAND":                "沙尘",
	"WIND":                "大风",
}

type WeatherDataType struct {
	OK                  bool   `json:"OK"`
	Weather             string `json:"weather,omitempty"`
	WeatherKey          string `json:"weatherKey,omitempty"`
	Temperature         string `json:"temperature,omitempty"`
	TemperatureRange    string `json:"temperatureRange,omitempty"`
	Comfort             string `json:"comfort,omitempty"`
	Keypoint            string `json:"keypoint,omitempty"`
	TomorrowWeather     string `json:"tomorrowWeather,omitempty"`
	TomorrowTemperature string `json:"tomorrowTemperature,omitempty"`
}

func getWeatherInfo(latlng string, token string) (w WeatherDataType) {
	url := "https://api.caiyunapp.com/v2.5/" + token + "/" + latlng + "/weather.json"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		w.OK = false
		return
	}
	res, err := client.Do(req)
	if err != nil {
		w.OK = false
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.OK = false
		log.Println(err)
		return
	}
	var srcmsg map[string]interface{}
	err = json.Unmarshal(body, &srcmsg)
	if err != nil {
		log.Println(err)
	}

	temperature, ok1 := srcmsg["result"].(map[string]interface{})["realtime"].(map[string]interface{})["temperature"]
	skycon, ok2 := srcmsg["result"].(map[string]interface{})["realtime"].(map[string]interface{})["skycon"]
	forecast_keypoint, ok3 := srcmsg["result"].(map[string]interface{})["forecast_keypoint"]
	life_index, ok4 := srcmsg["result"].(map[string]interface{})["realtime"].(map[string]interface{})["life_index"].(map[string]interface{})["comfort"].(map[string]interface{})["desc"]
	skycon_tom, ok5 := srcmsg["result"].(map[string]interface{})["daily"].(map[string]interface{})["skycon_08h_20h"].([]interface{})[1].(map[string]interface{})["value"]
	temp_max, _ := srcmsg["result"].(map[string]interface{})["daily"].(map[string]interface{})["temperature"].([]interface{})[0].(map[string]interface{})["max"]
	temp_min, _ := srcmsg["result"].(map[string]interface{})["daily"].(map[string]interface{})["temperature"].([]interface{})[0].(map[string]interface{})["min"]
	temp_t_max, _ := srcmsg["result"].(map[string]interface{})["daily"].(map[string]interface{})["temperature"].([]interface{})[1].(map[string]interface{})["max"]
	temp_t_min, _ := srcmsg["result"].(map[string]interface{})["daily"].(map[string]interface{})["temperature"].([]interface{})[1].(map[string]interface{})["min"]

	if ok1 && ok2 && ok3 && ok4 && ok5 {
		w = WeatherDataType{
			OK:                  true,
			Weather:             weatherType[fmt.Sprintf("%s", skycon)],
			WeatherKey:          fmt.Sprintf("%s", skycon),
			Temperature:         fmt.Sprintf("%.1f", temperature),
			Comfort:             fmt.Sprintf("%s", life_index),
			Keypoint:            fmt.Sprintf("%s", forecast_keypoint),
			TemperatureRange:    fmt.Sprintf("%.0f-%.0f C", temp_min, temp_max),
			TomorrowWeather:     weatherType[fmt.Sprintf("%s", skycon_tom)],
			TomorrowTemperature: fmt.Sprintf("%.0f-%.0f C", temp_t_min, temp_t_max),
		}
		return
	}
	w.OK = false
	return
}
