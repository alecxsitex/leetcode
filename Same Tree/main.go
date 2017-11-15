package main

/**
 * Definition for a binary tree node.
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**Input:     1         1
*          / \       / \
*         2   3     2   3
*
*        [1,2,3],   [1,2,3]
*
*Output: true
*
*Input:     1         1
*          / \       / \
*         2   1     1   2
*
*        [1,2,1],   [1,1,2]
*
*Output: false
*
 **/
func main() {

}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}
