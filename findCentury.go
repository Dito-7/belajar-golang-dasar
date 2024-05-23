package main

import "fmt"

func century(year int) int {
	if year <= 0 {
		fmt.Println("not in the century")
	} else if year%100 == 0 {
		return year / 100
	} else {
		return (year / 100) + 1
	}

	return year
}

func main() {
	cen := century(1999)
	fmt.Println(cen)

	cen = century(2000)
	fmt.Println(cen)

	cen = century(2001)
	fmt.Println(cen)
}
