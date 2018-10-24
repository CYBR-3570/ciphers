package main

import (
	"fmt"
)

func main() {
	// declare the input sequence

	b4 := 1
	b3 := 0
	b2 := 0
	b1 := 0

	for i := 0; i <= 80; i++ {
		// print the output (b1)
		fmt.Printf("%d", b1)
		// calculate the new b5
		temp := (b3 + b2 + b1) % 2
		b1 = b2
		b2 = b3
		b3 = b4
		b4 = temp
	}
}
