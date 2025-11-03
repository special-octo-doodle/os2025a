package main

import (
	"fmt"
)

func main() {
	arrayBool := [3]bool{true, false, true}
	arrayint := [3]int{-9, 10, 4}
	for i := 0; i < 3; i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayint[i])
	}

	// fmt.Println(reflect.TypeOf(arrayBool))
	// fmt.Printf("%#v\n", arrayBool)
	// fmt.Printf("%#v\n", arrayint)
}
