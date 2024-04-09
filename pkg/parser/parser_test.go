package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	m := NewMessage("test")
	var result byte
	fmt.Println(result)
	m.AddPID(1234).AddQR(1).AddOpCode(0).AddAA(0).AddTC(0).AddRD(0).AddRA(0).AddZ(0).AddRcode(0)
	fmt.Println(m.Header)
}

/* 1 0 0 0 1 0 0 1 */
/* 1 0 0 1 */
/* 1 1 0 0 1 0 0 0*/
