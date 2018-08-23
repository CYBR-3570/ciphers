package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text4", "relative path to text file")
	modPtr := flag.Int("mod", 6, "modulo: the amount of lines to wrap around cylinder")

	flag.Parse()

	// read in text file
	dat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// print current file contents
	fmt.Print(string(dat))
	fmt.Println("")

	// set up counter variables
	mod := *modPtr
	cur := 0
	counter := 0
	ba := []byte(dat)
	sz := len(ba)

	// iderate through all characters
	for count := 0; count < sz-1; count++ {
		fmt.Printf("%c", ba[cur])

		// make sure we dont outrun array
		if (cur + mod) < sz-1 {
			cur = cur + mod
		} else {
			counter++
			cur = counter
		}
	}
	fmt.Printf("\n")
}
