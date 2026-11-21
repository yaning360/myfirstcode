package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}

//切片可以省略它的上界和下界从而使用它们的默认值
//下界的默认值为零，而上界的默认值为（底层切片或）数组长度
