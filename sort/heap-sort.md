## 堆排序

堆就是数组实现二叉树，没有父指针或者子指针，堆根据堆属性来排序，堆属性决定了树中节点的位置

### 堆属性

堆分为两种：最大堆和最小堆，差别：父节点的值比每个子节点的值都要大（小）

parent(i) = floor((i-1)/2)
left(i) = 2 * i + 1
right(i) = 2 * i + 2

arr[parent(i)] >= arr[i]

高度：`h = floor(log2(n))` 

节点数：`2 ^ (h+1) - 1`

叶节点位置：`floor(n/2) 和 n-1 `之间

### 堆操作

- 位置上升：shiftUp()  如果节点比父节点大（最大堆），则需要将这个节点上升操作

- 位置下降：shiftDown() 如果节点比子节点小（最大堆），则需要将这个节点向下移动

上面2个操作时间复杂度 O(logn)

- insert(value): 在堆尾部插入元素，然后使用shiftUp来修复

- remove() 移动并返回最大值或者最小值，然后使用shiftDown()来修复堆

- removeAtIndex(index) 和remove()一样，差别在于可以移动堆中任意节点，而不仅仅是根节点。
当它与子节点比较无序则使用shiftDown()修复，如果与父节点比较无序则使用shiftUp()修复

- replace(index, value) 将一个更小的值赋值给一个节点（最小堆），使用shiftUp()修复

上面几个操作都是O(logn)时间复杂度

- search(value) 堆不是为搜索建立的， replace() 和 removeAtIndex() 操作是找到index节点，时间复杂度O(n)

- buildHeap(array)  反复调用  insert() 将一个数组转换为堆。

- 堆排序，由于堆是一个数组，可以使用独特的属性将数组从低到高排序，时间复杂度O(n)

- peak() 不删除返回最大值，最小值，时间复杂度O(1)


### 堆应用场景

- 构建优先队列
- 支持堆排序
- 快速找出集合中最小值或者最大值
- 在朋友面前装逼

### 参考

[数据结构：堆（Heap）](https://www.jianshu.com/p/6b526aa481b1)