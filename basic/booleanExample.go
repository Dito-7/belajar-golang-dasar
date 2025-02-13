package main

import "fmt"

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}

func booleanExample() {
	var employed bool = false
	var vacation bool = false

	result := SetAlarm(employed, vacation)

	fmt.Println(result)
}
