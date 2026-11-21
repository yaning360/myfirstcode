package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //删除元素
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

//最后双赋值查询的解释
//假设Answer不存在则返回0，false;但若存在则返回0，true,go允许同时获得这两个值
//简单的说，只要使用v,ok := m["Answer"]类似的写法查询映射，第一个接收值第二个判断是否存在
