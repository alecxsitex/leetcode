package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	levelSlice(root, &result, 0)
	var sorted [][]int
	for i := len(result) - 1; i >= 0; i-- {
		sorted = append(sorted, result[i])
	}
	return sorted

}

func levelSlice(root *TreeNode, list *[][]int, level int) {
	if root == nil {
		return
	}
	if level >= len(*list) {
		var newlist []int
		*list = append(*list, newlist)
	}

	levelSlice(root.Left, list, level+1)
	levelSlice(root.Right, list, level+1)
	(*list)[level] = append((*list)[level], root.Val)

}

/**
    3
   / \
  9  20
    /  \
   15   7
  /  \   \
 21  22   23
**/
var rrr = TreeNode{Val: 23, Left: nil, Right: nil}
var rlr = TreeNode{Val: 22, Left: nil, Right: nil}
var rll = TreeNode{Val: 21, Left: nil, Right: nil}
var rr = TreeNode{Val: 7, Left: nil, Right: &rrr}
var rl = TreeNode{Val: 15, Left: &rll, Right: &rlr}
var r = TreeNode{Val: 20, Left: &rl, Right: &rr}
var l = TreeNode{Val: 9, Left: nil, Right: nil}
var root = TreeNode{Val: 3, Left: &l, Right: &r}
var exp = [][]int{{15, 7}, {9, 20}, {3}}

func main() {
	fmt.Println(levelOrderBottom(&root))
}
