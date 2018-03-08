package main

import "fmt"

func main() {
	cetak(luaslingkaran(5, 10))

}
func luaslingkaran(a float64, b float64) float64 {
	const phi float64 = 3.14
	return phi * a * a * b

}
func cetak(a float64) {
	fmt.Println("luas =", a)
}
