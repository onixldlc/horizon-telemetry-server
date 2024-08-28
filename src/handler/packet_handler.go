package handler

import (
	"fmt"
	"forza_telemetry/parser"
	"net"
)

// Handler processes incoming UDP packets.
func PacketHandler(data []byte, len int, addr *net.UDPAddr, manager *Manager) {
	switch len {
	case (120):
		fmt.Print("Seems like a SLED DATA")
		data_sled := parser.NewDataSled(data)
		message := data_sled.ParseSledData()
		manager.SendMessage([]byte(message))
	case (218):
		fmt.Print("Seems like a CAR DASH DATA")
		data_sled := parser.NewDataCarDash(data)
		message := data_sled.ParseCarDashData()
		manager.SendMessage([]byte(message))
	default:
		fmt.Printf("im unfamiliar with this packet size (%d) !", len)
	}

	fmt.Printf("Received packet from %s with data:\n", addr)
	hexDump(data)
	// Add more processing logic here as needed.
}
