package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	lengthPtr := flag.Int("length", 5, "length of column")
	padPtr := flag.Bool("pad", true, "whether or not to pad")
	padLetterPtr := flag.String("pad_letter", "T", "letter to pad with")
	flag.Parse()

	// read in text file
	udat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// strip spaces and newline
	dat := strings.Replace(string(udat), " ", "", -1)
	dat = strings.Replace(string(dat), "\n", "", -1)

	// print current file contents
	fmt.Println(string(dat))

	col := *lengthPtr
	row := int(math.Ceil(float64(len(dat)/col) + 1))
	//fmt.Printf("maxArr is max size [%d][%d]\n", row, col)
	maxArr := make([][]string, row)
	for i := range maxArr {
		maxArr[i] = make([]string, col)
	}

	rowCt := 0
	colCt := 0
	// start putting chars in array
	for _, i := range dat {
		//fmt.Printf("putting %s in [%d][%d]\n", string(i), rowCt, colCt%col)
		maxArr[rowCt][colCt%col] = string(i)
		colCt++
		if colCt%col == 0 {
			rowCt++
		}
	}

	// fill spaces with padding if set
	if *padPtr {
		for i := 0; i < col; i++ {
			if maxArr[row-1][i] == "" {
				maxArr[row-1][i] = *padLetterPtr
			}
		}
	}

	// print out final array
	//for i := 0; i < row; i++ {
	//	fmt.Printf("final: %s\n", maxArr[i])
	//}

	// print out string
	for i := 0; i < col; i++ {
		for y := 0; y < row; y++ {
			fmt.Printf("%s", maxArr[y][i])
		}
	}
	fmt.Printf("\n")
}
