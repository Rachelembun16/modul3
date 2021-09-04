package main

import (
	"fmt"
)

func main() {
	var panjang int
	var lebar int

	fmt.Println("===HITUNGAN LUAS & KELILING===")
	fmt.Print("Masukkan panjang= ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan leber= ")
	fmt.Scanln(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("HASIL PERHITUNGAN")
	fmt.Printf("Luas = %d", luas)
	fmt.Println("")
	fmt.Printf("Keliling= %d", keliling)
}
