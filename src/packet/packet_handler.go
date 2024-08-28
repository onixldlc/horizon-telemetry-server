package packet

import (
	"fmt"
	"net"
)

// Handler processes incoming UDP packets.
func Handler(data []byte, len int, addr *net.UDPAddr) {
	fmt.Printf("Received packet from %s with data: %s\n", addr, string(data))
	hexDump(data)
	// Add more processing logic here as needed.
}
