package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var JumlahProduk int
	var JumlahUang int
	var KodeNamaBarang string
	var kembalian int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Selamat Datang di Deltamart")
	fmt.Println("1.Indomie Rp 3500,-")
	fmt.Println("2.swallow Rp 100000,-")
	fmt.Println("3.Green Tea Rp 500")

	fmt.Println()
	fmt.Println("Masukkan Kode Barang atau Nama Barang")
	scanner.Scan()
	KodeNamaBarang = scanner.Text()

	fmt.Println("Masukkan Jumlah Uang")
	fmt.Scanln(&JumlahUang)
	fmt.Println("Masukkan Jumlah Produk")
	fmt.Scanln(&JumlahProduk)

	fmt.Println("============================================")
	if KodeNamaBarang == "1" || KodeNamaBarang == "Indomie" {
		fmt.Println("Anda  membeli Indomie dengan Harga Rp.3500,-")
	} else if KodeNamaBarang == "2" || KodeNamaBarang == "swallow" {
		fmt.Println("Anda  membeli Swallow dengan Harga Rp.100000,-")
	} else if KodeNamaBarang == "3" || KodeNamaBarang == "Green Tea" {
		fmt.Println("Anda  membeli Green Tea dengan Harga Rp 500,-")
	} else {
		fmt.Println("Maaf barang tidak tersedia")
	}

	fmt.Println("============================================")

	switch KodeNamaBarang {
	case "1", "Indomie":
		fmt.Println("Anda  membeli Indomie dengan Harga Rp.3500,-")
		kembalian = JumlahUang - JumlahProduk*3500
		if kembalian < 0 {
			fmt.Println("Uang kurang", kembalian)
		} else {
			fmt.Println("Jumlah Harga", JumlahProduk*3500)
			fmt.Println("Kembalian", kembalian)
			fmt.Println("Terima Kasih telah berbelanja di Deltamart")

		}
	case "2", "Swallow":
		fmt.Println("Anda  membeli Swallow dengan Harga Rp.100000,-")
		kembalian = JumlahUang - JumlahProduk*100000
		if kembalian < 0 {
			fmt.Println("Uang kurang", kembalian)
		} else {
			fmt.Println("Jumlah Harga", JumlahProduk*100000)
			fmt.Println("Kembalian", kembalian)
			fmt.Println("Terima Kasih telah berbelanja di Deltamart")

		}
	case "3", "Green Tea":
		fmt.Println("Anda  membeli Green Tea dengan Harga Rp.500,-")
		kembalian = JumlahUang - JumlahProduk*500
		if kembalian < 0 {
			fmt.Println("Uang kurang", kembalian)
		} else {
			fmt.Println("Jumlah Harga", JumlahProduk*500)
			fmt.Println("Kembalian", kembalian)
			fmt.Println("Terima Kasih telah berbelanja di Deltamart")
		}
	}
}
