package main

import (
	"fmt"
)

func main() {
	// declare the input sequence

	b6 := 1
	b5 := 1
	b4 := 1
	b3 := 0
	b2 := 0
	b1 := 0

	for i := 0; i <= 200; i++ {
		// print the output (b1)
		fmt.Printf("%d", b1)
		// calculate the new b6
		temp := (b4 + b2) % 2
		b1 = b2
		b2 = b3
		b3 = b4
		b4 = b5
		b5 = b6
		b6 = temp
	}
}
