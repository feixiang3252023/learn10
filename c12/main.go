package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
}

// 定义结构体的方法
func (p person) sayHello() {
	fmt.Printf("Hi, I'm %s, %d years old.\n", p.name, p.age)
}

func main() {
	// 初始化结构体
	p := person{name: "Jack", age: 23}

	// 调用方法
	p.sayHello() // 输出: Hi, I'm Jack, 23 years old.
}
