package main

import (
	//"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var uang int
	fmt.Println("===KONVERSI RUPIAH===")
	fmt.Print("Masukkan Rupiah = ")
	fmt.Scan(&uang)
	newText1 := ""
	//buf := &bytes.Buffer{}
	var strUang = strconv.Itoa(uang)

	arrUang := []rune(strUang)
	for i, j := 0, len(arrUang)-1; i < j; i, j = i+1, j-1 {
		arrUang[i], arrUang[j] = arrUang[j], arrUang[i]
	}

	var uangRev = string(arrUang)

	i := len(uangRev)
	for i >= 0 {
		//if i%3 == 0 && i != 0 {
		newText1 = strings.Replace(uangRev, "", ".", 2)
		//}
		i = i - 3
	}
	fmt.Print(newText1)
	// var uangRev =
	//var zero = strings.Count(strUang, "0")
	//var zero1 = strings.Split(strUang, "")
	// for i := 0; i < len(strUang)-1; i++ {
	// 	if i%3 == 0 {
	// 		// var angka = string(strUang[i])
	// 		var money = strings.Join(strUang, ",")
	// 		fmt.Printf("Rp %s,00", money)
	// 	} else {
	// 		fmt.Print(strUang[i])
	// 	}
	// }

	//var fltUang, err = strconv.ParseFloat(uang, 32)

	// fltuang := float32(uang)
	// fmt.Printf("f is %f", fltuang)

	// if err == nil {
	// fmt.Printf("Rp %s,00", money)
	// }
}
