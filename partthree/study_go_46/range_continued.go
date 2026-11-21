package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //==2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

//若想跳过索引或值可以通过赋值 _
//如果只需要索引则省略第二个变量即可
