package main

import "fmt"

func main() {
	var arrayBool [3]bool
	var arrayint [3]int
	fmt.Println(arrayBool[1])
	arrayint[1]++
	arrayint[1]++
	fmt.Println(arrayint[1])
}
