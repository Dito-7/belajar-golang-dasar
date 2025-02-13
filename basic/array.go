package main

import "fmt"

func basicArray() {
	names := [3]string{"dito", "aditya", "nugoroh"}
	age := [3]int{12, 13, 14}
	gender := [...]string{"laki-laki", "perempuan"}

	var house [2]string
	house[0] = "sidoarjo"
	house[1] = "surabaya"

	fmt.Println(names[1])
	fmt.Println(house[1])
	fmt.Println(age)
	fmt.Println(gender)
	fmt.Println(len(names))
	fmt.Println(len(gender))
}
