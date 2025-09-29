package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var text string = "mao"
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("hee lo"))
	fmt.Println(strings.Title(text))

	name := "mao mao mao"

	fmt.Println(name, reflect.TypeOf(name))
}
