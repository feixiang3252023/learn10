package main

import (
	"errors"
	"fmt"
)

// 错误处理
func doSomething(flag bool) error {
	if !flag {
		// 返回一个错误
		return errors.New("something went wrong")
	}
	return nil
}

func main() {
	err := doSomething(false)
	if err != nil {
		fmt.Println(err)
	}
}
