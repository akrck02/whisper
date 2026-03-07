package service

import (
	"net/http"

	"github.com/akrck02/whisper/sdk/logger"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing
	},
}

func WebSocketSignaling(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Read every message in a new go routine
	go func() {
		for {
			messageType, p, err := ws.ReadMessage()
			if err != nil {
				logger.Errorf(err)
				return
			}
			logger.Info(string(p))

			if err := ws.WriteMessage(messageType, p); err != nil {
				logger.Errorf(err)
				return
			}
		}
	}()
}
