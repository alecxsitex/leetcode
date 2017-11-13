package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(nums)
	fmt.Println(target)
}

func twoSum(nums []int, target int) []int {
	var result = make([]int, 2)
	var diff int
	m := make(map[int]int)
	for key, num := range nums {
		diff = target - num
		if v, ok := m[diff]; ok {
			result[0] = v
			result[1] = key
		}
		m[num] = key
	}
	return result
}
