package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//append可用于nil切片
	s = append(s, 0)
	printSlice(s)

	//切片会按需增长
	s = append(s, 1)
	printSlice(s)

	//我们可以一次添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//append()默认添加元素到末尾
//append 返回的切片包含原切片的所有元素以及所提供的值
//append 的第一个参数是要追加元素的切片，其余参数是要追加到该切片末尾的元素。
