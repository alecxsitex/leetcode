package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 2147483648
	fmt.Println(isPalindrome(num))
}
func isPalindrome(x int) bool {
	numStr := strconv.Itoa(x)
	if x < 0 {
		return false
	}
	var reverse string
	for _, num := range numStr {
		reverse = string(num) + reverse
	}
	if numStr == reverse {
		return true
	}
	return false
}
