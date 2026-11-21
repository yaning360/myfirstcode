package main

import (
	"fmt"
	"strings"
)

//切片可以包括任何类型包括其他切片

func main() {
	//创建一个井字棋棋盘
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//玩家轮流落子
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "	"))
	}
}
