package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

//T(v)可以将值v的类型转换为类型T
//或可以更简洁
/*
x,y:=3,4
f:=math.Sqrt(float(x*x+y*y))
z:=unit(f)
*/
