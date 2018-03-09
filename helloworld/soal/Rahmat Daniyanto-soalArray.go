package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input int
	fmt.Println("masukkan input")
	fmt.Scanln(&input)
	fmt.Println("soal 1")
	soal1(input)
	fmt.Println("=========================================")

	fmt.Println("soal 2")
	soal2(input)
	fmt.Println("=========================================")

	fmt.Println("soal 3")
	soal3(input)
	fmt.Println("=========================================")

	fmt.Println("soal 4")
	soal4(input)
	fmt.Println("=========================================")
}
func soal1(input int) {
	matrix := make([][]string, input) // dim*dim matrix

	//genap:=16
	for i := 0; i < input; i++ {
		ganjil := 1
		genap := 0
		matrix[i] = make([]string, input)
		vector := make([]string, input)
		for j := 0; j < input; j++ {
			//	if i <= j && i+j <= input-1 { //atas
			if i == j { //diagonal kiri
				vector[j] = strconv.Itoa(ganjil)
				matrix[i][j] = vector[j]

			} else if i+j == input-1 { //diagonal kanan
				vector[j] = strconv.Itoa(genap)
				matrix[i][j] = vector[j]

			} else if i <= j && i+j <= input-1 { //atas
				vector[j] = "A"
				matrix[i][j] = vector[j]
			} else if i >= j && i+j >= input-1 { //bawah
				vector[j] = "B"
				matrix[i][j] = vector[j]

			} else {
				vector[j] = " "
				matrix[i][j] = vector[j]

			}
			ganjil = ganjil + 2
			genap = genap + 2
		}
	}
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print(matrix[i][j], "\t")
		}
		fmt.Println()
	}

}

func soal2(input int) {
	matrix := make([][]string, input) // dim*dim matrix

	ganjil2 := 1
	//genap:=16
	for i := 0; i < input; i++ {
		ganjil := 1
		genap := 0

		if i != 0 {
			ganjil2 = ganjil2 + 2
		}
		matrix[i] = make([]string, input)
		vector := make([]string, input)
		for j := 0; j < input; j++ {

			if i == j { // diagonal kiri
				vector[j] = strconv.Itoa(ganjil)
				matrix[i][j] = vector[j]

			} else if i+j == input-1 { // diagonal ganal
				vector[j] = strconv.Itoa(genap)
				matrix[i][j] = vector[j]

			} else if i == (input-1)/2 { //horizontal
				vector[j] = strconv.Itoa(genap)
				matrix[i][j] = vector[j]
			} else if j == (input-1)/2 { //vertikal
				vector[j] = strconv.Itoa(ganjil2)
				matrix[i][j] = vector[j]

			} else {
				vector[j] = " "
				matrix[i][j] = vector[j]

			}
			ganjil = ganjil + 2
			genap = genap + 2
		}

	}
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print(matrix[i][j], "\t")
		}
		fmt.Println()
	}

}
func soal3(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		ganjil := 1

		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if j < i {
				vector[j] = strconv.Itoa(ganjil)

				matrix[i][j] = vector[j] //insert value
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			ganjil = ganjil + 2 //ganti kolom
		}

	}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf(matrix[i][j])
		}
		fmt.Println()
	}

}
func soal4(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris

	for i := range matrix {
		ganjil := 1

		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {

			vector[j] = strconv.Itoa(ganjil)

			matrix[i][j] = vector[j] //insert value

			if j < (input-1)/2 {
				ganjil = ganjil + 2
			} else {
				ganjil = ganjil - 2
			}

		}

	}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf(matrix[i][j])
		}
		fmt.Println()
	}

}
