package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	tmp := sortList(node)
	for tmp != nil {
		fmt.Printf("%d\t", tmp.Val)
		tmp = tmp.Next
	}
}

// 二分，这里用快慢指针，二分的时候注意截断，注意细节，为了让slow可以用来截断，初始化时slow慢一个指针
// merge，添加头指针比较容易计算
// 单链表排序归并
func sortList(head *ListNode) *ListNode {
	// 1. 判断边界条件，为空或者只有一个元素直接返回（肯定有序）
	if head == nil || head.Next == nil {
		return head
	}
	// 2. 找链表的中间点
	p1, p2 := head, head.Next
	if p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 3. 排序子链表，注意链表拆分后要截断
	right := sortList(p1.Next)
	p1.Next = nil
	left := sortList(head)

	// 4. 合并排序后的链表
	tmp := &ListNode{}
	return mergeSort(tmp, left, right)
}

func mergeSort(res, left, right *ListNode) *ListNode {
	head := res
	for left != nil && right != nil {
		if left.Val > right.Val {
			head.Next = right
			right = right.Next
		} else {
			head.Next = left
			left = left.Next
		}
		head = head.Next
	}
	if left != nil {
		head.Next = left
	}
	if right != nil {
		head.Next = right
	}
	return res.Next
}

// 单链表快速排序Time Complexity: O(nlogn), Space: O(n)
func sortList2(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head // 返回排序后的链表
}

// 定义两个辅助指针first，second,这两个指针均往next方向移动，移动的过程中保持first之前的值都小于选定的pivot，
// first和second之间的值都大于选定的pivot，那么当second走到末尾时交换first的值与pivot便完成了一次切分
// first 指向的永远是最后一个小于pivot的元素
// 切分之后，first之前的值 <= fist < second
func quickSort(first, last *ListNode) {
	// 如果区间没有元素或者只有一个元素则返回
	if first == last || first.Next == last {
		return
	}
	// 进行一次排序，p 之前的比p小，p之后的比p大
	p := partition(first, last)
	quickSort(first, p)
	quickSort(p.Next, last)
}

func partition(first, last *ListNode) *ListNode {
	// 如果区间没有元素或者只有一个元素则返回
	if first == last || first.Next == last {
		return first
	}

	// first永远指向最后一个比pivot小的元素
	// 双指针，快指针比pivot小就交换快慢指针
	// 因为first之前的比pivot小于等于，first与second之间的都比pivot大
	pivot, root, second := first.Val, first, first.Next
	for second != last {
		if second.Val < pivot {
			first = first.Next
			first.Val, second.Val = second.Val, first.Val
		}
		second = second.Next
	}
	// 	// first表示的是小于pivot的最后一个元素，
	// 交换之后，first之前的比first小，之后的比first大
	first.Val, root.Val = root.Val, first.Val
	return first
}
