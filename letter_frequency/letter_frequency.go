package main

import (
	//"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read in flags
	//shiftPtr := flag.Int("shift", 0, "amount to shift")
	//flag.Parse()
	//shift := *shiftPtr

	// read in runtime args
	upperArgs := strings.ToUpper(os.Args[1])
	byteArgs := []byte(upperArgs)

	// print unshifted data
	fmt.Println(upperArgs)
	fmt.Println(byteArgs)
	fmt.Println("")

	// set dictionary aka the reverse of the english alphabet
	dictionary := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//freq := []int{}
	var freq [26]int

	// loop through text string
	for _, x := range byteArgs {
		// convert ascii to 0-26 aka alphabet
		a := x - 65

		// add letter to int array
		curLetterCount := freq[a]
		newLetterCount := curLetterCount + 1
		freq[a] = newLetterCount

		//fmt.Printf("found %s. it is now at frequency %d\n", dictionary[a], newLetterCount)
	}

	// loop through and print out frequencies
	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			fmt.Printf("%s: %d\n", dictionary[i], freq[i])
		}
	}
}
