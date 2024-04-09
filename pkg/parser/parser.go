package parser

import (
	"encoding/binary"
	"fmt"
)

type Message struct {
	Header []byte
}

func NewMessage(data string) Message {
	return Message{
		Header: make([]byte, 12),
	}
}

func (m *Message) AddPID(pid uint16) {
	value := uint16(pid)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[0:2], intBytes)
}

func (m *Message) AddQR(flag uint8) {
	m.Header[2] |= flag << 7
}

func (m *Message) AddOpCode(flag uint8) {
	m.Header[2] |= byte((flag & 0xF) << 3)
}

func (m *Message) AddAA(flag uint8) {
	m.Header[2] |= (flag & 1) << 2
}

func (m *Message) AddTC(flag uint8) {
	m.Header[2] |= (flag & 1) << 1
}

func (m *Message) AddRD(flag uint8) {
	m.Header[2] |= (flag & 1)
}

func (m *Message) AddRA(flag uint8) {
	m.Header[3] |= (flag & 1) << 7
}

func (m *Message) AddZ(flag uint8) {
	m.Header[3] |= byte((flag & 0x3) << 4)
}

func (m *Message) AddRcode(flag uint8) {
	m.Header[3] |= byte((flag & 0xF))
}

func (m *Message) AddQdCount(flag uint16) {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[4:6], intBytes)
}

func (m *Message) AddAnCount(flag uint16) {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[6:8], intBytes)
}

func (m *Message) AddNsCount(flag uint16) {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[8:10], intBytes)
}

func (m *Message) AddArCount(flag uint16) {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[10:12], intBytes)
}

func (m *Message) AddDefaultHeader() {
	m.AddPID(1234)
	m.AddQR(1)
	m.AddOpCode(0)
	m.AddAA(0)
	m.AddTC(0)
	m.AddRD(0)
	m.AddRA(0)
	m.AddZ(0)
	m.AddRcode(0)
}

func (m *Message) GetHeader() []byte {
	return m.Header
}

func (m *Message) PrintHeader(header []byte) {
	for _, b := range m.Header {
		fmt.Println(b)
	}
}
