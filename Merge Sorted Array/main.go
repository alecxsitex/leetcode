package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3, 5, 0, 0}
	nums2 := []int{2, 4}

	merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	cur1 := m - 1
	cur2 := n - 1
	tail := m + n - 1
	for cur2 >= 0 {
		if cur1 >= 0 && nums1[cur1] > nums2[cur2] {
			nums1[tail] = nums1[cur1]
			cur1--
		} else {
			nums1[tail] = nums2[cur2]
			cur2--
		}
		tail--
	}

}
