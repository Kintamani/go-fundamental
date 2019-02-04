package jenis

import "fmt"

func AddUmur(x int, y int) int {
	return x + y
}

func Keterangan(name string, umur int) string {
	// fmt.Sprintf == menggabungkan dua variabel tanpa ketentuan
	return fmt.Sprintf("- Nama : %v\n- Umur : %v\n", name, umur)
}

//FUNCTION CLOSEURE
func Closeure(nameSatu string) func(string) string {
	return func(nameDua string) string {
		return fmt.Sprintf("%s DAN %s", nameSatu, nameDua)
	}
}
