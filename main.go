package main

import (
	"fmt"
)

func main() {
	var n float64
	var b, d int
	fmt.Print("Provide a number ")
	_, err := fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Don't use negative numbers, make it simple")
	}
	for n < 0 || err != nil {
		fmt.Print("Provide a number ")
		_, err = fmt.Scan(&n)
		if n < 0 {
			fmt.Println("Don't use negative numbers, make it simple")
		}
	}
	fmt.Print("Provide a base ")
	fmt.Scan(&b)
	if b != 2 && b != 8 && b != 16 {
		fmt.Printf("%d is not a valid base, it can only be 2 or 8 or 16\n", b)
	}
	for b != 2 && b != 8 && b != 16 {
		fmt.Print("Provide a base ")
		fmt.Scan(&b)
		if b != 2 && b != 8 && b != 16 {
			fmt.Printf("%d is not a valid base, it can only be 2 or 8 or 16\n", b)
		}
	}
	fmt.Print("Provide number of decimal digits (default 6) ")
	_, err = fmt.Scan(&d)
	if d < 0 || err != nil {
		fmt.Println("Using default...")
		d = 6
	}
	fmt.Printf("The number %.*f in base %d is %s", d, n, b, GetBase(b, n, d))
}
