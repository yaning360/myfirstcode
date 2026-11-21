package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
v:=math.Pow(x,n)就是简短语句，if语句可以在条件判断前执行一条简短语句
格式为:if 简短语句; 条件
math.Pow(x,y)就是计算x的y次方



*/
