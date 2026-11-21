package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界！"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//常量的声明方式与变量类似，但使用const关键字
//常量可以是字符串，字符，布尔值或数值，但不能使用：=声明
