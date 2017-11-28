package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(27))
	//fmt.Println(rune('A'))
}

func convertToTitle(n int) string {
	var result = ""

	for n > 0 {
		n--
		result = string(rune(65+n%26)) + result
		n /= 26
	}
	return result
}
