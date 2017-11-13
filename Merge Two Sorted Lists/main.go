package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var start *ListNode

	if l1.Val < l2.Val {
		start = l1
		start.Next = mergeTwoLists(l1.Next, l2)
	} else {
		start = l2
		start.Next = mergeTwoLists(l2.Next, l1)
	}
	return start
}
