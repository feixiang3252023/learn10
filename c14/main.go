package main

import "fmt"

type Walker interface {
	walk() string
}

type Talker interface {
	talk() string
}

type Human struct{}

func (h Human) walk() string {
	return "I'm walking."
}

func (h Human) talk() string {
	return "I'm talking."
}

func main() {
	var h Human

	// Human实现了Walker接口
	var w Walker = h
	fmt.Println(w.walk())

	// Human实现了Talker接口
	var t Talker = h
	fmt.Println(t.talk())
}
