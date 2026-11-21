package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil")
	}
}

//切片的零值为nil
//切片的长度和容量均为零并且没有底层数组
