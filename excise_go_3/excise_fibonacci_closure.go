package main

import "fmt"

func fibonacci() func() int {
	numbers := []int{0, 1}
	i := 0
	return func() int {
		if i >= len(numbers) {
			a := numbers[len(numbers)-1]
			b := numbers[len(numbers)-2]
			sum := a + b
			numbers = append(numbers, sum)
		}
		number := numbers[i]
		i++
		return number
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
