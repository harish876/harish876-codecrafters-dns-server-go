package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

var (
	TYPE_A     = 1
	CLASS_IN   = 1
	HEADER_LEN = 12
)

type Message struct {
	Header   []byte
	Question []byte
	Answer   []byte
}

type HeaderSection struct {
	PacketId uint16 //16 bits A random ID assigned to query packets. Response packets must reply with the same ID
	QR       uint8  //1 bit	1 for a reply packet, 0 for a question packet.
	OpCode   uint8  //4 bits	Specifies the kind of query in a message.
	AA       uint8  //1 bit	1 if the responding server "owns" the domain queried, i.e., it's authoritative.
	TC       uint8  //1 if the message is larger than 512 bytes. Always 0 in UDP responses.
	RD       uint8  //1 bit	Sender sets this to 1 if the server should recursively resolve this query, 0 otherwise.
	RA       uint8  //Recursion Available 1 bit	Server sets this to 1 to indicate that recursion is available.
	Z        uint8  //Reserved 3 bits	Used by DNSSEC queries. At inception, it was reserved for future use.
	Rcode    uint8  //4 bits	Response code indicating the status of the response.
	QdCount  uint16 //16 bits	Number of questions in the Question section.
	AnCount  uint16 //16 bits	Number of records in the Answer section.
	NsCount  uint16 //Authority Record Count 16 bits	Number of records in the Authority section.
	ArCount  uint16 // Additional Record Count16 bits	Number of records in the Additional section.
}

func NewHeaderSection() HeaderSection {
	return HeaderSection{}
}

func (h *HeaderSection) AddPID(pid uint16) *HeaderSection {
	h.PacketId = pid
	return h
}

func (h *HeaderSection) AddQR(flag uint8) *HeaderSection {
	h.QR = flag
	return h
}

func (h *HeaderSection) AddOpCode(flag uint8) *HeaderSection {
	h.OpCode = flag
	return h
}

func (h *HeaderSection) AddAA(flag uint8) *HeaderSection {
	h.AA = flag
	return h
}

func (h *HeaderSection) AddTC(flag uint8) *HeaderSection {
	h.TC = flag
	return h
}

func (h *HeaderSection) AddRD(flag uint8) *HeaderSection {
	h.RD = flag
	return h
}

func (h *HeaderSection) AddRA(flag uint8) *HeaderSection {
	h.RA = flag
	return h
}

func (h *HeaderSection) AddZ(flag uint8) *HeaderSection {
	h.Z = flag
	return h
}

func (h *HeaderSection) AddRcode(flag uint8) *HeaderSection {
	h.Rcode = flag
	return h
}

func (h *HeaderSection) AddQdCount(flag uint16) *HeaderSection {
	h.QdCount = flag
	return h
}

func (h *HeaderSection) AddAnCount(flag uint16) *HeaderSection {
	h.AnCount = flag
	return h
}

func (h *HeaderSection) AddNsCount(flag uint16) *HeaderSection {
	h.NsCount = flag
	return h
}

func (h *HeaderSection) AddArCount(flag uint16) *HeaderSection {
	h.ArCount = flag
	return h
}

func (h *HeaderSection) ToBytes() []byte {
	result := make([]byte, 12)

	binary.BigEndian.PutUint16(result[0:2], h.PacketId)

	result[2] |= h.QR << 7
	result[2] |= byte((h.OpCode & 0xF) << 3)
	result[2] |= (h.AA & 1) << 2
	result[2] |= (h.TC & 1) << 1
	result[2] |= (h.RD & 1)

	result[3] |= (h.RA & 1) << 7
	result[3] |= byte((h.Z & 0x3) << 4)
	result[3] |= byte((h.Rcode & 0xF))

	binary.BigEndian.PutUint16(result[4:6], h.QdCount)
	binary.BigEndian.PutUint16(result[6:8], h.AnCount)
	binary.BigEndian.PutUint16(result[8:10], h.NsCount)
	binary.BigEndian.PutUint16(result[10:12], h.ArCount)

	return result
}

type QuestionSection struct {
	Name  []byte // A domain name, represented as a sequence of "labels" (more on this below)
	Type  uint16 // 2-byte int; the type of record (1 for an A record, 5 for a CNAME record etc., full list here)
	Class uint16 // 2-byte int; usually set to 1 (full list here)
}

func NewQuestionSection() QuestionSection {
	return QuestionSection{
		Name: make([]byte, 0),
	}
}

func EncodeLabelSequence(domain string) []byte {
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
	q.Type = typ
	return q
}

func (q *QuestionSection) AddClass(class uint16) *QuestionSection {
	q.Class = class
	return q
}

func (q *QuestionSection) AddName(domain string) *QuestionSection {
	q.Name = EncodeLabelSequence(domain)
	return q
}

func (q *QuestionSection) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, q.Name...)
	result = binary.BigEndian.AppendUint16(result, q.Type)
	result = binary.BigEndian.AppendUint16(result, q.Class)
	return result
}

type AnswerSection struct {
	Name   []byte // Label Sequence	The domain name encoded as a sequence of labels.
	Type   uint16 // 2-byte Integer	1 for an A record, 5 for a CNAME record etc., full list here
	Class  uint16 // 2-byte Integer	Usually set to 1 (full list here)
	TTL    uint32 // 4-byte Integer	The duration in seconds a record can be cached before requerying.
	Length uint16 // 2-byte Integer	Length of the RDATA field in bytes.
	Data   []byte // Variable	Data specific to the record type.
}

func NewAnswerSection() AnswerSection {
	return AnswerSection{
		Name: make([]byte, 0),
		Data: make([]byte, 0),
	}
}

func (a *AnswerSection) AddName(domain string) *AnswerSection {
	a.Name = EncodeLabelSequence(domain)
	return a
}

func (a *AnswerSection) AddType(typ uint16) *AnswerSection {
	a.Type = typ
	return a
}

func (a *AnswerSection) AddClass(class uint16) *AnswerSection {
	a.Class = class
	return a
}

func (a *AnswerSection) AddTTL(ttl uint32) *AnswerSection {
	a.TTL = ttl
	return a
}

func (a *AnswerSection) AddLength(length uint16) *AnswerSection {
	a.Length = length
	return a
}

func (a *AnswerSection) AddData(data string) *AnswerSection {
	a.Data = EncodeDomain(data)
	return a
}

func (a *AnswerSection) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, a.Name...)
	result = binary.BigEndian.AppendUint16(result, a.Type)
	result = binary.BigEndian.AppendUint16(result, a.Class)
	result = binary.BigEndian.AppendUint32(result, a.TTL)
	result = binary.BigEndian.AppendUint16(result, a.Length)
	result = append(result, a.Data...)

	return result
}

func DeserializeHeader(buf []byte) (HeaderSection, int) {
	header := buf[:12]
	var packetId uint16
	var qdCount uint16
	var anCount uint16
	var nsCount uint16
	var arCount uint16
	binary.Read(bytes.NewReader(header[:4]), binary.BigEndian, &packetId)
	binary.Read(bytes.NewReader(header[4:6]), binary.BigEndian, &qdCount)
	binary.Read(bytes.NewReader(header[6:8]), binary.BigEndian, &anCount)
	binary.Read(bytes.NewReader(header[8:10]), binary.BigEndian, &nsCount)
	binary.Read(bytes.NewReader(header[10:12]), binary.BigEndian, &arCount)

	return HeaderSection{
		PacketId: packetId,
		QR:       (header[2] >> 7) & 1,
		OpCode:   (header[2] >> 3) & 0x0F,
		AA:       (header[2] >> 2) & 1,
		TC:       (header[2] >> 1) & 1,
		RD:       header[2] & 1,
		RA:       (header[3] >> 7) & 1,
		Z:        (header[3] >> 4) & 0x3,
		Rcode:    header[3] & 0x0F,
		QdCount:  qdCount,
		AnCount:  anCount,
		NsCount:  nsCount,
		ArCount:  arCount,
	}, 12
}

func DeserializeQuestionSection(data []byte) (QuestionSection, int) {
	name, endIdx := DecodeLabelSequence(data)
	var typ uint16
	var class uint16

	if endIdx >= len(data) {
		return QuestionSection{}, 0
	}

	endIdx++
	binary.Read(bytes.NewBuffer(data[endIdx:endIdx+2]), binary.BigEndian, &typ)
	binary.Read(bytes.NewBuffer(data[endIdx+2:endIdx+4]), binary.BigEndian, &class)

	q := QuestionSection{
		Name:  name,
		Type:  typ,
		Class: class,
	}
	return q, len(q.ToBytes())
}

func DeserializeAnswerSection(data []byte) AnswerSection {
	var typ uint16
	var class uint16
	var ttl uint32
	var length uint16
	name, endIdx := DecodeLabelSequence(data)
	if endIdx >= len(data) {
		return AnswerSection{}
	}
	endIdx++
	binary.Read(bytes.NewBuffer(data[endIdx:endIdx+2]), binary.BigEndian, &typ)
	binary.Read(bytes.NewBuffer(data[endIdx+2:endIdx+4]), binary.BigEndian, &class)
	binary.Read(bytes.NewBuffer(data[endIdx+4:endIdx+8]), binary.BigEndian, &ttl)
	binary.Read(bytes.NewBuffer(data[endIdx+8:endIdx+10]), binary.BigEndian, &length)
	domain := DecodeDomain(data[endIdx+10 : endIdx+14])
	fmt.Println(domain)
	return AnswerSection{
		Name:   name,
		Type:   typ,
		Class:  class,
		TTL:    ttl,
		Length: length,
		Data:   []byte(domain),
	}
}

func DecodeLabelSequence(data []byte) ([]byte, int) {
	curr := 0
	var name []byte
	for {
		if data[curr] == 0 || curr >= len(data) {
			break
		}

		length := data[curr]
		curr++
		var word []byte
		for i := 0; i < int(length); i++ {
			word = append(word, data[curr])
			curr++
		}
		name = append(name, word...)
		name = append(name, '.')
	}
	//removed the trailing .
	return name[:len(name)-1], curr
}

func EncodeDomain(data string) []byte {
	ip := make([]byte, 0)
	ipArray := strings.Split(data, ".")

	for _, val := range ipArray {
		num, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		ip = append(ip, byte(num))
	}
	return ip
}

// Todo. Incomplete
func DecodeDomain(buf []byte) string {
	var value uint32
	binary.Read(bytes.NewBuffer(buf), binary.BigEndian, &value)
	return fmt.Sprint(value)
}
