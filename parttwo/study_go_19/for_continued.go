package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
for语句中的初始化语句和后置语句都是可选的
这句话的意思就是初始化语句和后置语句可以不写或直接写进循环体内,并且不会有默认值之类的东西
*/
