package main

import (
	"fmt"
)

func main() {
	/*
	* x1x2x3 | x(x1x2x3)
	* -------+----------
	* 000    | 11
	* 001    | 01
	* 010    | 00
	* 011    | 10
	* 100    | 01
	* 101    | 00
	* 110    | 11
	* 111    | 10
	*
	* t1t2 = S(x3x4x3 xor k1k2k3)
	* u1u2 = x1x2 xor t1t2
	* E(x1x2x3x4, k1k2k3) = x3x4u1u2
	 */

	// input binary
	input := make([]string, 4)
	input[0] = "0100"
	input[1] = "0001"
	input[2] = "0100"
	input[3] = "1000"

	tmpInput := make([]string, 4)

	// key
	key := "001"

	// 3 layer
	for y := 0; y < 3; y++ {
		for i, block := range input {
			// t1t2
			t := tFunc(fmt.Sprintf("%s%s%s", string(block[2]), string(block[3]), string(block[2])), key)
			//fmt.Printf("DEBUG: t1t2=%s\n", t)

			// u1u2
			u := uFunc(fmt.Sprintf("%s%s", string(block[0]), string(block[1])), t)
			//fmt.Printf("DEBUG: u1u2=%s\n", u)

			// E(x1x2x3x4, k1k2k3)
			//fmt.Printf("DEBUG:: E(%s, %s) = %s%s%s\n", input[i], key, string(block[2]), string(block[3]), u)
			final := fmt.Sprintf("%s%s%s", string(block[2]), string(block[3]), u)
			tmpInput[i] = final
		}

		// print layer
		fmt.Printf("LAYER %d: ", y)
		for i := range input {
			fmt.Printf(tmpInput[i])
		}
		fmt.Printf("\n")

		// switch inputs
		for i := range input {
			input[i] = tmpInput[i]
		}
	}
}

func tFunc(block, key string) string {
	// declare t vars
	t := make([]string, 3)

	// loop through comparing
	for i := range block {
		if block[i] != key[i] {
			t[i] = "1"

		} else {
			t[i] = "0"
		}
	}

	// init map based off of sbox data
	var m = map[string]string{
		"000": "11",
		"001": "01",
		"010": "00",
		"011": "10",
		"100": "01",
		"101": "00",
		"110": "11",
		"111": "10",
	}

	final := m[fmt.Sprintf("%s%s%s", t[0], t[1], t[2])]
	return final
}

func uFunc(x, t string) string {
	// declare t vars
	out := make([]string, 3)

	// loop through comparing
	for i := range x {
		if x[i] != t[i] {
			out[i] = "1"

		} else {
			out[i] = "0"
		}
	}

	final := fmt.Sprintf("%s%s%s", out[0], out[1], out[2])
	return final
}
