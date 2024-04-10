package parser

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	h := NewHeaderSection()
	h.AddPID(1234).AddQR(1).AddOpCode(0).AddAA(0).AddTC(0).AddRD(0).AddRA(0).AddZ(0).AddRcode(0).AddQdCount(1).AddAnCount(1)
	fmt.Println(h.Header)
}

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
