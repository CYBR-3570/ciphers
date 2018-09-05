package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	magnitudePtr := flag.Int("magnitude", 11, "magnitude of cipher aka the '11' in '(11x+8) MOD 26'")
	shiftPtr := flag.Int("shift", 8, "shift of the cipher aka the '8' in '(11x+8) MOD 26'")
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
	dictionary := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// loop through text string
	for _, x := range dat {
		// convert ascii to 0-26 aka alphabet
		a := x - 65

		// calculate modular inverse
		inverse := new(big.Int).ModInverse(big.NewInt(int64(*magnitudePtr)), big.NewInt(26))
		s := fmt.Sprintf("%d", inverse)
		i, _ := strconv.Atoi(s)

		// compute for every character
		switch {
		case (a >= 0) && (a <= 26):

			// D(y) = c(y - b) mod26
			b := (i * (int(a) - *shiftPtr)) % 26
			// make the mod remainder if negative
			if b < 0 {
				b = b + 26
			}
			fmt.Printf("%s", dictionary[b])
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
