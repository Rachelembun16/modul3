package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("===HITUNG KARAKTER===")
	fmt.Println("")
	fmt.Print("Masukkan suatu kata= ")
	fmt.Scanln(&input)
	hasil := len(input)
	fmt.Printf("jumlah karakter = %d", hasil)
	fmt.Println("")
	fmt.Printf("Karakter yang diinputkan = %s", input)
	fmt.Println("")

	for i := 0; i < hasil-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if string(input[i]) == string(input[j]) {
				var jumlah = strings.Count(input, string(input[i]))
				fmt.Printf("Jumlah %s = %d", string(input[i]), jumlah)
				fmt.Println("")
			}
		}
	}

}
