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
	integerPart := strings.Split(s, ".")[0]
	Reverse(&integerPart)
	if d > 0 {
		return integerPart + "." + strings.Split(s, ".")[1]
	}
	return integerPart
}
