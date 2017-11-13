package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 3, 5, 6}
	var target = 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); {
		if nums[i] < target {
			i++
		} else {
			return i
		}

	}
	return len(nums)
}
