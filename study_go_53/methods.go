package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

//方法就是一个带有接收者参数的函数。接收者出现在 func 关键字和方法名之间。
// 上面例子中，Abs 方法拥有一个名为 v 的 Vertex 类型的接收者。
// 通过这种方式，Abs 方法可以访问调用它的 Vertex 的字段。
