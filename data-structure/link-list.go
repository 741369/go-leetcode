/**********************************************
** @Des: reverse link list
** @Author: lzw
** @Last Modified time: 2020/5/20 上午10:07
***********************************************/

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
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	list2 := reverseBetween2(list, 0, 4)
	for list2 != nil {
		fmt.Printf("%#v \t", list2.Val)
		list2 = list2.Next
	}
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
		fmt.Println("====", prev.Val)
	}
	return prev
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return nil
	}
	dummy := &ListNode{Next: head}
	d := dummy
	i, j := m, n-m
	for i > 1 {
		d = d.Next
		i--
	}

	var prev *ListNode
	cur := d.Next
	for j >= 0 {
		cur.Next, prev, cur = prev, cur, cur.Next
		j--
	}
	return prev
}
