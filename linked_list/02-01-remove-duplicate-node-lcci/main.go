/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/26 下午5:20
***********************************************/

package main

import "fmt"

//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	tmp := removeDuplicateNodes3(node)
	for tmp != nil {
		fmt.Printf("%d\t", tmp.Val)
		tmp = tmp.Next
	}
}

// 一次hash
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	valMap := map[int]bool{
		head.Val: true,
	}
	node := head
	for node.Next != nil {
		if valMap[node.Next.Val] {
			// 删除node.Next节点
			node.Next = node.Next.Next
		} else {
			// 记录hash
			valMap[node.Next.Val] = true
			// 往后移动一个节点
			node = node.Next
		}
	}
	node.Next = nil
	return head
}

// 两次循环
func removeDuplicateNodes3(head *ListNode) *ListNode {
	oneNode := head
	for oneNode != nil {
		twoNode := oneNode
		for twoNode.Next != nil {
			if oneNode.Val == twoNode.Next.Val {
				// 删除twoNode
				twoNode.Next = twoNode.Next.Next
			} else {
				// twoNode 后移动
				twoNode = twoNode.Next
			}
		}
		oneNode = oneNode.Next
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	occurred := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}
