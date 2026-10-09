package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
fmt.Sprint的作用是将内容格式化成字符串但不会直接输出到屏幕上
go语言中的if语句与for语句类似，不需要（），但必须{}
*/
