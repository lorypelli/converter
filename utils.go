package main

import (
	"math"
)

func Reverse(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp
	}
	*s = string(runes)
}

func GetBase(b int, n float64, d int) string {
	switch b {
	case 2:
		{
			return ToBinary(n, d)
		}
	case 8:
		{
			return ToOctal(n, d)
		}
	case 16:
		{
			return ToHex(n, d)
		}
	}
	return ""
}

func GetDecimal(num float64) float64 {
	return num - math.Floor(num)
}
