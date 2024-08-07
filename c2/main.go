package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u) // 输出：42 42 42

	// 字符串转换为整型
	if intValue, err := strconv.Atoi("42"); err == nil {
		fmt.Println(intValue) // 输出：42
	}
}
