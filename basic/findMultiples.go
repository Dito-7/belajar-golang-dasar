package main

import "fmt"

func FindMultiples(numbers []int) []int {
	multiple := 0

	for _, integer := range numbers {
		if len(numbers) == len(numbers) {
			multiple = numbers[0] + integer
			numbers = append(numbers, multiple)
			return numbers
		}
	}
	return numbers
}

func findMult() {
	numbers := []int{2, 6}
	// slice := numbers[1:2]
	fmt.Println(FindMultiples(numbers))
}
