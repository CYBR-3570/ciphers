package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	/*
	* ECIAO EALOE OBATB RULDL NHSOK DNEUY AE
	* E C I A O _ E A L O E _ O B A T B _ R
	*  U L D L _ N H S O K _ D N E U Y _ A E
	 */

	// read in flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	//railPtr := flag.Int("rail_key", 2, "amount to shift")
	flag.Parse()

	// read in text file
	udat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// remove newline char from data
	dat := strings.Replace(string(udat), "\n", "", -1)

	// print current file contents
	fmt.Println(string(dat))
	//fmt.Printf("legnth of dat: %d\n", len(dat))

	// calculate ceiling function of half for the top half
	cHalf := int(math.Ceil(float64(len(dat) / 2)))
	//fmt.Printf("cHalf is: %d\n", cHalf)

	for i := 0; i < cHalf; i++ {
		// dont print spaces
		if string(dat[i]) != " " {
			fmt.Printf("%s", string(dat[i]))
		}
		if string(dat[(i+cHalf)]) != " " {
			fmt.Printf("%s", string(dat[(i+cHalf)]))
		}
	}
	fmt.Printf("\n")
}
