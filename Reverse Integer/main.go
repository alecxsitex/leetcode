package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 123
	fmt.Println(reverse(num))
}
func reverse(x int) int {
	var minus = false
	var resultStr string
	num := strconv.Itoa(x)
	if x < 0 {
		minus = true
	}
	if minus {
		num = num[1:len(num)]
	}

	for _, v := range num {
		resultStr = string(v) + resultStr
	}

	result, err := strconv.ParseInt(resultStr, 10, 32)
	if err != nil {
		return 0
	}
	if minus {
		result = -result
	}
	return int(result)
}
