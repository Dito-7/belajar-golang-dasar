package main

import (
	"fmt"
)

func great(convert []int) string {
	result := ""

	for _, code := range convert {
		result += string(rune(code))
	}

	return result
}

func convertToUint(s string) ([]uint16, error) {
	var resultU []uint16

	for _, char := range s {
		if char > 65535 {
			continue
		}
		resultU = append(resultU, uint16(char))
	}

	return resultU, nil
}

func konversiData() {
	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)

		name = "eko"
		// strn = string(nilai16)
		str = "Hello World!"

		// convert = []int{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33}
	)
	convertUint, _ := convertToUint(str)
	var convert []int
	for _, val := range convertUint {
		convert = append(convert, int(val))
	}
	fmt.Println(great(convert))

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	fmt.Println(name)

	fmt.Println(great(convert))

	// fmt.Println(convertToUint(str))
}
