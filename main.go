package main

import (
	"log"
	"net/http"
)

func main() {

	// HTTP Server   http://127.0.0.1:22726/eink-weather/v1?XXX
	http.HandleFunc("/eink-weather/v1", EinkWeatherFunc)
	log.Println("Http start at port:22726")
	err := http.ListenAndServe(":22726", nil)
	if err != nil {
		log.Printf("http server failed, err:%v\n", err)
		return
	}
}
