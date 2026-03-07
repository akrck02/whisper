package main

import (
	"net/http"

	"github.com/akrck02/whisper/sdk/logger"
	"github.com/akrck02/whisper/service"
)

func main() {

	logger.Pretty()

	//startMediaServices()
}

func startMediaServices() {
	go service.StartRTC()

	http.HandleFunc("/ws", service.WebSocketSignaling)

	logger.Info("http server started on :8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err.Error())
	}
}
