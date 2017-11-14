package main

import (
	"fmt"
)

func main() {
	//a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	//b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	//exp := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"
	a := "11"
	b := "1"
	exp := "100"
	fmt.Println(addBinary(a, b))
	fmt.Println(exp)
}

func addBinary(a string, b string) string {
	var adigit, bdigit, carry int
	var apos = len(a) - 1
	var bpos = len(b) - 1
	var result string

	for {
		if apos >= 0 || bpos >= 0 || carry == 1 {
			adigit = 0
			bdigit = 0

			if apos >= 0 {
				if a[apos] == []byte("1")[0] {
					adigit = 1
					apos--
				}
			}
			if bpos >= 0 {
				if b[bpos] == []byte("1")[0] {
					bdigit = 1
					bpos--
				}
			}

			if adigit+bdigit+carry >= 2 {
				carry = 1
			} else {
				carry = 0
			}
			result = string((adigit+bdigit+carry)%2) + result
		} else {
			break
		}
	}
	return result
}
