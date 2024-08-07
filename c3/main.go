package main

import "fmt"

type Celsius float64    // 为float64起一个别名Celsius，表示摄氏温度
type Fahrenheit float64 // 为float64起一个别名Fahrenheit，表示华氏温度

// CToF 将摄氏温度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	var c Celsius = 100
	var f Fahrenheit = CToF(c)

	fmt.Println(c, "C is", f, "F") // 输出：100 C is 212 F
}
