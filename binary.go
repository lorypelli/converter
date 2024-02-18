package main

import (
	"fmt"
)

func ToBinary(n int) string {
	var s string
	for i := n; i > 0; i /= 2 {
		s += fmt.Sprint(i % 2)
	}
	return s
}
