package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("My favorite number is", rand.Intn(10))
}

//每个go程序都是由包组成的，包是Go程序的基本组成单元。每个Go程序都是由包组成的，包是Go程序的基本组成单元。
//程序从包main开始执行。
//该程序导入了两个包：fmt和math/rand。fmt包实现了格式化的I/O函数，rand包实现了伪随机数生成器。
