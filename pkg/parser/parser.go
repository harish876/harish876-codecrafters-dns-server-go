package parser

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type Message struct {
	Header   []byte
	Question []byte
}

type HeaderSection struct {
	Header []byte
}

func NewHeaderSection() HeaderSection {
	return HeaderSection{
		Header: make([]byte, 12),
	}
}

func (h *HeaderSection) AddPID(pid uint16) *HeaderSection {
	value := uint16(pid)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(h.Header[0:2], intBytes)
	return h
}

func (h *HeaderSection) AddQR(flag uint8) *HeaderSection {
	h.Header[2] |= flag << 7
	return h
}

func (h *HeaderSection) AddOpCode(flag uint8) *HeaderSection {
	h.Header[2] |= byte((flag & 0xF) << 3)
	return h
}

func (h *HeaderSection) AddAA(flag uint8) *HeaderSection {
	h.Header[2] |= (flag & 1) << 2
	return h
}

func (h *HeaderSection) AddTC(flag uint8) *HeaderSection {
	h.Header[2] |= (flag & 1) << 1
	return h
}

func (h *HeaderSection) AddRD(flag uint8) *HeaderSection {
	h.Header[2] |= (flag & 1)
	return h
}

func (h *HeaderSection) AddRA(flag uint8) *HeaderSection {
	h.Header[3] |= (flag & 1) << 7
	return h
}

func (h *HeaderSection) AddZ(flag uint8) *HeaderSection {
	h.Header[3] |= byte((flag & 0x3) << 4)
	return h
}

func (h *HeaderSection) AddRcode(flag uint8) *HeaderSection {
	h.Header[3] |= byte((flag & 0xF))
	return h
}

func (h *HeaderSection) AddQdCount(flag uint16) *HeaderSection {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(h.Header[4:6], intBytes)
	return h
}

func (h *HeaderSection) AddAnCount(flag uint16) *HeaderSection {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(h.Header[6:8], intBytes)
	return h
}

func (h *HeaderSection) AddNsCount(flag uint16) *HeaderSection {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(h.Header[8:10], intBytes)
	return h
}

func (h *HeaderSection) AddArCount(flag uint16) *HeaderSection {
	value := uint16(flag)
	intBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(intBytes, value)
	copy(h.Header[10:12], intBytes)
	return h
}

func (m *HeaderSection) GetHeader() []byte {
	return m.Header
}

func (m *HeaderSection) PrintHeader(header []byte) {
	for _, b := range m.Header {
		fmt.Println(b)
	}
}

type QuestionSection struct {
	/* A domain name, represented as a sequence of "labels" (more on this below) */
	Name []byte
	/* 2-byte int; the type of record (1 for an A record, 5 for a CNAME record etc., full list here) */
	Type []byte
	/* 2-byte int; usually set to 1 (full list here) */
	Class []byte
}

func NewQuestionSection() QuestionSection {
	return QuestionSection{
		Name:  make([]byte, 0),
		Type:  make([]byte, 2),
		Class: make([]byte, 2),
	}
}

func EncodeName(domain string) []byte {
	name := make([]byte, 0)
	domainArray := strings.Split(domain, ".")
	for _, d := range domainArray {
		name = append(name, byte(len(d)))
		name = append(name, []byte(d)...)
	}
	name = append(name, 0x00)
	return name
}

func (q *QuestionSection) AddType(typ uint16) *QuestionSection {
	binary.BigEndian.PutUint16(q.Type, typ)
	return q
}

func (q *QuestionSection) AddClass(typ uint16) *QuestionSection {
	binary.BigEndian.PutUint16(q.Class, typ)
	return q
}

func (q *QuestionSection) AddName(domain string) *QuestionSection {
	q.Name = EncodeName(domain)
	return q
}
