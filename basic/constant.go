package main

import (
	"fmt"
)

func constant() {
	//constant adalah data yang tidak bisa diubah
	const (
		kelahrian string = "sidoarjo"
		lastname         = "nugroho"
	)
	fmt.Println(kelahrian, lastname)

	// error
	// kelahrian = "asdasd"

}
