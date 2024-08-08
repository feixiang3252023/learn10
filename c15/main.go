package main

import "fmt"

type Mover interface {
	move() string
}

type Shaker interface {
	shake() string
}

// Combiner接口组合了Mover和Shaker接口
type Combiner interface {
	Mover
	Shaker
}

type Animal struct{}

func (a Animal) move() string {
	return "Animal is moving."
}

func (a Animal) shake() string {
	return "Animal is shaking."
}

func main() {
	var a Animal

	// Animal实现了Combiner接口
	var c Combiner = a
	fmt.Println(c.move())
	fmt.Println(c.shake())
}
