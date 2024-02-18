package main

import (
	"fmt"
)

func ToOctal(n int) string {
	var s string
	for i := n; i > 0; i /= 8 {
		s += fmt.Sprint(i % 8)
	}
	Reverse(&s)
	return s
}
