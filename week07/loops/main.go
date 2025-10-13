package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("int", int(length)*width)
	fmt.Println("float64", length*float64(width))

	fmt.Println("int value:", reflect.TypeOf(width), width)
	fmt.Println("float64 value:", reflect.TypeOf(length), length)

	fmt.Println("int to float64 value:", reflect.TypeOf(float64(length)), float64(length))
	fmt.Println("float64 to int value:", reflect.TypeOf(int(width)), int(width))

	var now time.Time = time.Now()
	var day int = now.Day()
	fmt.Println(day)

	univ := "Go$ inha$"
	changer := strings.NewReplacer("$", "!")
	replaced := changer.Replace(univ)
	fmt.Println(replaced)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)
}
