package main

import (
	"fmt"
)

func main() {
	var str = "["
	fmt.Println(isValid(str))

}

func isValid(s string) bool {
	var open = map[rune]int{40: 0, 91: 1, 123: 2}
	var close = map[rune]int{41: 0, 93: 1, 125: 2}
	var current []int

	str := []rune(s)

	if len(str) == 0 {
		return true
	}

	for _, char := range str {
		if _, ok := open[char]; ok {
			current = append(current, open[char])
		}

		if _, ok := close[char]; ok {
			if len(current) < 1 || close[char] != current[len(current)-1] {
				return false
			}
			current = current[0 : len(current)-1]
		}
	}

	if len(current) != 0 {
		return false
	}
	return true
}
