package main

import (
	"fmt"
)

func main() {
	// declare the input sequence

	b5 := 1
	b4 := 0
	b3 := 1
	b2 := 0
	b1 := 1

	for i := 0; i <= 80; i++ {
		// print the output (b1)
		fmt.Printf("%d", b1)
		// calculate the new b5
		temp := (b4 + b1) % 2
		b1 = b2
		b2 = b3
		b3 = b4
		b4 = b5
		b5 = temp
	}
}
