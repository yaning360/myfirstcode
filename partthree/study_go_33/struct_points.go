package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

//可以通过结构体指针访问结构体字段
//(*p).X=1e9是另一种比较繁琐的写法
