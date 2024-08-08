package main

import "fmt"

func main() {
	// 映射的定义和初始化
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	// 访问映射
	fmt.Println("map:", m)

	// 删除操作
	delete(m, "k2")
	fmt.Println("map:", m)
}
