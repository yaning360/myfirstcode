package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

//type是用来定义新的类型，而var则是用来创建某个类型的变量
//
