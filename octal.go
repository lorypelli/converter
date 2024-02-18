package main

import (
	"fmt"
)

func ToOctal(n float64) string {
	var s string
	for i := int(n); i > 0; i /= 8 {
		s += fmt.Sprint(i % 8)
	}
	Reverse(&s)
	return s
}
