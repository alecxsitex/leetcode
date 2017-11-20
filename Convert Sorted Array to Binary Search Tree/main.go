package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	halfLenght := len(nums) / 2
	var root = TreeNode{Val: nums[halfLenght]}
	root.Left = sortedArrayToBST(nums[:halfLenght])
	root.Right = sortedArrayToBST(nums[halfLenght+1 : len(nums)])
	return &root
}
