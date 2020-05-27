/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/11 下午7:34
***********************************************/

package main

import "fmt"

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println("1===", cache.Get(1)) // 返回  1
	cache.Put(3, 3)                   // 该操作会使得密钥 2 作废
	fmt.Println("2===", cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)                   // 该操作会使得密钥 1 作废
	fmt.Println("1===", cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println("3===", cache.Get(3)) // 返回  3
	fmt.Println("4===", cache.Get(4)) // 返回  4
}

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		this.remove(node)
		this.putHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		if len(this.m) >= this.cap {
			deleteKey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, deleteKey)
		}
		newNode := Node{
			key:   key,
			value: value,
		}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

// 删除node节点
func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// 将node放到头节点
func (this *LRUCache) putHead(node *Node) {
	originNext := this.header.next
	this.header.next = node
	node.next = originNext

	originNext.pre = node
	node.pre = this.header

}
