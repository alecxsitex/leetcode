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
	sum := -5

	var root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Left.Left = &TreeNode{Val: 5}

	var r = &TreeNode{Val: -2}
	r.Right = &TreeNode{Val: -3}
	fmt.Println(hasPathSum(r, sum))
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var counter = 0
	if helper(root, counter, sum) {
		return true
	}

	return false

}

func helper(root *TreeNode, counter int, sum int) bool {
	if root == nil {
		return false
	}
	counter += root.Val
	if root.Right == nil && root.Left == nil {
		if counter == sum {
			return true
		}
	}
	if root.Left != nil {
		if helper(root.Left, counter, sum) {
			return true
		}
	}

	if root.Right != nil {
		if helper(root.Right, counter, sum) {
			return true
		}
	}

	return false
}
