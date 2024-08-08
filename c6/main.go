package main

import "fmt"

func main() {
	// 使用break跳出循环
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	// 使用continue跳过某次循环迭代
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// 使用goto跳转
	i := 0
loop: // 标签
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop // 跳转到标签位置
	}
}
