package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("BB"))
}

func titleToNumber(s string) int {
	// A - rune(65)
	// Z - rune(90)
	var result int
	var inc = 1
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		result += (int(runes[i]) - 64) * inc
		inc *= 26
	}
	return result
}
