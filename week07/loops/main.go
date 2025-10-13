package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("int", int(length)*width)
	fmt.Println("float64", length*float64(width))

	fmt.Println("int value:", reflect.TypeOf(width))
	fmt.Println("float64 value:", reflect.TypeOf(length))

	fmt.Println("int to float64 value:", reflect.TypeOf(float64(length)))
	fmt.Println("float64 to int value:", reflect.TypeOf(int(width)))
}
