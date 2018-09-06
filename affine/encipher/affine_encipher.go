package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	magnitudePtr := flag.Int("magnitude", 3, "magnitude of cipher aka the '3' in '(3x+1) MOD 26'")
	shiftPtr := flag.Int("shift", 1, "shift of the cipher aka the '1' in '(3x+1) MOD 26'")
	flag.Parse()

	// read in text file
	dat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// print current file contents
	fmt.Print(string(dat))
	fmt.Println("")

	// set dictionary aka the reverse of the english alphabet
	//dictionary := []string{"Z", "Y", "X", "W", "V", "U", "T", "S", "R", "Q", "P", "O", "N", "M", "L", "K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A"}
	dictionary := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// loop through text string
	for _, x := range dat {
		// convert ascii to 0-26 aka alphabet
		a := x - 65

		// compute for every character
		switch {
		case (a >= 0) && (a <= 26):
			// convert
			b := ((*magnitudePtr * int(a)) + *shiftPtr) % 26
			fmt.Printf("%s", dictionary[b])
			//fmt.Printf("%s", dictionary[a])
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
