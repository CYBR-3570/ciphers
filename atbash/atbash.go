package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read in text file
	dat, err := ioutil.ReadFile("text")
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	fmt.Print(string(dat))
	dictionary := []string{"Z", "Y", "X", "W", "V", "U", "T", "S", "R", "Q", "P", "O", "N", "M", "L", "K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A"}

	for _, x := range dat {
		a := x - 65
		switch {
		case (a >= 0) && (a <= 26):
			// convert
			fmt.Printf("%s", dictionary[a])
		case a == 223:
			// space character
			fmt.Printf(" ")
		case a == 201:
			// end of line
			fmt.Printf("\n")
		default:
		}
	}
}
