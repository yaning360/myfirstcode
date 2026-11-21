package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//不过这里不能使用v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

//%g是浮点数的一种简洁格式输出
//声明赋值v的简短语句的作用域在if条件，if代码块和else代码块
//但if-else结束后就不能访问v了
