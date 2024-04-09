package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/pkg/parser"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("Listening on Port 2053...")
	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)
		msg := parser.NewMessage(receivedData)
		msg.AddPID(1234).AddQR(1).AddOpCode(0).AddAA(0).AddTC(0).AddRD(0).AddRA(0).AddZ(0).AddRcode(0)
		response := msg.GetHeader()

		sentByteCount, err := udpConn.WriteToUDP(response, source)
		fmt.Println("Byte Count:", sentByteCount)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
