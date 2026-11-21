package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

//没有条件的Switch语句就相当于switch true
//也就是每个case分支自己写一个布尔条件，自己判断是否成立
//这种写法可以代替写非常长的if-else代码
