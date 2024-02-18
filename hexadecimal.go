package main

import (
	"fmt"
	"strings"
)

func ToHex(n float64, d int) string {
	var s string
	for i := int(n); i > 0; i /= 16 {
		s += fmt.Sprint(i % 16)
	}
	decimalPart := GetDecimal(n)
	var i int
	s += "."
	for i < d || decimalPart != 0 {
		decimal := decimalPart * 16
		s += strings.Split(fmt.Sprint(decimal), ".")[0]
		decimalPart = 15 - decimal
		if (decimalPart == 15) {
			decimalPart = 0
		}
		i++
	}
	Reverse(&strings.Split(s, ".")[0])
	return strings.Split(s, ".")[0] + "." + strings.Split(s, ".")[1]
}
