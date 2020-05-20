/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/19 下午9:47
***********************************************/

package main

type Node struct {
	Next *Node
	Val  int
}

type MyStack struct {
	top *Node
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	node := &Node{Next: s.top, Val: x}
	s.top = node
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	if s.top == nil {
		return -1
	}
	r := s.top.Val
	s.top.Next, s.top = nil, s.top.Next // 这里注意把pop的节点的Next置为nil，防止内存泄露
	return r
}

/** Get the top element. */
func (s *MyStack) Top() int {
	if s.top == nil {
		return -1
	}
	return s.top.Val
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return s.top == nil
}
