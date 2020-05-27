/**********************************************
** @Des: 
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/2 下午10:06
***********************************************/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
 	Val int
 	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	curr := reverseList(head.Next)

	next := curr

	for next.Next != nil {
		next = next.Next
	}
	next.Next = head
	head.Next = nil

	return curr
}

func reverseList2(head *ListNode) *ListNode {
	var pre  *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func main() {
	tempList := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}

	revList := reverseList(&tempList)
	fmt.Printf("%#v", revList)

}
/**
     * 迭代方法
     * 1 -> 2 -> 3 -> 4 -> null
     * null <- 1 <- 2 <- 3 <- 4
     *
     * @param head
     * @return
public static ListNode reverseListIterative(ListNode head) {
ListNode prev = null; //前指针节点
ListNode curr = head; //当前指针节点
//每次循环，都将当前节点指向它前面的节点，然后当前节点和前节点后移
while (curr != null) {
ListNode nextTemp = curr.next; //临时节点，暂存当前节点的下一节点，用于后移
curr.next = prev; //将当前节点指向它前面的节点
prev = curr; //前指针后移
curr = nextTemp; //当前指针后移
}
return prev;
}

Line 55: Char 37: cannot use param_1 (type *leetcode.ListNode) as type *ListNode in argument to __helper__ (solution.go)
 */
