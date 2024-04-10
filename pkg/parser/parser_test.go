package parser

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	h := NewHeaderSection()
	h.AddPID(1234).AddQR(1).AddOpCode(1)
	fmt.Println(h.ToBytes())
}

func TestDeserializeHeader(t *testing.T) {
	h := NewHeaderSection()
	h.AddPID(1234).AddQR(1).AddOpCode(1).AddRcode(4)

	b := h.ToBytes()
	ph := DeserializeHeader(b)
	fmt.Println("Deserialized Packet Id - ", ph.PacketId)
	fmt.Println("Deserialized QR - ", ph.QR)
	fmt.Println("Deserialized Op code - ", ph.OpCode)
	fmt.Println("Deserialized Rcode - ", ph.Rcode)
}

// 1 0 0 0
/* 1 1 0 0 0 0 0 0 */
/* 4 210 137 0 0 1 0 1 0 0 0 0 */

func TestQuestion(t *testing.T) {
	q := NewQuestionSection()
	q.AddName("google.com").AddType(1).AddClass(1)
	fmt.Printf("%X\n", q.Name)
}

// 06 67 6F 6F 67 6C 65 03 63 6F 6D 00

func TestAnswer(t *testing.T) {
	a := NewAnswerSection()
	a.AddName("codecraftery.io").AddType(1).AddClass(1).AddTTL(60).AddLength(4).AddData("8.8.8.8")
	fmt.Printf("%X\n", a.Name)
}
