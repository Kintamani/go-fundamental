package main

import (
	"basic/6-fungsi/jenis"
	"fmt"
)

func main() {
	// FUNCTION SEDERHANA
	x := 5
	y := 5
	valName := "Jaka Widada"
	valUmur := jenis.AddUmur(x, y)
	valKeterangan := jenis.Keterangan(valName, valUmur)
	fmt.Print(valKeterangan)
	fmt.Println("=================================")

	// FUNGSI SEBAGAI VALUE
	valFungsi := func(index1, index2 int) int {
		return index1 + index2
	}
	fmt.Println("- FUNGSI SEBAGAI VALUE = ", valFungsi(valUmur, 10))
	fmt.Println("=================================")

	// FUNGSI CLOSURE
	closeure := jenis.Closeure("- NAMA SATU")
	fmt.Println(closeure("- NAMA DUA"))
	fmt.Println("=================================")

	//FUNGSI SEBAGAI PARAMETER
	functionPertama := func(valuePertama string) bool { //HARUS SAMA DENGAN FUNCTION YANG DIDEKLARASI DIBAWAH
		return valuePertama == "ASAL"
	}
	panggil := parameter("JANGAN ASAL", functionPertama)
	fmt.Print(panggil)
	fmt.Println("=================================")
}

func parameter(valueKedua string, functionKedua func(string) bool) string {
	if functionKedua(valueKedua) {
		return fmt.Sprintln("- BENAR")
	}
	return fmt.Sprintln("- SALAH")
}
