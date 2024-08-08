package main

import "fmt"

// 定义一个加法函数
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13)) // 调用函数并输出结果
}
