package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	l1 := fake([]int{5, 6, 7})
	l2 := fake([]int{5, 5, 5})
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := fake([]int{5, 6, 7})
	l2 := fake([]int{5, 5, 5})
	fmt.Printf("add two nums with list, %v + %v = ", linklistToValue(l1), linklistToValue(l2))
	nums := addTwoNumbers(l1, l2)
	fmt.Printf("%v\n", linklistToValue(nums))

	fmt.Printf("add two nums2 with list, %v + %v = ", linklistToValue(l1), linklistToValue(l2))
	nums2 := addTwoNumbers2(l1, l2)
	fmt.Printf("%v\n", linklistToValue(nums2))
}

func fake(list []int) *ListNode {
	l := &ListNode{}
	o := l
	for k, v := range list {
		o.Val = v
		if len(list) > k+1 {
			o.Next = &ListNode{}
			o = o.Next
		}
	}
	return l
}

func linklistToValue(root *ListNode) (res int) {
	count := 1
	for root != nil {
		res += root.Val * count
		count *= 10
		root = root.Next
	}
	return res
}
