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
			* WLOWAPELNHNHLEGYSOLDNDWNITUIEEFHMDRIEBYTCWEOHARRUE
			* W     E     L     L     I
			* W    P    N    Y
			* WELLIFICA
			* LLEDTHEWR
			* ONGNUMBER
			* WHYDIDYOU
			* ANSWERTHE
			* PHONE
			*
			* there are 50 characters..
			* we know there will be 9 columns
			* floor function of 50/9 = 5 full rows
			* 50 mod 9 give you 5 meaning there will be 5 letters remaining
		  *
			* we can denote that the first 5 columns will have:
			*   celing of 50/9 aka 6 rows
			* and the remaining will have:
			*   floor of 50/9 aka 6-1 aka 5
	*/

	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	lengthPtr := flag.Int("length", 9, "length of column")
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

	// initialize array
	col := *lengthPtr
	row := int(math.Ceil(float64(len(dat)/col) + 1))
	//fmt.Printf("maxArr is max size [%d][%d]\n", row, col)
	maxArr := make([][]string, row)
	for i := range maxArr {
		maxArr[i] = make([]string, col)
	}

	cCount := 0
	iCount := 0

	// print first longer rows
	for i := 0; i < len(dat)%col; i++ {
		for y := 0; y <= int(math.Ceil(float64(len(dat)/col))); y++ {
			maxArr[y][i] = string(dat[cCount])
			cCount++
		}
		iCount++
	}

	// print second shorter rows
	for i := iCount; i < col; i++ {
		for y := 0; y < int(math.Floor(float64(len(dat)/col))); y++ {
			//fmt.Printf("on maxArr[%d][%d]: dat[%d]: %s\n", y, i, cCount, string(dat[cCount]))
			maxArr[y][i] = string(dat[cCount])
			cCount++
		}
	}

	// print out final array
	for i := 0; i < row; i++ {
		//fmt.Printf("final: %s\n", maxArr[i])
		for y := 0; y < col; y++ {
			fmt.Printf("%s", maxArr[i][y])
		}
	}
	fmt.Printf("\n")
}
