package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	tmp := reverse3(list)
	fmt.Printf("%#v\n", tmp)
	for tmp != nil {
		fmt.Printf("%v\t", tmp.Val)
		tmp = tmp.Next
	}
}

// 迭代 O(n) O(1)
func reverse3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var res *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = res
		res = curr
		curr = temp
	}
	return res
}

// 迭代 O(n) O(1)
func reverse2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := head      //左指针
	r := head.Next //右指针
	for r != nil {
		tmp := r.Next //临时
		r.Next = l    //右指针.next 指向左指针
		l = r         //左指针移动
		r = tmp       //右指针移动
	}
	head.Next = nil
	return l
}

// 递归 O(n) O(n)
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}
