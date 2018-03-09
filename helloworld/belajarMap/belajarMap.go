package main

import (
	"fmt"
)

func main() {
	/*
		var daymonth = map[string]int{"januari": 31, "february": 28, "maret": 31, "april": 30}
		for key, value := range daymonth {
			fmt.Println(key, " :\t", value)
		}*/
	tesMap()

}
func tesMap() {
	var mapSlice = []map[string]string{
		{"nama": "tuti", "umur": "22", "alamat": "jakarta", "pekerjaan": "A"},
		{"nama": "joni", "umur": "24", "alamat": "bekasi", "pekerjaan": "B"},
	}

	for _, val := range mapSlice {
		fmt.Println("nama ", val["nama"])
		fmt.Println("umur ", val["umur"])
		fmt.Println("alamat : ", val["alamat"])
		fmt.Println("pekerjaan", val["pekerjaan"])

	}
	fmt.Println(len(mapSlice))

}
