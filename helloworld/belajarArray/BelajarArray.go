package main

import (
	"fmt"
)

func main() {

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
