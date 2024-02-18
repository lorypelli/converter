package main

import (
	"fmt"
	"strings"
)

func ToOctal(n float64, d int) string {
	var s string
	for i := int(n); i > 0; i /= 8 {
		s += fmt.Sprint(i % 8)
	}
	decimalPart := GetDecimal(n)
	var i int
	s += "."
	for i < d || decimalPart != 0 {
		decimal := decimalPart * 8
		s += strings.Split(fmt.Sprint(decimal), ".")[0]
		decimalPart = 7 - decimal
		if (decimalPart == 7) {
			decimalPart = 0
		}
		i++
	}
	Reverse(&strings.Split(s, ".")[0])
	return strings.Split(s, ".")[0] + "." + strings.Split(s, ".")[1]
}
