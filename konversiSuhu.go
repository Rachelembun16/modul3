package main

import (
	"fmt"
)

func main() {
	var pilih int

	fmt.Println("===KONVERSI SUHU===")
	fmt.Println("")
	fmt.Println("1. Celcius ==> Kelvin")
	fmt.Println("2. Celcius ==> Fahrenheit")
	fmt.Println("3. Kelvin ==> Celcius")
	fmt.Println("4. Kelvin ==> Fahtenheit")
	fmt.Println("5. Fahrenheit ==> Kelvin")
	fmt.Println("6. Fahrenheit ==> Celcius")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		var celcius float32
		fmt.Print("Masukkan suhu dalam Celcius: ")
		fmt.Scan(&celcius)

		hasil := celcius + 273.15

		fmt.Println(fmt.Sprintf("Hasil konversi= %f `K", hasil))
	case 2:
		var celcius float32
		fmt.Print("Masukkan suhu dalam Celcius: ")
		fmt.Scan(&celcius)

		hasil := (celcius * 9 / 5) + 32
		fmt.Println(fmt.Sprintf("Hasil konversi = %f `F", hasil))
	case 3:
		var kelvin float32
		fmt.Print("Masukkan suhu dalam Kelvin: ")
		fmt.Scan(&kelvin)

		hasil := kelvin - 273.15

		fmt.Print(fmt.Sprintf("Hasil konversi = %f `C", hasil))
	case 4:
		var kelvin float32
		fmt.Printf("Masukkan suhu dalam Kelvin: ")
		fmt.Scan(&kelvin)

		hasil := (kelvin-273.15)*9/5 + 32

		fmt.Printf(fmt.Sprintf("Hasil konversi = %f `F", hasil))
	case 5:
		var fahrenheit float32
		fmt.Printf("Masukkan suhu dalam Fahrenheit: ")
		fmt.Scan(&fahrenheit)

		hasil := (fahrenheit-32)*5/9 + 273.15

		fmt.Printf(fmt.Sprintf("Hasil konversi = %f `K", hasil))
	case 6:
		var fahrenheit float32
		fmt.Printf("Masukkan suhu dalam Fahrenheit: ")
		fmt.Scan(&fahrenheit)

		hasil := (fahrenheit - 32) * 5 / 9

		fmt.Printf(fmt.Sprintf("Hasil konversi = %f `C", hasil))
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}
