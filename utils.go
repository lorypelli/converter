package main

func Reverse(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp
	}
	*s = string(runes)
}