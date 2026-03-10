package main

import (
	"net/http"

	"github.com/akrck02/whisper/modules/api"
	"github.com/akrck02/whisper/sdk/logger"
	services "github.com/akrck02/whisper/service"
)

func main() {

	logger.Pretty()

	startApi()
	startMediaServices()
}

func startApi() {
	api.Start()
}

func startMediaServices() {

	go services.StartRTC()

	http.HandleFunc("/ws", services.WebSocketSignaling)

	logger.Info("http server started on :8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err.Error())
	}
}
