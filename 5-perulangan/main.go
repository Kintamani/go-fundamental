package main

import (
	"basic/5-perulangan/jenis"
	"fmt"
)

func main() {
	var nilai int
	nilai = 5

	//pertama
	var pertama = jenis.Pertama(nilai)
	fmt.Println(pertama)

	//kedua
	var kedua = jenis.Kedua(nilai)
	fmt.Println(kedua)

	//ketiga
	var ketiga = jenis.Ketiga(nilai)
	fmt.Println(ketiga)

}
