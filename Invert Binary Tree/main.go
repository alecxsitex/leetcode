package main

/*
     4
   /   \
  2     7
 / \   / \
1   3 6   9
*/
func main() {
	var ll = &TreeNode{Val: 1}
	var lr = &TreeNode{Val: 3}
	var rl = &TreeNode{Val: 6}
	var rr = &TreeNode{Val: 9}
	var l = &TreeNode{Val: 2, Left: ll, Right: lr}
	var r = &TreeNode{Val: 7, Left: rl, Right: rr}
	var root = &TreeNode{Val: 4, Left: l, Right: r}

	invertTree(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}
