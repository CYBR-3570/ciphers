package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
	* WASHINGTON
	* 22 0 18 7 8 13 6 19 14 13
	*
	* [3 7] [22] = [3*22 - 7*0] = [66 ]  = [14]
	* [8 9] [0 ]   [8*22 - 9*0]   [176]    [20]
	*
	* or (matrixArgs[0] * dat[i]) - (matrixArgs[1] * dat[i+1])%26
	*    (matrixArgs[2] * dat[i]) - (matrixArgs[3] * dat[i+1])%26
	*
	* 14 20 25 25 11 25 21 11 3 21
	* OUZZLZVLDV
	 */

	// read in runtime args
	matrixArgs := os.Args[1:]

	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	flag.Parse()

	// read in text file
	udat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}
	// strip spaces and newline
	dat := strings.Replace(string(udat), " ", "", -1)
	dat = strings.Replace(string(dat), "\n", "", -1)

	fmt.Println(matrixArgs)
	fmt.Println("")
	fmt.Println(string(dat))
	//fmt.Println("")

	if len(dat)%2 != 0 {
		fmt.Printf("Message needs padding\n")
		os.Exit(0)
	}

	// start doing math on letter pairs
	for i := 0; i < len(dat); i = i + 2 {
		// print pairs
		//println("--------------------")
		//fmt.Printf("ORIG: [%d, %d]\n", dat[i]-65, dat[i+1]-65)
		w, _ := strconv.Atoi(matrixArgs[0])
		x, _ := strconv.Atoi(matrixArgs[1])
		y, _ := strconv.Atoi(matrixArgs[2])
		z, _ := strconv.Atoi(matrixArgs[3])
		//fmt.Printf("DEBUG: (%d*%d + %d*%d)\n", w, int(dat[i]-65), x, int(dat[i+1]-65))
		//fmt.Printf("DEBUG: (%d*%d + %d*%d)\n", y, int(dat[i]-65), z, int(dat[i+1]-65))
		a := ((w * int(dat[i]-65)) + (x * int(dat[i+1]-65))) % 26
		b := ((y * int(dat[i]-65)) + (z * int(dat[i+1]-65))) % 26

		//fmt.Printf("NEW: [%d, %d]\n", a, b)
		//println("--------------------")

		dictionary := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		fmt.Printf("%s", dictionary[a])
		fmt.Printf("%s", dictionary[b])
	}
	fmt.Printf("\n")
}
