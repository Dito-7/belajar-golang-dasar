package main

import (
	"fmt"
	"strconv"
)

func intToString(num int) string {
	return strconv.Itoa(num)
}

func main() {
	num1 := 123
	num2 := 999
	num3 := -100

	fmt.Println(intToString(num1))
	fmt.Println(intToString(num2))
	fmt.Println(intToString(num3))
}
