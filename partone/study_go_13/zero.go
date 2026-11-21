package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//%q以带引号的形式输出值
//声明变量是如何没有明确指定初始值，它会被赋于零值
