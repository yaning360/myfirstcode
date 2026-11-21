package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//func adder()后的func(int) int是一整个函数类型，表示接受一个int类型返回一个int类型
//adder函数return部分也是一个整体，表示返回一个函数
