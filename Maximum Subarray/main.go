package main

import "fmt"

func main() {
	var nums = []int{-1, -2}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, v := range nums[1:] {
		if sum+v > v {
			sum += v
		} else {
			sum = v
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
