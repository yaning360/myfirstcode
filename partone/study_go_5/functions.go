package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
}

//这个程序演示了如何定义一个函数
//与c语言不同，类型在变量名之后，即x int
//在这个案例中，函数可以接受零个或多个参数
/*当两个或多个连续的具名函数参数有相同类型时，可省略除最后
一个参数之外其他参数的类型,比如本程序里的add函数可写为func(x , y int)*/
