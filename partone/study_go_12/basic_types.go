//go语言中基本类型包括bool布尔类型，string字符串类型，
/*
整数类型有int int8 int16 int32 int64
无符号整数类型有uint uint8 uint16 uint32 uint64 uintptr
浮点数类型主要有float32 float64
复数类型有complex64 complex128
*/
/*
byte是 uint8 的别名
rune是 int32 的别名,表示一个 Unicode 码点
*/
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T,Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T,Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T,Value: %v\n", z, z)
}

/*
int,uint,uintptr在32位系统中为32位宽，在64位系统中为64位宽
*/
//%T输出变量的类型，%v按默认格式输出变量的值
//"<<"是左移运算符，二进制往左进n位，1<<64-1就是2的64次方-1
