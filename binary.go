package main

import (
	"fmt"
)

func ToBinary(n float64) string {
	var s string
	for i := int(n); i > 0; i /= 2 {
		s += fmt.Sprint(i % 2)
	}
	Reverse(&s)
	return s
}
