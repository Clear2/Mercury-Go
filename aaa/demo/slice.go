package main

import "fmt"

/**
len() 切片长度
cap() 切片数组容量
append() 对切片追加元素
copy() 复制一个切片
*/
func main() {
	initMap()
}

func initSlice() {
	// make创建的切片各元素被默认初始化为切片元素类型的零值，例如
	// len = 10, cap = 10
	a := make([]int, 10)

	// len = 10, cap = 10
	b := make([]int, 10, 15)
	fmt.Println(a, b)

	var c []int
	fmt.Printf("%v\n", c) // 结果为[]， 没有意义的
}

func initMap() {
	// 作为无序兼职对集合，字典要求key必须是支持相等(==, !=)运算符的数据类型
	m := make(map[string]int)
	m["a"] = 1

	m2 := map[int]struct { // 值是匿名结构体
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2)
}
