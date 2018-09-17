package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// read in runtime flags
	filePtr := flag.String("file", "text1", "relative path to text file")
	keyPtr := flag.String("key", "PRUDENT", "key to encode/decode")
	encipherPtr := flag.Bool("encipher", true, "weather to encipher or not")
	decipherPtr := flag.Bool("decipher", false, "weather to decipher or not")
	flag.Parse()

	// read in text file
	dat, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("could not read in input text file")
	}

	// print current file contents
	fmt.Print(string(dat))
	fmt.Println("")

	if *encipherPtr {
		encoded := encipher(string(dat), *keyPtr)
		fmt.Printf("enciphered message:\n%s\n", encoded)
	}

	if *decipherPtr {
		decoded := decipher(string(dat), *keyPtr)
		fmt.Printf("deciphered message:\n%s\n", decoded)
	}
}

func sanitize(in string) string {
	out := []rune{}
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func quartets(in string) string {
	out := make([]rune, 0, len(in))
	for i, v := range in {
		if i%4 == 0 && i != 0 {
			out = append(out, rune(32))
		}
		out = append(out, v)
	}
	return string(out)
}

func encodePair(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func decodePair(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func encipher(msg, key string) string {
	smsg, skey := sanitize(msg), sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, encodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func decipher(msg, key string) string {
	smsg, skey := sanitize(msg), sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, decodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}
