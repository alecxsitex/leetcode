package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	var count = 0
	var num int

	for _, v := range nums {
		if count == 0 {
			num = v
		}

		if v == num {
			count++
		} else {
			count--
		}
	}
	return num
}
