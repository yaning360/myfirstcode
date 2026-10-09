package main

import "fmt"

func main() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

/*
在函数内部，可以使用 := 短赋值语句，替代带隐式类型的 var 声明。

在函数外部，每条语句都以关键字开头（ var、 func 等），因此不能使用 := 这种写法。
*/
