package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

//在go中，导出名就是可以被其他包访问的名字，如果一个名称以大写字母开头，那么它就是导出名
//Pi就是从包math中导出的导出名，导入一个包后，你只能引用其中的导出名
//任何未导出的名称都无法从包外访问
