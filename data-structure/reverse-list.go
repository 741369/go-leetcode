/**********************************************
** @Des: 反转链表
** @Author: lzw
** @Last Modified time: 2020/5/21 下午4:42
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
	list2 := reverseList(list)
	for list2 != nil {
		fmt.Printf("%#v \t", list2.Val)
		list2 = list2.Next
	}
	fmt.Println("")
}

// 递归倒序链表
func reverseList2(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	p := reverseList2(node.Next)
	node.Next.Next = node // node.Next 即为p，将node放到p的后面
	node.Next = nil
	return p
}

func reverseList(node *ListNode) *ListNode {
	cur := node
	if cur == nil {
		return nil
	}

	var prev *ListNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
