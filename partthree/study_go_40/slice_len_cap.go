package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//切片具有长度(length)和容量(capacity),分别通过len()和cap()
//长度就是切片所包含的长度，容量就是切片的第一个元素开始计算，底层数组的元素个数
