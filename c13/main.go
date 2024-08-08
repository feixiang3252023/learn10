package main

import "fmt"

// 定义一个接口
type Greeter interface {
	greet() string
}

// 实现接口的具体类型
type English struct{}

func (e English) greet() string {
	return "Hello!"
}

// 实现接口的具体类型
type Chinese struct{}

func (c Chinese) greet() string {
	return "你好！"
}

func greetSomeone(g Greeter) {
	fmt.Println(g.greet())
}

func main() {
	var e English
	var c Chinese

	greetSomeone(e) // 输出：Hello!
	greetSomeone(c) // 输出：你好！
}
