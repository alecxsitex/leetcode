package main

import (
	"fmt"
)

func main() {
	var num = []int{0}
	fmt.Println(plusOne(num))
}

func plusOne(digits []int) []int {
	var l = len(digits)
	digits[l-1]++

	for i := l - 1; i >= 0; i-- {
		if digits[i] < 10 {
			break
		} else {
			if i-1 < 0 {
				digits[i] = 1
				digits = append(digits, 0)
			} else {
				digits[i-1]++
				digits[i] = 0
			}
		}
	}

	return digits
}
