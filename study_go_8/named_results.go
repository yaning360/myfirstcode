package main

import "fmt"

func spilt(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(spilt(17))
}

/*
函数的返回值可以被命名，如果命名了返回值那它们将被视为定义在函数顶部的变量
比如本程序中的第五行代码，其中(x,y int)就是对返回值的命名
这时候不带参数的返回值将返回被命名的返回值
这被称为“裸返回”
*/
