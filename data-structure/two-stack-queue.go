/**********************************************
** @Des: 2个栈实现一个队列
** @Author: lzw
** @Last Modified time: 2020/5/19 下午9:32
***********************************************/

package main

type MyQueue struct {
	stackPush []int
	stackPop  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

// 把全部数据压入stackPop栈内
func (queue *MyQueue) pushToPop() {
	if len(queue.stackPop) <= 0 {
		for _, val := range queue.stackPush {
			queue.stackPop = append(queue.stackPop, val)
		}
		// 清空stackPush
		queue.stackPush = nil
	}
}

/** Push element x to the back of queue. */
func (queue *MyQueue) Push(x int) {
	queue.stackPush = append(queue.stackPush, x)
	queue.pushToPop()
}

/** Removes the element from in front of queue and returns that element. */
func (queue *MyQueue) Pop() int {
	if len(queue.stackPop) <= 0 && len(queue.stackPush) <= 0 {
		return 0
	}
	queue.pushToPop()
	ret := queue.stackPop[0]
	queue.stackPop = queue.stackPop[1:]
	return ret
}

/** Get the front element. */
func (queue *MyQueue) Peek() int {
	if len(queue.stackPop) <= 0 && len(queue.stackPush) <= 0 {
		return 0
	}
	queue.pushToPop()
	return queue.stackPop[0]
}

/** Returns whether the queue is empty. */
func (queue *MyQueue) Empty() bool {
	return len(queue.stackPop) <= 0 && len(queue.stackPush) <= 0
}
