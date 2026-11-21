package main

import "fmt"

func main() {
	a := make([]int, 5) //len(a)=5
	printSlice("a", a)

	b := make([]int, 0, 5) //len(a)=0,cap(a)=5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//内置函数make可以创建切片，括号内参数格式为
//([]Type,len,cap)分别为数组类型，切片长度，切片容量
//make会分配一个元素均为零值的数组并返回一个引用该数组的切片
//这就是创建动态数组的方式
