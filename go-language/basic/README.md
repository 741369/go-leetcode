## Go学习总结

多看示例，多谢代码，多总结

学习的过程中对知识点汇总，试着自己去整体梳理，掌握基本知识大纲，最后转换为自己的知识体系。

### 各大语言对比

Golang PHP Python JavaScript Java C++

多核计算 编译速度 学习难度 执行效率 桌面支持 框架&生态 社区活跃

### Go语言的关键字

golang一共25个关键字，体现了设计者们大道至简的思想

package import const type struct interface func go return defer
if else for range switch case default break continue select
var map chan goto fallthrough

### Go语言中的主要概念

基本类型 数组 切片 结构体 指针 字典 鸭子类型 并发协程
函数 接口 类型转换 异常处理 垃圾回收 反射  
包 文件 I/O 安全 网络 内存 UTF-8

### 值类型与引用类型

| 区别\类型 |  值 | 引用 |
|---|---|---|
| 类型成员 | int,float,bool,string,array,struct | slice,map,chan,interface |
| 内存分配 | 栈 | 堆 |
| 编译器 | 声明后默认分配 | 申请后才可使用 |
| 初始化方法 | new() | make() |
| 赋值后改值 | 不变 | 变 |

### 数据类型

基本数据类型

- 布尔型 只可以是常量true或者false
- 数值型 整型 int,int8,int16,int32,int64,uint,uintptr,uint8,byte,uint16,uint32,uint64 浮点型 float32,float64,complex64,complex128
- 字符串 string

派生数据类型

指针 数组 结构体 通道(channel) 切片(slice) 函数 接口(interface) 字典(map) 

