package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 3, 5}
	var val = 3
	fmt.Println(removeElement(nums, val))
}
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i = 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
