package main

import (
	"fmt"
	"forza_telemetry/packet"
	"net"
)

// Define a callback type
type DataHandler func(data []byte, len int, addr *net.UDPAddr)

func setupUDPServer(address string) {
	listenAddr := address
	udpAddr, _ := net.ResolveUDPAddr("udp", listenAddr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	// Start listening
	go listenUDP(conn, packet.Handler)
	fmt.Println("Starting udp server, listening at", address)

	// Keep the main goroutine running.
	select {}
}

// Listen for UDP packets and invoke the callback.
func listenUDP(conn *net.UDPConn, callback DataHandler) {
	buffer := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error on reading:", err)
			continue
		}
		// Trigger the callback with data and sender address
		callback(buffer[:n], n, addr)
	}
}

func main() {
	// Setup UDP server on localhost port 5300
	setupUDPServer("127.0.0.1:5300")
	// setupUDPServer("0.0.0.0:5300")
}
