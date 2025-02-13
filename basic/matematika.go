package main

import "fmt"

func matematika() {
	var (
		a            = 10
		b            = 10
		c            = 10
		hasilJumlah  = a + b + c
		hasilKurang  = a - b - c
		hasilKali    = a * b * c
		hasilbagi    = a / b * c
		hasilModulus = a * b % c

		i = 10
	)
	i += 10 // augmented assigments
	i += b + 10

	fmt.Println(hasilJumlah)
	fmt.Println(hasilKurang)
	fmt.Println(hasilKali)
	fmt.Println(hasilbagi)
	fmt.Println(hasilModulus)
	fmt.Println(i)

	// unary operator atau increment dan decrement
	j := 1
	j++
	j++

	fmt.Println(j)
	j--

	fmt.Println(j)

}
