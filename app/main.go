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
	count := 0
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
		count++
		_ = receivedData
		fmt.Printf("Received %d bytes from %s for count %d\n", size, source, count)

		reqHeader, offset := parser.DeserializeHeader(buf[:12])

		fmt.Printf("Number Of Questions : %d\n", reqHeader.QdCount)
		var questions []parser.QuestionSection
		for i := 0; i < int(reqHeader.QdCount); i++ {
			reqQuestion, _ := parser.DeserializeQuestionSection(buf[offset:])
			questions = append(questions, reqQuestion)
		}
		var answers []parser.AnswerSection
		for _, question := range questions {
			answer := parser.Answer(question)
			answers = append(answers, answer)
		}

		reqHeader.AddQR(1).AddAnCount(uint16(len(answers))).AddRcode(4)
		header := reqHeader.ToBytes()

		response := []byte{}
		response = append(response, header...)
		for _, question := range questions {
			q := parser.NewQuestionSection()
			q.AddName(string(question.Name)).AddType(1).AddClass(1)
			response = append(response, q.ToBytes()...)
		}
		for _, answer := range answers {
			response = append(response, answer.ToBytes()...)
		}

		sentByteCount, err := udpConn.WriteToUDP(response, source)
		_ = sentByteCount
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
