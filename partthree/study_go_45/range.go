package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

//range是for循环的一种形式可用于遍历切片或映射
//遍历切片时每次迭代都会返回两个值，第一个是元素的索引，第二个是该索引处元素的副本
//副本就是把原来的值复制一份，得到一个独立的值。修改这个副本，不会直接修改原来的元素。
