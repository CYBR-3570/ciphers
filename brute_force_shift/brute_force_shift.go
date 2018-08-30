package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read in runtime args
	upperArgs := strings.ToUpper(os.Args[1])
	byteArgs := []byte(upperArgs)

	// print unshifted data
	fmt.Println(upperArgs)
	fmt.Println(byteArgs)
	fmt.Println("")

	// shift 26 times aka english alphabet
	for i := 0; i < 26; i++ {
		fmt.Printf("shift %d: ", i)
		shift(byteArgs, i)
	}
}

func shift(byteArgs []byte, shift int) {
	// set dictionary aka the reverse of the english alphabet
	dictionary := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// loop through text string
	for _, x := range byteArgs {
		// convert ascii to 0-26 aka alphabet
		a := x - 65

		// compute modulo around alphabet
		//m := (a + 6) % 26
		m := (int(a) + shift) % 26

		// compute for every character
		//fmt.Printf("%d:%s ", m, dictionary[m])
		fmt.Printf("%s", dictionary[m])
	}
	fmt.Println("")
}
