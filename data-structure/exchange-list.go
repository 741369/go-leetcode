/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/20 下午12:10
***********************************************/

package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h := &ListNode{
		Val:  0,
		Next: head,
	}
	p, cur := h, h.Next
	for cur != nil && cur.Next != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		p = cur
		cur = cur.Next
	}

	return h.Next
}
