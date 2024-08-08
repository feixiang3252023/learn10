package main

import "fmt"

// 定义一个分割整数的函数，返回商和余数
func divmod(a, b int) (int, int) {
	quo := a / b // 商
	rem := a % b // 余数
	return quo, rem
}

// 使用命名返回值的方式来返回多个值
func divmodNamed(a, b int) (quo, rem int) {
	quo = a / b // 商
	rem = a % b // 余数
	return      // 不需要明确指定返回值，因为已经在函数签名中命名
}

func main() {
	quo, rem := divmod(7, 3)
	fmt.Println("分割7和3得到：商 =", quo, "余数 =", rem) // 分割7和3得到：商 = 2 余数 = 1

	quoNamed, remNamed := divmodNamed(7, 3)
	fmt.Println("命名返回值分割7和3得到：商 =", quoNamed, "余数 =", remNamed) // 命名返回值分割7和3得到：商 = 2 余数 = 1
}
