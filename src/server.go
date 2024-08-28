package main

import (
	"fmt"
	"net"
)

// Define a callback type
type DataHandler func(data string, addr *net.UDPAddr)

func main() {
	listenAddr := "0.0.0.0:5300"
	udpAddr, _ := net.ResolveUDPAddr("udp", listenAddr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	// Start listening
	go listenUDP(conn, handleData)

	// Keep the main goroutine running.
	select {}
}

// Listen for UDP packets and invoke the callback.
func listenUDP(conn *net.UDPConn, callback DataHandler) {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error on reading:", err)
			continue
		}
		// Trigger the callback with data and sender address
		callback(string(buffer[:n]), addr)
	}
}

// Define your handler function
func handleData(data string, addr *net.UDPAddr) {
	fmt.Printf("Received from %v: %s\n", addr, data)
}
