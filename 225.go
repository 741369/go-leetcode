/**********************************************
** @Des: 225.用队列实现栈操作
** @Author: liuzhiwang
** @Last Modified time: 2019/9/16 下午4:09
***********************************************/

package main

import "fmt"

type MyStack struct {
	stack []int
}

// 初始化
func Constructor() MyStack {
	return MyStack{
		stack: make([]int, 0),
	}
}

/** Push element x onto stack */
func (myStack *MyStack) Push(x int)  {
	myStack.stack = append(myStack.stack, x)
}

/** Removes the element on top of the stack and returns that element */
func (myStack *MyStack) Pop() int {
	if len(myStack.stack) == 0 {
		return 0
	}
	tmp := myStack.stack[len(myStack.stack)-1]
	myStack.stack = myStack.stack[:len(myStack.stack)-1]
	return tmp
}

/** Get the top element. */
func (myStack *MyStack) Top() int {
	if len(myStack.stack) == 0 {
		return 0
	}
	return myStack.stack[len(myStack.stack)]
}

/** Returns whether the stack is empty. */
func (myStack *MyStack) Empty() bool {
	return len(myStack.stack) == 0
}

func main() {
	obj := Constructor()

	obj.Push(5)

	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()

	fmt.Println(param_2, "==", param_3, "==", param_4)

}