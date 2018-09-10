package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text2", "relative path to text file")
	keywordPtr := flag.String("keyword", "PRIME MINISTER", "keyword to use")
	flag.Parse()

	// read in text file
	udat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// strip spaces and newline
	dat := strings.Replace(string(udat), " ", "", -1)
	dat = strings.Replace(string(dat), "\n", "", -1)
	keyword := strings.Replace(*keywordPtr, " ", "", -1)

	// remove duplicates from keyword
	keyword = removeDuplicates(keyword)
	//fmt.Printf("original keyword:  %s\n", *keywordPtr)
	//fmt.Printf("new keyword:       %s\n", keyword)

	// create new alphabet
	fAlpha := fmt.Sprintf("%sABCDEFGHIJKLMNOPQRSTUVWXYZ", keyword)
	alpha := removeDuplicates(fAlpha)
	fmt.Printf("original alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ\n")
	//fmt.Printf("mid alphabet:      %s\n", alpha)

	// PRIME MINISTER

	// PRIMENST
	// ABCDFGHJ
	// KLOQUVWX
	// YZ

	// create 2d array to hold alphabet
	row := len(keyword)
	col := ((row + 26) % row) + 2
	maxArr := make([][]string, row)
	for i := range maxArr {
		maxArr[i] = make([]string, col)
	}
	finArr := make([][]string, row)
	for i := range finArr {
		finArr[i] = make([]string, col)
	}

	// populate array with alphabet
	r := 0
	for i, y := range alpha {
		maxArr[i%row][r] = string(y)
		if i%row == row-1 {
			r = r + 1
		}
	}
	fmt.Printf("new alphabet:      ")
	ac := 0

	var newAlpha [26]string
	actualAlpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// store alphabet to slice
	for i := 0; i < row; i++ {
		for y := 0; y < col; y++ {
			if maxArr[i][y%row] != "" {
				fmt.Printf("%s", maxArr[i][y%row])
				newAlpha[ac] = maxArr[i][y%row]
				ac++
			}
		}
	}
	fmt.Printf("\n")

	// print current file contents
	fmt.Println("")
	fmt.Printf("original message:  %s\n", string(dat))
	fmt.Printf("new message:       ")

	// run substitution
	for i := range dat {
		// search for character in newAlph
		for j := range newAlpha {
			if newAlpha[j] == string(dat[i]) {
				// print translated letter
				fmt.Printf("%s", actualAlpha[j])
			}
		}
	}
	fmt.Printf("\n")
}

func removeDuplicates(l string) string {
	//yeet := strings.Split(*keywordPtr, "")
	duplicate := strings.Split(l, "")
	cleaned := []string{}
	for _, value := range duplicate {
		if !stringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}

	return strings.Join(cleaned, "")
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
