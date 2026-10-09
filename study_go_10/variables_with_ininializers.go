package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

/*
var声明可以包含初始值，每个变量对应着一个初始值
如果指定了初始值，可以省略类型；变量会采用初始值的类型。
比如var i = 12,12很显然是个int类型，所以go会自动将i识别为int类型
*/
