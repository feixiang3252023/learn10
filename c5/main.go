package main

import "fmt"

func main() {
	// 像while一样使用for
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) // 输出: 1024

	// 传统的for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
