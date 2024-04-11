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
	ph, _ := DeserializeHeader(b)
	fmt.Println("Deserialized Packet Id - ", ph.PacketId)
	fmt.Println("Deserialized QR - ", ph.QR)
	fmt.Println("Deserialized Op code - ", ph.OpCode)
	fmt.Println("Deserialized Rcode - ", ph.Rcode)
}

func TestQuestion(t *testing.T) {
	q := NewQuestionSection()
	q.AddName("google.com").AddType(1).AddClass(1)
	fmt.Println("Encoded Label Sequence - ", q.Name)
	fmt.Println("Entire Question Section - ", q.ToBytes())
}

func TestDeserializeQuestionSection(t *testing.T) {
	q := NewQuestionSection()
	q.AddName("google.com").AddType(1).AddClass(1)
	b := q.ToBytes()
	dq, _ := DeserializeQuestionSection(b)
	fmt.Println("Deserialised Question Section Name - ", string(dq.Name))
	fmt.Println("Deserialised Question Section Type - ", dq.Type)
	fmt.Println("Deserialised Question Section Class - ", dq.Class)
}

func TestAnswer(t *testing.T) {
	a := NewAnswerSection()
	a.AddName("codecraftery.io").AddType(1).AddClass(1).AddTTL(60).AddLength(4).AddData("8.8.8.8")
	fmt.Printf("%X\n", a.Name)
}

func TestDeserializeAnswerSection(t *testing.T) {
	a := NewAnswerSection()
	a.AddName("codecrafters.io").AddType(1).AddClass(1).AddTTL(60).AddLength(4).AddData("8.8.8.8")
	b := a.ToBytes()
	da := DeserializeAnswerSection(b)
	fmt.Println("Deserialised Answer Section Name - ", string(da.Name))
	fmt.Println("Deserialised Answer Section Type - ", da.Type)
	fmt.Println("Deserialised Answer Section Class - ", da.Class)
	fmt.Println("Deserialised Answer Section TTL - ", da.TTL)
	fmt.Println("Deserialised Answer Section Length - ", da.Length)
	fmt.Println("Deserialised Answer Section Data - ", da.Data)
}
