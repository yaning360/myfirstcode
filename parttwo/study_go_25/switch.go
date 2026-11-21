package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		//freeBSD,openBSD
		//plan9,windows...
		fmt.Printf("%s.\n", os)
	}
}

/*
switch语句就是编写一系列if-else语句的更简洁方式
与if语句一样可以先用简短语句声明变量再判断
与c语言中的Switch类似，但不同的是
Switch智慧之星选中的分支，而不会继续执行后续的分支
实际上在c语言中每个分支末尾的break语句在go语言中会自动完成
另一个重要的区别是Go中的switch分支不必是常量也不必是整数
*/
