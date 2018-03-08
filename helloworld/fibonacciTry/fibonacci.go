package main

import "fmt"

func main() {

	var banyak int

	fmt.Scanln(&banyak)
	for i := 1; i <= banyak; i++ {

		fmt.Println(fibonacci(i))

	}
}
func fibonacci(a int) int {

	if a <= 1 {
		return a
	}
	return fibonacci(a-1) + fibonacci(a-2)
}
