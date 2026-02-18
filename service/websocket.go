package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing
	},
}

func HandleWebsocketConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Read every message in a new go routine
	go func() {
		for {
			messageType, p, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(p))

			if err := ws.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}()
}
