package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

//go的指针保存一个值的地址
/*
*T是指向T类型值的指针，它的零值且nil,nil就是指针的零值
运算符&会生成一个指向其操作数的指针
运算符*表示指针所指向的底层值，这被称为“解引用”或“间接引用”
*/
