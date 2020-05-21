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
	if head == nil || m > n {
		return nil
	}

	//申请节点指向头节点
	fake := &ListNode{}
	fake.Next = head
	prve := fake

	//走到将要翻转节点的前一个节点 prev
	for i := 0; i < m-1; i++ {
		prve = prve.Next
	}

	//cur 第m个节点，也就是将要翻转的节点
	cur := prve.Next
	for i := m; i < n; i++ {
		cur.Next, cur.Next.Next, prve.Next = cur.Next.Next, prve.Next, cur.Next
		//tmp := cur.Next            //保存要反转节点的下一个节点
		//cur.Next = tmp.Next    //当前节点指向 要放转节点的next节点，最终指向第m个节点的next
		//tmp.Next = prve.Next  //第n个节点的next指向前一个节点
		//prve.Next = tmp           // 第m个节点指向后面一个节点
	}

	return fake.Next
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

	cur := d.Next.Next
	pre := d.Next
	for j > 0 {
		pre.Next = cur.Next
		cur.Next = d.Next
		d.Next = cur
		cur = pre.Next
		//cur.Next, prev, cur = prev, cur, cur.Next
		j--
	}
	return dummy.Next
}
