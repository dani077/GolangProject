package main

import (
	"fmt"
)

func main() {

	var panjang int
	var lebar int
	var tinggi int
	var jari2 int
	var bangun int

	fmt.Println("masukkan panjang  :")
	fmt.Scanln(&panjang)
	fmt.Println("masukkan lebar   :")
	fmt.Scanln(&lebar)
	fmt.Println("masukkan tinggi  :")
	fmt.Scanln(&tinggi)
	fmt.Println("masukkan jari-jari: ")
	fmt.Scanln(&jari2)

	fmt.Println("pilih bangun")
	fmt.Println("1. kerucut\n 2.tabung\n 3.balok")
	fmt.Println("masukkan jenis bangun: ")
	fmt.Scanln(&bangun)

	if bangun == 1 {
		fmt.Println("volume kerucut: ", volumekerucut(jari2, tinggi))
	} else if bangun == 2 {
		fmt.Println("volume tabung: ", volumetabung(jari2, tinggi))
	} else if bangun == 3 {

		fmt.Println("volume balok ", volumebalok(panjang, lebar, tinggi))
	} else {
		fmt.Println("Ngopi ngapa Ngopi ,kaga ada di pilihannya tong")
	}

}
func volumekerucut(a int, b int) float64 {
	const phi float64 = 3.14
	return float64(a*a*b*1/3) * phi
}
func volumetabung(a int, b int) float64 {
	const phi float64 = 3.14
	return float64(a*a*b) * phi
}
func volumebalok(a int, b int, c int) int {
	return a * b * c
}
