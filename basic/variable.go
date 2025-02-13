package main

import (
	"fmt"
)

func basic() {
	var name string

	name = "ditoaditya"
	fmt.Println(name)

	name = "nugroho"
	fmt.Println(name)

	var namee = "nasya"
	fmt.Println(namee)

	//:= itu seperti var yang mendefinisikan tipe data
	nameee := "tidak perlu menggunakan var"
	fmt.Println(nameee)

	var (
		firstname = "dito"
		lastname  = "aditya"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	//constant adalah data yang tidak bisa diubah
	const kelahrian string = "sidoarjo"
	fmt.Println(kelahrian)
}
