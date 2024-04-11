package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/pkg/parser"
)

// add better doc support for all the header fields - enums
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

		reqHeader, offset1 := parser.DeserializeHeader(buf[:12])
		reqQuestion, _ := parser.DeserializeQuestionSection(buf[offset1:])

		q := parser.NewQuestionSection()
		q.AddName(string(reqQuestion.Name)).AddType(1).AddClass(1)

		a := parser.NewAnswerSection()
		a.AddName(string(reqQuestion.Name)).AddType(1).AddClass(1).AddTTL(60).AddLength(4).AddData("8.8.8.8")

		reqHeader.AddQR(1).AddAnCount(1).AddRcode(4)
		fmt.Println("Request RCode", reqHeader.Rcode)
		header := reqHeader.ToBytes()
		question := q.ToBytes()
		answer := a.ToBytes()

		response := []byte{}
		response = append(response, header...)
		response = append(response, question...)
		response = append(response, answer...)

		sentByteCount, err := udpConn.WriteToUDP(response, source)
		fmt.Println("Byte Count:", sentByteCount)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
