package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*函数标准写法
func 函数名(参数名 参数类型) 返回值类型 {
    // 具体内容
    return 返回值
}
*/

/*该程序中的:=表示短变量声明，等价于先声明变量再赋值，即
var a,b string
a,b = swap("hello","world")
需要注意的是这种写法只能在函数内部写
函数可以返回任意数量的返回值
*/
