package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := makeNewTree(1)
	root.Left = makeNewTree(2)
	root.Right = makeNewTree(2)
	root.Right.Right = makeNewTree(3)
	root.Right.Right.Right = makeNewTree(4)
	root.Left.Left = makeNewTree(3)
	root.Left.Left.Left = makeNewTree(4)
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	l := float64(calcDeep(root.Left))
	r := float64(calcDeep(root.Right))
	if math.Abs(l-r) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(calcDeep(root.Left), calcDeep(root.Right))
}

func makeNewTree(val int) *TreeNode {
	var t = &TreeNode{}
	t.Val = val
	return t
}
