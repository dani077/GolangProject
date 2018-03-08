package main

import "fmt"

func main() {

	var input int
	fmt.Scanln(&input)
	//gambar(input)

	//muter2(input)
	//segitiga(input)
	//segitiga2(input)
	segitiga3(input)
	//segitiga4(input)
	//segitiga5(input)
	//segitiga6(input)
}
func gambar(input int) {
	ganjil := 1
	genap := 0
	for baris := 1; baris <= input; baris++ { //baris

		for kolom := 1; kolom <= input; kolom++ { //kolom

			//fmt.Print(kolom, ",", baris)
			if baris == kolom {

				fmt.Print(ganjil, " \t")
				if baris == (input+1)/2 {
					genap = genap + 2
				}
				ganjil = ganjil + 2
			} else if baris+kolom == input+1 {
				fmt.Print(genap, " \t")
				genap = genap + 2
			} else if baris == (input+1)/2 {
				fmt.Print(kolom, " \t")

			} else if kolom == (input+1)/2 {
				fmt.Print(baris, " \t")
			} else {
				fmt.Printf("\t")
			}

		}

		fmt.Println()
	}

}
func segitiga(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if kolom <= baris {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")

			}
		}

		/*for bintang := 1; bintang <= baris; bintang++ {
			fmt.Print("*")
			fmt.Print(" ")

		}*/

		fmt.Println()
	}
}
func segitiga2(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if baris+kolom < input+1 {
				fmt.Printf(" ")
			} else {
				fmt.Print("*")

			}
		}

		/*for bintang := 1; bintang <= baris; bintang++ {
			fmt.Print("*")
			fmt.Print(" ")

		}*/

		fmt.Println()
	}
}
func segitiga3(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if baris <= kolom && baris+kolom <= input+1 { //atas
				fmt.Printf("*")

			} else if baris >= kolom && kolom+baris >= input+1 { //bawah
				fmt.Printf("*")

			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	/*for bintang := 1; bintang <= baris; bintang++ {
		fmt.Print("*")
		fmt.Print(" ")

	}*/

	fmt.Println()
}

func segitiga4(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if baris <= kolom && baris+kolom >= input+1 { //kanan
				fmt.Printf("*")

			} else if baris >= kolom && kolom+baris <= input+1 { //kiri
				fmt.Printf("*")

			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	/*for bintang := 1; bintang <= baris; bintang++ {
		fmt.Print("*")
		fmt.Print(" ")

	}*/

	fmt.Println()
}
func segitiga5(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if baris == kolom {

				fmt.Print("* \t")

			} else if baris+kolom == input+1 {
				fmt.Print("* \t")

			} else if baris == 1 || kolom == 1 {
				fmt.Print("* \t")

			} else if kolom == input || baris == input {
				fmt.Print("* \t")
			} else {
				fmt.Printf("\t")
			}

			/*for bintang := 1; bintang <= baris; bintang++ {
				fmt.Print("*")
				fmt.Print(" ")

			}*/

		}
		fmt.Println()
	}

}

func segitiga6(input int) {

	for baris := 1; baris <= input; baris++ {
		for kolom := 1; kolom <= input; kolom++ {
			if baris <= kolom && baris+kolom >= input+1 && baris >= (input+1)/2 {
				fmt.Printf("*")

			} else if baris >= kolom && kolom+baris <= input+1 && baris <= (input+1)/2 {
				fmt.Printf("*")

			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	/*for bintang := 1; bintang <= baris; bintang++ {
		fmt.Print("*")
		fmt.Print(" ")

	}*/

	fmt.Println()
}
