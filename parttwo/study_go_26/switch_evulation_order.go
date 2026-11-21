package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's the Saturday?")
	switch today := time.Now().Weekday(); today {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//go语言中的switch语句的各个分支按从上到下求值，一旦有分支匹配成功便停止
