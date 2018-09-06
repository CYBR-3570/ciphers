package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
)

var (
	FunctionA int
	FunctionB int
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	orgAPtr := flag.String("originalA", "T", "original 'A' character")
	orgBPtr := flag.String("originalB", "O", "original 'B' character")
	newAPtr := flag.String("newA", "H", "new 'A' character")
	newBPtr := flag.String("newB", "E", "new 'B' character")
	flag.Parse()

	// define starting letters
	originalA := *orgAPtr
	originalB := *orgBPtr
	newA := *newAPtr
	newB := *newBPtr

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

	// convert to ascii
	oLetA := int(originalA[0] - 65)
	nLetA := int(newA[0] - 65)
	oLetB := int(originalB[0] - 65)
	nLetB := int(newB[0] - 65)

	// loop through all combos

	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			//fmt.Printf("trying: %d * %d + %d = %d mod26\n", a, oLetA, b, nLetA)
			if ((a*oLetA + b) % 26) == (nLetA % 26) {
				//fmt.Printf("match found at: a=%d b=%d\n", a, b)
				// now try all b's for a match

				for i := 0; i < 26; i++ {
					for y := 0; y < 26; y++ {
						if ((i*oLetB + y) % 26) == (nLetB % 26) {
							if (a == i) && (b == y) {
								//fmt.Printf("match found at: a=%d b=%d\n", a, b)
								FunctionA = a
								FunctionB = b
							}
						}
					}
				}

			}
		}
	}

	// print end result
	fmt.Printf("Result: a=%d b=%d\n", FunctionA, FunctionB)

	// encipherment formula
	fmt.Printf("Encipherment function: E(x) = (%dx+%d) mod26\n", FunctionA, FunctionB)

	// decipherment formula
	// calculate modular inverse
	inverse := new(big.Int).ModInverse(big.NewInt(int64(FunctionA)), big.NewInt(26))
	s := fmt.Sprintf("%d", inverse)
	i, _ := strconv.Atoi(s)
	fmt.Printf("Decipherment function: %d(y-%d) mod26\n", i, FunctionB)

	// loop through text string
	for _, x := range dat {
		// convert ascii to 0-26 aka alphabet
		a := x - 65

		// compute for every character
		switch {
		case (a >= 0) && (a <= 26):

			// D(y) = c(y - b) mod26
			b := (i * (int(a) - FunctionB)) % 26
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
