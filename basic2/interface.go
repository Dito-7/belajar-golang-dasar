package basic2

import "fmt"

func interfaceExample() {
	var data interface{}
	data = 2
	fmt.Println(data)

	data = "hello"
	fmt.Println(data)
}
