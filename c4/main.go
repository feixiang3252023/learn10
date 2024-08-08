package main

import "fmt"

func main() {
	// 使用if-else进行条件判断
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// 使用switch进行条件判断
	switch num := 7; num {
	case 1:
		fmt.Println("One")
	case 7:
		fmt.Println("Seven")
	default:
		fmt.Println("Unknown Number")
	}
}
