/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/16 下午9:47
***********************************************/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	root *TreeNode
}

// 插入
func (bst *BST) Insert(val int) {
	newNode := &TreeNode{Val: val}

	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}
func insertNode(root, newNode *TreeNode) {
	// 插入节点放在根节点左边
	if newNode.Val < root.Val {
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertNode(root.Left, newNode)
		}
	} else if newNode.Val > root.Val {
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertNode(root.Right, newNode)
		}
	}
	// 否则等于根节点
}

// 删除
func (bst *BST) Remove(val int) bool {
	_, existed := remove(bst.root, val)
	return existed
}
func remove(root *TreeNode, val int) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	var existed bool
	// 从左边找
	if val < root.Val {
		root.Left, existed = remove(root.Left, val)
		return root, existed
	}

	// 从右边找
	if val > root.Val {
		root.Right, existed = remove(root.Right, val)
		return root, existed
	}

	// 如果此节点正是要移除的节点，那么返回此节点，返回之前需要调整
	existed = true

	// 如果此节点没有孩子，直接返回
	if root.Left == nil && root.Right == nil {
		root = nil
		return root, existed
	}

	// 如果左子节点为空，提升右子节点
	if root.Left == nil {
		root = root.Right
		return root, existed
	}

	// 如果右子节点为空，提升左子节点
	if root.Right == nil {
		root = root.Left
		return root, existed
	}

	// 如果左右节点都存在，那么从右边节点找到最小的节点提升，这个节点肯定比左子树所有节点都要大，
	// 也可以从左子树中找到一个最大的提升，道理是一样的

	smallestInRight, _ := min(root.Right)
	// 提升
	root.Val = smallestInRight
	// 从右子树删除此节点
	root.Right, _ = remove(root.Right, smallestInRight)

	return root, existed
}

// 搜索
func (bst *BST) Search(val int) bool {
	return search(bst.root, val)
}
func search(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val < root.Val {
		return search(root.Left, val)
	}
	if val > root.Val {
		return search(root.Right, val)
	}
	return false
}

// Min 二叉树最小值
func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}
func min(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, false
	}

	n := root
	// 从左边找
	for {
		if n.Left == nil {
			return n.Val, true
		}
		n = n.Left
	}
}

// Max 二叉树最大值
func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}
func max(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, false
	}
	// 从右边找
	n := root
	for {
		if n.Right == nil {
			return n.Val, true
		}
		n = n.Right
	}
}

// 遍历
// 前序遍历
func (bst *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(root *TreeNode, f func(int)) {
	if root != nil {
		f(root.Val)
		preOrderTraverse(root.Left, f)
		preOrderTraverse(root.Right, f)
	}
}

// 中序
func (bst *BST) InOrderTraverse(f func(int)) {
	inOrderTraverse(bst.root, f)
}
func inOrderTraverse(root *TreeNode, f func(int)) {
	if root != nil {
		inOrderTraverse(root.Left, f)
		f(root.Val)
		inOrderTraverse(root.Right, f)
	}
}

// 后序
func (bst *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(bst.root, f)
}
func postOrderTraverse(root *TreeNode, f func(int)) {
	if root != nil {
		postOrderTraverse(root.Left, f)
		postOrderTraverse(root.Right, f)
		f(root.Val)
	}
}
