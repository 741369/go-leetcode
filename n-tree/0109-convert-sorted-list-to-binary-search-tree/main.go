/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/18 下午8:22
***********************************************/

package main

import "fmt"

//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 深度优先搜索 链表


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &ListNode{
		Val:  -10,
		Next: &ListNode{
			Val:  -3,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 9},
				},
			},
		},
	}
	node := sortedListToBST2(head)
	dfs(node)
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	fmt.Printf("%d\t", node.Val)
	dfs(node.Right)
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{
			Val:   head.Val,
			Left:  nil,
			Right: nil,
		}
	}

	middleNode, preNode := middleNodeAndPreNode(head)
	if middleNode == nil {
		return nil
	}
	if preNode != nil {
		preNode.Next = nil
	}
	if middleNode == head {
		head = nil
	}
	return &TreeNode{
		Val:   middleNode.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(middleNode.Next),
	}
}

func middleNodeAndPreNode(head *ListNode) (middle *ListNode, pre *ListNode) {
	if head == nil || head.Next == nil {
		return nil, head
	}

	p1, p2 := head, head
	// 找中节点以及中节点前面的一个节点
	for p2.Next != nil && p2.Next.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1, pre
}

//leetcode submit region end(Prohibit modification and deletion)

// 方法2 分治
// 在找出了中位数节点之后，我们将其作为当前根节点的元素，并递归地构造其左侧部分的链表对应的左子树，以及右侧部分的链表对应的右子树。
func sortedListToBST2(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := getMedian(left, right)
	root := &TreeNode{
		Val:   mid.Val,
		Left:  buildTree(left, mid),
		Right: buildTree(mid.Next, right),
	}
	return root
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}