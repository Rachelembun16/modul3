package main

import (
	"fmt"
	"strconv"
)

func main() {
	var uang int
	fmt.Println("===KONVERSI RUPIAH===")
	fmt.Print("Masukkan Rupiah = ")
	fmt.Scan(&uang)

	fltUang := float64(uang)

	var strUang = strconv.Itoa(uang)
	usdUang := fltUang * 0.0000701662
	// intUang := int32(usdUang)

	arrUang := []rune(strUang)
	for i, j := 0, len(arrUang)-1; i < j; i, j = i+1, j-1 {
		arrUang[i], arrUang[j] = arrUang[j], arrUang[i]
	}

	var uangRev = string(arrUang)
	var money = ""
	i := 0

	for i < len(uangRev) {
		if i%3 == 0 && i != 0 {
			money = money + "." + string(uangRev[i])
			i = i + 1
		} else {
			money = money + string(uangRev[i])
			i = i + 1
		}

	}
	arrMoney := []rune(money)

	for i, j := 0, len(arrMoney)-1; i < j; i, j = i+1, j-1 {
		arrMoney[i], arrMoney[j] = arrMoney[j], arrMoney[i]
	}

	fmt.Printf("Konversi Rupiah= Rp%s,00", string(arrMoney))
	fmt.Println("")
	fmt.Printf("Konversi USD= $%.2f", usdUang)

}
