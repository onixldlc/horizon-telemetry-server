package server

import (
	"fmt"
	"forza_telemetry/handler"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type MessageHandler func(manager *handler.Manager, messageType int, message []byte)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var manager = handler.NewManager()

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	manager.Register(conn)

	log.Println("Client connected !")
}

func WebsocketHandler(address string) *handler.Manager {
	http.HandleFunc("/connect", receiveHandler)
	fmt.Printf("Websocket server started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
	return manager
}
