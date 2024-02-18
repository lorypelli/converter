package main

import (
	"fmt"
)

func ToHex(n int) string {
	var s string
	for i := n; i > 0; i /= 16 {
		s += fmt.Sprint(i % 16)
	}
	Reverse(&s)
	return s
}
