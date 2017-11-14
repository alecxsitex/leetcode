package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a "
	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.Trim(s, " ")
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
