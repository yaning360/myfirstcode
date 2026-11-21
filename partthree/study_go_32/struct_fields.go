package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	V := Vertex{1, 2}
	V.X = 4
	fmt.Println(V.X)
}

//可以使用点号来访问结构体字段
