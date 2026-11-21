package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

//import语句用于导入包。每个包都有一个路径，路径是包的唯一标识符。包的路径通常是包所在的目录的路径。包的路径可以是相对路径，也可以是绝对路径。
//使用括号分组导入是更好的代码风格，也可以逐行导入包
//这里导入math,使用的Sqrt可以用来开平方根

//注意这里使用了Printf而不是Println,前者多用于控制格式的输出如出现%g%f等,且不会自动换行，后者则是直接输出内容
