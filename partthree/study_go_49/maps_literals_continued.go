package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.042202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

//如果顶层类型只是一个类型名，那么在字面量的各个元素中可以省略该类型
