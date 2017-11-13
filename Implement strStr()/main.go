package main

import (
	"fmt"
)

func main() {
	var haystack = ""
	var needle = ""
	fmt.Println(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 || len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < ((len(haystack) - len(needle)) + 1); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1

}
