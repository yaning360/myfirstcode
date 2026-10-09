package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
此处的分号可以省略，go语言中的for循环可以直接
不写初始化语句和后置语句及其分号，直接当做while语句使用
*/
