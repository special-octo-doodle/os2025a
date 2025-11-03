package main

import "fmt"

func main() {
	var arrayBool [3]bool = [3]bool{true, false, true}
	var arrayint [3]int
	fmt.Println(arrayBool[1])
	arrayint[1] = 2
	fmt.Println(arrayint[1])
}
