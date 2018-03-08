package main

import (
	"fmt"
	"math"
)

func main() {

	var akhir float64 = 100

	fmt.Println("masukkan batas akhir")
	fmt.Scanln(&akhir)
	var i float64
	for i = 2; i < akhir; i++ {

		cekPrima(i)

	}

}
func cekPrima(a float64) {

	var flag = 1
	var batas = a

	var primehelper = math.Sqrt(float64(batas))
	var i float64
	for i = 2; i < primehelper; i++ {
		if math.Mod(batas, i) == 0 {
			flag = 0

		}
	}
	if flag == 1 {
		fmt.Println(batas, " is Prime ")

	} else if flag == 0 {
	}

}
