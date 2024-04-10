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

		h := parser.NewHeaderSection()
		h.AddPID(1234).AddQR(1).AddQdCount(1).AddAnCount(1)

		q := parser.NewQuestionSection()
		q.AddName("codecrafters.io").AddType(1).AddClass(1)

		a := parser.NewAnswerSection()
		a.AddName("codecrafters.io").AddType(1).AddClass(1).AddTTL(60).AddLength(4).AddData("8.8.8.8")

		msg := parser.Message{}

		msg.Header = append(msg.Header, h.Header...)

		msg.Question = append(msg.Question, q.Name...)
		msg.Question = append(msg.Question, q.Type...)
		msg.Question = append(msg.Question, q.Class...)

		msg.Answer = append(msg.Answer, a.Name...)
		msg.Answer = append(msg.Answer, a.Type...)
		msg.Answer = append(msg.Answer, a.Class...)
		msg.Answer = append(msg.Answer, a.TTL...)
		msg.Answer = append(msg.Answer, a.Length...)
		msg.Answer = append(msg.Answer, a.Data...)

		response := []byte{}
		response = append(response, msg.Header...)
		response = append(response, msg.Question...)
		response = append(response, msg.Answer...)

		sentByteCount, err := udpConn.WriteToUDP(response, source)
		fmt.Println("Byte Count:", sentByteCount)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
