package main

import "fmt"

func main() {
	// 数组的定义和初始化
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	// 省略的部分为零值

	// 切片的定义和初始化
	slice := []int{1, 2, 3, 4, 5} // 切片字面量

	// 切片操作
	fmt.Println(slice[1:3]) // 输出切片的第2个到第4个元素，不包括索引为3的元素
}
