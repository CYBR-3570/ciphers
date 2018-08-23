package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text7", "relative path to text file")
	keyPtr := flag.String("key", "ITALI", "key to use to advance cipher")
	shiftPtr := flag.Int("shift", 8, "amount of space between the cipher disk shift")

	flag.Parse()

	// read in text file
	dat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// print current file contents after removing spaces
	stripDat := strings.Replace(string(dat), " ", "", -1)
	fmt.Printf("%s/n", stripDat)

	// set up rings
	// inner ring is ciphertext
	// outer ring is plaintext
	innerRing := []string{"k", "v", "p", "&", "m", "r", "d", "l", "g", "a", "z", "e", "n", "b", "o", "s", "f", "c", "h", "t", "y", "q", "i", "x"}
	outerRing := []string{"A", "B", "C", "D", "E", "F", "G", "I", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "V", "X", "Z", "1", "2", "3", "4"}

	// set up counter variables
	key := []byte(*keyPtr)
	shift := *shiftPtr
	keyCount := 0
	shiftCount := 0
	ba := []byte(stripDat)
	sz := len(ba)

	// iderate through all characters
	for count := 0; count < sz-1; count++ {
		// start with key
		for b := 0; b < len(key); b++ {
			fmt.Printf("%c", key[b])
		}
	}
	fmt.Printf("\n")
}
