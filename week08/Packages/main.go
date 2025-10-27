package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("점수: ")
	number, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	if number >= 90 {
		fmt.Println("A")
	} else if number >= 80 {
		fmt.Println("B")
	}
}
