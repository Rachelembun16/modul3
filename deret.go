package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Println("===DERET GANJIL GENAP===")
	fmt.Print("Masukkan suatu angka= ")
	fmt.Scan(&input)

	fmt.Print("Deret Ganjil= ")
	for i := 1; i <= input*2; i++ {
		if i%2 != 0 {
			ganjil := i
			fmt.Printf("%d ", ganjil)
		}
	}

	fmt.Println(" ")
	fmt.Print("Deret Genap= ")
	for i := 1; i <= input*2; i++ {
		if i%2 == 0 {
			genap := i
			fmt.Printf("%d ", genap)
		}
	}
}
