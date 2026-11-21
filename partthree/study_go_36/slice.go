package main

import "fmt"

func main() {
	primes := [6]int{3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

/*
类型 []T 表示一个元素类型为 T 的切片。
通过指定两个索引来形成切片，即由冒号分隔的下界和上界：
a [LOW:HIGH]包含low但不包含high
*/
