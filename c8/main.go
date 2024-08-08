package main

import "fmt"

// 定义一个交换两个数的函数，这里使用指针来直接修改变量的值
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a := 20
	b := 30
	fmt.Println("交换前 a 和 b 的值：", a, b) // 交换前 a 和 b 的值： 20 30
	swap(&a, &b)                       // 这里传递的是变量的地址，也就是指向这些变量的指针
	fmt.Println("交换后 a 和 b 的值：", a, b) // 交换后 a 和 b 的值： 30 20
}
