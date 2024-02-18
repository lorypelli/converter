package main

import (
	"fmt"
)

func ToHex(n float64) string {
	var s string
	for i := int(n); i > 0; i /= 16 {
		s += fmt.Sprint(i % 16)
	}
	Reverse(&s)
	return s
}
