package main

import (
	"fmt"
)

func main() {
	newHead := deleteDuplicates(&l1)
	for {
		if newHead == nil {
			break
		}
		fmt.Printf("%+v\n", newHead)
		newHead = newHead.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Test data
var l5 = ListNode{Val: 3, Next: nil}
var l4 = ListNode{Val: 3, Next: &l5}
var l3 = ListNode{Val: 2, Next: &l4}
var l2 = ListNode{Val: 1, Next: &l3}
var l1 = ListNode{Val: 1, Next: &l2}

//  Output: 1 -> 2 -> 3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var next *ListNode
	var cur *ListNode
	next = head
	cur = next.Next

	for {
		if cur == nil {
			next.Next = nil
			break
		}
		if cur.Val > next.Val {
			next.Next = cur
			next = next.Next
		}
		cur = cur.Next

	}
	return head
}
