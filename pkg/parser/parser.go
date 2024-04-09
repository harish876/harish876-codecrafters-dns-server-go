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

func (m *Message) AddPID(pid uint16) *Message {
	value := uint16(pid)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[0:2], intBytes)
	return m
}

func (m *Message) AddQR(flag uint8) *Message {
	m.Header[2] |= flag << 7
	return m
}

func (m *Message) AddOpCode(flag uint8) *Message {
	m.Header[2] |= byte((flag & 0xF) << 3)
	return m
}

func (m *Message) AddAA(flag uint8) *Message {
	m.Header[2] |= (flag & 1) << 2
	return m
}

func (m *Message) AddTC(flag uint8) *Message {
	m.Header[2] |= (flag & 1) << 1
	return m
}

func (m *Message) AddRD(flag uint8) *Message {
	m.Header[2] |= (flag & 1)
	return m
}

func (m *Message) AddRA(flag uint8) *Message {
	m.Header[3] |= (flag & 1) << 7
	return m
}

func (m *Message) AddZ(flag uint8) *Message {
	m.Header[3] |= byte((flag & 0x3) << 4)
	return m
}

func (m *Message) AddRcode(flag uint8) *Message {
	m.Header[3] |= byte((flag & 0xF))
	return m
}

func (m *Message) AddQdCount(flag uint16) *Message {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[4:6], intBytes)
	return m
}

func (m *Message) AddAnCount(flag uint16) *Message {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[6:8], intBytes)
	return m
}

func (m *Message) AddNsCount(flag uint16) *Message {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[8:10], intBytes)
	return m
}

func (m *Message) AddArCount(flag uint16) *Message {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(m.Header[10:12], intBytes)
	return m
}

func (m *Message) GetHeader() []byte {
	return m.Header
}

func (m *Message) PrintHeader(header []byte) {
	for _, b := range m.Header {
		fmt.Println(b)
	}
}
