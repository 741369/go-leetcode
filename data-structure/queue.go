/**********************************************
** @Des: 循环队列
** @Author: lzw
** @Last Modified time: 2020/5/19 下午8:57
***********************************************/

package main

/**
先说说循环队列的结构体成员设计，队列结构体中除了用于存储元素的数组，还有用于标识队首元素下标的front、用于标识下一个元素入队下标的rear以及数组的最大长度，并且rear独占一个空间。也就是说，若队列最多存储k个元素，则数组长度为k+1。私以为，只要解决了循环队列的判空与判满的规则设计问题，其他的都迎刃而解。在本答案中，判空与判满规则如下：
判空：front == rear
判满：rear + 1 == front，要注意数组越界问题的转换
几个需要注意的点：
front或rear的数组越界转换问题：在移动过程中，若是越界则需要转换到有效下标值处
访问队尾元素时的越界问题：队尾元素的下标是rear - 1，因此若是rear=0，则需要进行有效下标转换
*/
type MyCircularQueue struct {
	data        []int
	front, rear int // front标识出队元素的下标，read标识队的下标
	maxSize     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:    make([]int, k+1, k+1),
		front:   0,
		rear:    0,
		maxSize: k + 1, // 多加一个用于给rear指针站位
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (queue *MyCircularQueue) EnQueue(value int) bool {
	if queue.IsFull() {
		return false
	}
	queue.data[queue.rear] = value
	queue.rear++
	if queue.rear == queue.maxSize {
		queue.rear = 0
	}

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (queue *MyCircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}
	queue.front++
	if queue.front == queue.maxSize {
		queue.front = 0
	}

	return true
}

/** Get the front item from the queue. */
func (queue *MyCircularQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.data[queue.front]
}

/** Get the last item from the queue. */
func (queue *MyCircularQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}
	lastIndex := queue.rear - 1
	// 注意边界问题
	if lastIndex < 0 {
		lastIndex = queue.maxSize - 1
	}
	return queue.data[lastIndex]
}

/** Checks whether the circular queue is empty or not. */
func (queue *MyCircularQueue) IsEmpty() bool {
	return queue.front == queue.rear
}

/** Checks whether the circular queue is full or not. */
func (queue *MyCircularQueue) IsFull() bool {
	next := queue.rear + 1
	if next == queue.maxSize {
		next = 0
	}
	return next == queue.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
