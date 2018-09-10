package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	keywordPtr := flag.String("keyword", "GERMAN", "keyword to use")
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
	fmt.Print(string(dat))
	fmt.Println("")

	keyword := *keywordPtr

	println(keyword)
	s := strings.Split(keyword, "")
	sort.Strings(s)
	println(strings.Join(s, ""))

	//  G E R M A N
	//  d e f e n d
	//  t h e e a s
	//  t w a l l o
	//  f t h e c a
	//  s t l e x x

	// create 2d array to hold all data
	row := len(keyword)
	col := (len(dat) / len(keyword)) + 2
	maxArr := make([][]string, row)
	for i := range maxArr {
		maxArr[i] = make([]string, col)
	}
	finArr := make([][]string, row)
	for i := range finArr {
		finArr[i] = make([]string, col)
	}

	// populate array with keyword
	for i, y := range keyword {
		maxArr[i][0] = string(y)
	}

	// populate the array with data
	// init remainder counter
	r := 1
	for i, y := range dat {
		//fmt.Printf("setting [%d][%d] : %s\n", i%len(keyword), r, string(y))
		maxArr[i%len(keyword)][r] = string(y)
		if i%(len(keyword)) == len(keyword)-1 {
			r = r + 1
		}
	}

	finCount := 0
	for _, y := range s {
		// find the row needed
		for i := 0; i < len(keyword); i++ {
			if maxArr[i][0] == y {
				//fmt.Printf("%v\n", maxArr[i])
				// update final array with proper data
				finArr[finCount] = maxArr[i]
				finCount++
			}
		}
	}

	// print out message
	for i := 0; i < col; i++ {
		for y := 1; y < row; y++ {
			fmt.Printf("%s", finArr[i][y%row])
		}
	}
	fmt.Printf("\n")
	// print array
	//fmt.Printf("%v\n", finArr)
}
