package main

import "fmt"

/**
数组的特点:
数组长度创建完就固定了，不可以再追加元素
数组是值类型的，数组赋值和作为函数参数都是值拷贝
数组长度是数组类型的组成部分，[10]int和[20]int表示不同的类型
可以根据数组创建切片
*/

func main() {
	var a0 [5]int
	// := 需要初始化值
	a1 := [3]int{1, 2, 3}   // 指定长度和初始化字面量
	a2 := [...]int{1, 2, 3} // 不指定长度，但是由后面的初始化列表数量来确定其长度

	fmt.Println(a0)
	for i, v := range a1 {
		fmt.Println(i, v)
	}
	a1[0] = 4
	alength := len(a2)
	for i := 0; i < alength; i++ {

	}
	changeArr(a1)
	changeArrPoint(&a1)
	fmt.Println(a1)
}

func changeArr(arr [3]int) {
	arr[1] = 10
}

func changeArrPoint(arr *[3]int) {
	(*arr)[2] = 20
}
