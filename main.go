package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akrck02/whisper/service"
)

func main() {
	// go service.StartRTC()

	http.HandleFunc("/ws", service.HandleWebsocketConnections)

	fmt.Println("http server started on :8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
