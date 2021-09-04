package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Println("===HITUNG KARAKTER===")
	fmt.Print("Masukkan suatu kata=")
	fmt.Scanln(&input)

	hasil := len(input)

	fmt.Printf("jumlah karakter = %d", hasil)
	fmt.Println("")
	fmt.Printf("Karakter yang diinputkan = %s", input)
}
