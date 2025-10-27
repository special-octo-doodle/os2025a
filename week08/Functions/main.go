package main

import (
	"fmt"
)

func swap(first int, second int) {
	temp := 0
	temp = first
	first = second
	second = temp
	fmt.Println(first, second)
}

func main() {
	var a, b int = 10, 20
	fmt.Print(a, b)
	swap(a, b)
	fmt.Print(a, b)
}
