package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

//映射就是将键映射到值，就如同Python中的字典的键值对
//映射的零值为nil,nil映射没有任何，键也不能想其中添加值
//make()的作用就是创建并初始化映射使它能够正常储值键值对
