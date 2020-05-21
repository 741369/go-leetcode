/**********************************************
** @Des: lru 删除最近未使用缓存算法
** @Author: lzw
** @Last Modified time: 2020/5/21 上午9:37
***********************************************/

package main

import "fmt"

func main() {
	lrcCache := Constructor(2)
	lrcCache.Put(1, 1)
	lrcCache.Put(2, 2)
	fmt.Println("1==", lrcCache.Get(1)) // 返回1
	lrcCache.Put(3, 3)                  // 该操作使得密钥2作废
	fmt.Println("2==", lrcCache.Get(2)) // 返回-1
	lrcCache.Put(4, 4)                  // 该操作使得密钥1作废

	fmt.Println("3==", lrcCache.Get(1)) // 返回-1
	fmt.Println("4==", lrcCache.Get(3)) // 返回3
	fmt.Println("5==", lrcCache.Get(4)) // 返回4
}

type LRUCacheNode struct {
	key   int
	value int
	prev  *LRUCacheNode
	next  *LRUCacheNode
}

// map + 双链表
type LRUCache struct {
	md   map[int]*LRUCacheNode
	head *LRUCacheNode
	tail *LRUCacheNode
	len  int
	cap  int
}

// 构建LRUCache
func Constructor(cap int) LRUCache {
	lc := LRUCache{
		md:  make(map[int]*LRUCacheNode),
		len: 0,
		cap: cap,
	}
	lc.head = &LRUCacheNode{}
	lc.tail = &LRUCacheNode{}
	lc.head.prev = nil
	lc.head.next = lc.tail
	lc.tail.prev = nil
	lc.tail.next = lc.head

	return lc
}

// 获取
func (lc *LRUCache) Get(key int) int {
	// 判断key对应的键值是否存在
	node, exist := lc.md[key]
	if !exist {
		return -1
	}
	// 存在，则将节点移动到头部
	lc.MoveToHead(node)
	return node.value
}

// 插入
func (lc *LRUCache) Put(key int, value int) {
	// 判断key是否存在
	if _, exist := lc.md[key]; exist {
		// 存在，移动到头部
		lc.md[key].value = value
		lc.MoveToHead(lc.md[key])
	} else { // 不存在，插入到节点头部
		if lc.len > lc.cap {
			panic("impossible error")
		}

		// 1. 检查是否已达到最大容量，达到则淘汰末尾节点
		if lc.len == lc.cap {
			lc.AbandonTail()
			lc.len--
		}

		// 2. 添加新节点到头部
		node := &LRUCacheNode{
			key:   key,
			value: value,
		}
		lc.AddToHead(node)
		lc.md[key] = node
		lc.len++
	}
}

// 插入到表头
func (lc *LRUCache) AddToHead(node *LRUCacheNode) {
	node.prev = lc.head
	node.next = lc.head.next

	lc.head.next.prev = node
	lc.head.next = node
}

// 移动元素到表头
func (lc *LRUCache) MoveToHead(node *LRUCacheNode) {
	lc.RemoveNode(node)
	lc.AddToHead(node)
}

// 删除表尾元素
func (lc *LRUCache) AbandonTail() {
	if lc.len > 0 {
		delete(lc.md, lc.tail.prev.key)
		lc.RemoveNode(lc.tail.prev)
	}
}

// 删除元素
func (lc *LRUCache) RemoveNode(node *LRUCacheNode) {
	node.next.prev = node.prev
	node.prev.next = node.next

	node.next = nil
	node.prev = nil
}
