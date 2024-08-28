package main

import (
	"forza_telemetry/handler"
	"forza_telemetry/server"
)

func main() {
	// Setup UDP server on localhost port 5300
	// setupUDPServer("127.0.0.1:5300")
	wsManager := server.WebsocketHandler("127.0.0.1:5400")
	server.SetupUDPServer("0.0.0.0:5300", handler.PacketHandler, wsManager)
}
