package main

import (
	"basic/4-kondisi/3-switchCase/jenis"
	"fmt"
)

func main() {
	var x int
	x = 10

	//PERTAMA
	fmt.Println("============================================")
	var pertama = jenis.Pertama(x)
	fmt.Println("- HASIL SWITCH CASE PERTAMA =", pertama)

	//KEDUA
	fmt.Println("============================================")
	var kedua = jenis.Kedua(x)
	fmt.Println("- HASIL SWITCH CASE KEDUA =", kedua)

	//KETIGA
	fmt.Println("============================================")
	var X interface{}
	X = 12

	fmt.Println("NILAI = ", X)

	switch X.(type) {
	case int:
		fmt.Println("TIPE DATA INT")
	case string:
		fmt.Println("TIPE DATA STRING")
	case float32:
		fmt.Println("TIPE DATA FLOAT23")
	case float64:
		fmt.Println("TIPE DATA FLOAT64")
	default:
		fmt.Println("TIDAK DITEMUKAN TYPE DATA")
	}

}
