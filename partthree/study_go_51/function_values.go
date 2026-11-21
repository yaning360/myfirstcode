package main

import (
	"fmt"
	"math"
)

// 函数也是值，也可以被传递
// 函数值可以用作函数参数和返回值
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		//此处为匿名函数的用法，即不写函数名直接将函数值赋值给hypot
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
