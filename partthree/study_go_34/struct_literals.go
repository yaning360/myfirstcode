package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} //Y:0是隐式的
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, p)
}

/*
Vertex{1,2}按字段顺序初始化
Vertex{X:1}指定字段初始化
Vertex{}全部使用零值
&Vertex{1,2}创建 Vertex 并取地址
*/
