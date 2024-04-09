package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	m := NewMessage("test")
	var result byte
	fmt.Println(result)
	m.AddDefaultHeader()
	fmt.Println(m.Header)
}

/* 1 0 0 0 1 0 0 1 */
/* 1 0 0 1 */
/* 1 1 0 0 1 0 0 0*/
