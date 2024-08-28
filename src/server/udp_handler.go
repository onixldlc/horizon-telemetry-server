package server

import (
	"fmt"
	"forza_telemetry/handler"
	"net"
)

// Define a callback type
type DataHandler func(data []byte, len int, addr *net.UDPAddr, manager *handler.Manager)

func SetupUDPServer(address string, callback DataHandler, manager *handler.Manager) {
	listenAddr := address
	udpAddr, _ := net.ResolveUDPAddr("udp", listenAddr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	// Start listening
	go listenUDP(conn, callback, manager)
	fmt.Println("Starting udp server, listening at", address)

	// Keep the main goroutine running.
	select {}
}

// Listen for UDP packets and invoke the callback.
func listenUDP(conn *net.UDPConn, callback DataHandler, manager *handler.Manager) {
	buffer := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error on reading:", err)
			continue
		}
		// Trigger the callback with data and sender address
		callback(buffer[:n], n, addr, manager)
	}
}
