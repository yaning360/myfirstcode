package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

//被延迟的函数调用会被压入一个栈中。当函数返回时， 这些延迟调用会按后进先出的顺序执行。
