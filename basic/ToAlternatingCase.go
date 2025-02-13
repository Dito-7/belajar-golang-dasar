package main

import (
	"fmt"
	"unicode"
)

func ToAlternatingCase(s string) string {
	result := ""
	for _, char := range s {
		if unicode.IsUpper(char) {
			result += string(unicode.ToLower(char))
		} else if unicode.IsLower(char) {
			result += string(unicode.ToUpper(char))
		} else {
			result += string(char)
		}
	}

	return result
}

func alternateCase() {
	fmt.Println(ToAlternatingCase("hello world"))
	fmt.Println(ToAlternatingCase("HELLO WORLD"))
	fmt.Println(ToAlternatingCase("hello WORLD"))
	fmt.Println(ToAlternatingCase("HeLLo WoRLD"))
	fmt.Println(ToAlternatingCase("12345"))
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))
	fmt.Println(ToAlternatingCase("String.prototype.toAlternatingCase"))
}
