package main

import "fmt"

func operasiBilangan() {
	var (
		nilaiAkhir = 90
		absensi    = 80

		lulusNilaiAkhir bool = nilaiAkhir > 80
		lulusAbsensi    bool = absensi > 70

		lulus bool = lulusNilaiAkhir && lulusAbsensi
	)

	fmt.Println(lulus)
}
