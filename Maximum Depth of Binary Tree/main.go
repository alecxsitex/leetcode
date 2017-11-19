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

func maxDepth(root *TreeNode) int {
	var depth = 0
	if root == nil {
		return depth
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return 1 + r

}

var mq3 = TreeNode{Val: 4, Left: nil, Right: nil}
var mp3 = TreeNode{Val: 5, Left: nil, Right: nil}
var nq3 = TreeNode{Val: 5, Left: nil, Right: nil}
var m3 = TreeNode{Val: 3, Left: &mq3, Right: &mp3}
var n3 = TreeNode{Val: 3, Left: &nq3, Right: nil}
var root3 = TreeNode{Val: 1, Left: &m3, Right: &n3}

var npq = TreeNode{Val: 8, Left: nil, Right: nil}
var npp = TreeNode{Val: 9, Left: nil, Right: nil}

// var mqq = TreeNode{Val: 8, Left: nil, Right: nil}
// var mqp = TreeNode{Val: 9, Left: nil, Right: nil}
var pm = TreeNode{Val: 3, Left: nil, Right: nil}
var qm = TreeNode{Val: 4, Left: nil, Right: nil}
var pn = TreeNode{Val: 4, Left: &npp, Right: &npq}
var qn = TreeNode{Val: 3, Left: nil, Right: nil}
var m = TreeNode{Val: 2, Left: &pm, Right: &qm}
var n = TreeNode{Val: 2, Left: &pn, Right: &qn}
var root = TreeNode{Val: 1, Left: &m, Right: &n}

func main() {
	fmt.Println(maxDepth(&root))
}
