package main

import "fmt"

func main() {
	v := 5 + 12i
	fmt.Printf("v is of type %T\n", v)
}

//声明变量但未指定显式类型时，变量的值都会根据右侧的值推断出来
//当声明右侧的值有明确的类型时，新变量也具有相同的类型
//但当右侧包含无类型数值常量，新变量可能是int,float64或complex128
