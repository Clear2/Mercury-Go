package main

import "fmt"

/**
 内置函数new 按指定类型长度分配零值内存, 返回指针， 并不关心内部构造和初始化方式

 引用类型则必须使用make函数创建，编译器会将make转换为目标类型专用的创建函数
	以确保完成全部内存分配和相关属性初始化

*/

func main() {
	data := [3]string{"a", "b", "c"}
	for i, s := range data {
		fmt.Println(i, s)
	}
}

func makeSlice() []int {
	s := make([]int, 10, 10)
	s = append(s, 100)
	return s
}

func makeMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func testNew() {
	p := new(map[string]int)
	// p["a"] = 1 Invalid operation: 'p["a"]' (type '*map[string]int' does not support indexing)
	m := *p
	m["a"] = 1 //  assignment to entry in nil map
	// new没有分配键值存储内存，因此无法正常工作
}
