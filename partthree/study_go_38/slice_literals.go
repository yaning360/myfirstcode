package main

import "fmt"

func main() {
	q := []int{3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

//字面量就是“直接写成某种类型的一个具体值”
//切片字面量就像省略了长度的数组字面量
