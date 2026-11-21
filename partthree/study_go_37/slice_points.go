package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

//go语言中的切片就类似于数组的引用，并不储存数据只是描述了底层数组的一段
//修改切片中的元素，会同时修改其底层数组中对应的元素
//共享同一个底层数组的其他切片也会看到这些变化
