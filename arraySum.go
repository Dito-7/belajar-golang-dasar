package main

import (
	"fmt"
)

func PositiveSum(numbers []int) int {
	sum := 0

	for _, arr := range numbers {
		if arr > 0 {
			sum += arr
		}
	}

	return sum
}

func countPositivesSumNegatives(numbers []int) []int {
	if numbers == nil {
		return []int{}
	}

	countPositive := 0
	sumNegative := 0

	for _, arr := range numbers {
		if arr > 0 {
			countPositive++
		} else if arr < 0 {
			sumNegative += arr
		}
	}

	return []int{countPositive, sumNegative}
}

func main() {
	numbers := []int{1, 2, -5, 10, -10}

	fmt.Println(PositiveSum(numbers))
	fmt.Println(countPositivesSumNegatives(numbers))

}
