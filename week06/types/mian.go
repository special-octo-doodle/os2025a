package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var text string = "mao"
	// fmt.Println(math.Floor(2.75))
	// fmt.Println(strings.Title("hee lo"))
	// fmt.Println(strings.Title(text))

	// name := "mao mao mao"

	var a float64
	var b string
	var c int32
	var d bool

	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))
}
