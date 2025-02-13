package main

import (
	"fmt"
)

func sliceGo() {
	names := [...]string{"halo", "nama", "aku", "dito", "aditya", "nugroho"}
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	newSlice := make([]string, 2, 5)

	slice1 := names[3:5]
	slice2 := names[:3]
	slice3 := names[3:]
	slice4 := names[:]

	daySlice1 := days[5:]

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"

	daySlice2 := append(daySlice1, "libur Baru")

	newSlice[0] = "dito"
	newSlice[1] = "aditya"
	newSlice1 := append(newSlice, "nugroho", "keren")

	fmt.Println(slice1)
	fmt.Println("len: ", len(slice1))
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println(daySlice1)
	fmt.Println(days)
	fmt.Println(daySlice2)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice1)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5, 6}
	iniSlice := []int{1, 2, 3, 4, 5, 6} //lebih baik gunakan slice saja

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
