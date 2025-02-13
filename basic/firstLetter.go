package main

import (
	"fmt"
	"unicode"
)

func firstLetter() {
	upper := "How can mirrors be real if our eyes aren't real"
	fmt.Println(upper)

	var output []rune
	isWord := true

	for _, code := range upper {
		if isWord && unicode.IsLetter(code) {
			output = append(output, unicode.ToUpper(code))
			isWord = false
		} else if unicode.IsSpace(code) {
			isWord = true
			output = append(output, code)
		} else {
			output = append(output, code)
		}
	}

	fmt.Println(string(output))

	return
}
