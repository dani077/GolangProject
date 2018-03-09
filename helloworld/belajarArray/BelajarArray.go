package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var cari string
	//fmt.Scanln(&cari)
	//arraypertama(cari)
	//var makanan = [2][3]string{{"coklat", "gula", "pecel"}, {"nasi", "mie", "oreo"}}
	/*var klub = []string{"A", "B", "C", "D"}

	var klub2 = klub[0:3]
	var klub3 = klub[1:4]
	var klub4 = klub2[2:3]
	fmt.Println(klub)

	fmt.Println(klub2)
	fmt.Println(klub3)

	fmt.Println(klub4)
	klub4[0] = "Z"
	fmt.Println(klub)

	fmt.Println(klub2)
	fmt.Println(klub3)

	fmt.Println(klub4)

	var klub5 = copy(klub3, klub2)
	fmt.Println(klub5)
	fmt.Println(klub4)
	fmt.Println(klub2)
	fmt.Println(klub3)
	var klub6 = append(klub4, "G")
	fmt.Println(klub6)
	/*
		for baris := 0; baris < len(klub); baris++ {
			fmt.Println("indeks ke ", baris, ":", klub[baris])
		}
		for i, klub := range klub {
			fmt.Println("indeks ke ", i, ":", klub)

		} /*
			for baris := 0; baris < len(makanan); baris++ {
				for kolom := 0; kolom < len(makanan); kolom++ {
					fmt.Println("indeks ke", baris, ",", kolom, ":", makanan[baris][kolom])
				}
			}*/

	var input int
	fmt.Scanln(&input)

	soal1(input)
	soal3(input)
	soal4(input)
	//	segitiga(input)
	//	segitiga2(input)
	//	segitiga3(input)
	//	segitiga4(input)
	//	segitiga5(input)
	//	segitiga6(input)
}
func segitiga(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if j < i {
				vector[j] = "*"
				matrix[i][j] = vector[j] //insert value
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			fmt.Printf(matrix[i][j]) //print

		}
		fmt.Println() // print enter

	}

}

func segitiga2(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if j+i >= input {
				vector[j] = "*"
				matrix[i][j] = vector[j] //insert value
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			fmt.Printf(matrix[i][j]) //print

		}
		fmt.Println() // print enter

	}

}
func segitiga3(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if i >= j && i+j >= input-1 {
				vector[j] = "*"
				matrix[i][j] = vector[j]

			} else if i <= j && i+j <= input-1 {
				vector[j] = "*"
				matrix[i][j] = vector[j]
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			fmt.Printf(matrix[i][j]) //print

		}
		fmt.Println() // print enter

	}

}

func segitiga4(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if i >= j && i+j <= input-1 {
				vector[j] = "*"
				matrix[i][j] = vector[j]

			} else if i <= j && i+j >= input-1 {
				vector[j] = "*"
				matrix[i][j] = vector[j]
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			fmt.Printf(matrix[i][j]) //print

		}
		fmt.Println() // print enter

	}

}
func segitiga5(input int) {
	matrix := make([][]string, input)
	for i := range matrix {
		matrix[i] = make([]string, input)
		for j := range matrix[i] {
			if i == j {
				matrix[i][j] = "*"
			} else if i+j == input-1 {
				matrix[i][j] = "*"
			} else if i == 0 || j == 0 {
				matrix[i][j] = "*"
			} else if i == input-1 || j == input-1 {
				matrix[i][j] = "*"
			} else {
				matrix[i][j] = " "
			}
			fmt.Printf(matrix[i][j])

		}
		fmt.Println()
	}
}

func segitiga6(input int) {
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if i >= j && i+j <= input-1 && i <= (input-1)/2 {
				vector[j] = "*"
				matrix[i][j] = vector[j]

			} else if i <= j && i+j >= input-1 && i >= (input-1)/2 {
				vector[j] = "*"
				matrix[i][j] = vector[j]
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}
			fmt.Printf(matrix[i][j]) //print

		}
		fmt.Println() // print enter

	}

}
func arraypertama(a string) { // seleksi kondisi array
	var club = [4]string{"persija", "PSMS", "Persikabo", "Arema"}

	for _, club := range club {
		if club == a {
			fmt.Println("nama klub sudah ada")
		}

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

func soal1(input int) {
	ganjil := 1
	genap := 16
	matrix := make([][]string, input) // input*input matrix //buat baris
	for i := range matrix {
		matrix[i] = make([]string, input) //buat slice 2 dimensi//buat kolom
		vector := make([]string, input)
		for j := range matrix[i] {
			if i == j {
				vector[j] = strconv.Itoa(ganjil)
				matrix[i][j] = vector[j]

			} else if i+j >= input-1 {
				vector[j] = strconv.Itoa(genap)
				matrix[i][j] = vector[j]
			} else if i+j >= input-1 {
				vector[j] = "A"
				matrix[i][j] = vector[j]

			} else if i+j <= input-1 {
				vector[j] = "B"
				matrix[i][j] = vector[j]
			} else {

				vector[j] = " "
				matrix[i][j] = vector[j] //insert  value
			}

		}

		genap = genap + 2
		ganjil = ganjil - 2

	}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf(matrix[i][j])
		}
		fmt.Println()
	}

}

/*func inputDimensi2Array(a int) {

	c := make([]int, a)

	for kolom := 0; kolom < a; kolom++ { //kolom
		fmt.Scanf(c[kolom])

		fmt.Println()
	}
}*/

func cobaslice(a int) {

	var club = []string{"persija", "PSMS", "Arema", "Semen Padang"}

	for i := 0; i < len(club); i++ {
		fmt.Println(club[i])
	}
	var ligaGojek = append(club[0:1], "bhayangkara FC")

	for i := 0; i < len(ligaGojek); i++ {
		if ligaGojek[i] == club[i] {
			fmt.Println(ligaGojek[i])
		}
	}
	var nomor = []int{2, 3, 4, 5, 7}
	hasil := 1
	for i := 0; i < len(nomor); i++ {

		hasil = hasil * nomor[i]
	}
	fmt.Println(hasil)

}
