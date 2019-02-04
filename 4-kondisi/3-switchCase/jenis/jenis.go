package jenis

import "fmt"

func Pertama(X int) string {
	value := ""
	fmt.Println("NILAI = ", X)

	switch X {
	case 1:
		value = "NILAI D"
	case 5:
		value = "NILAI C"
	case 8:
		value = "NILAI B"
	case 10:
		value = "NILAI A"
	default:
		value = "TIDAK DITEMUKAN NILAI"
	}
	return value
}

func Kedua(X int) string {
	value := ""
	fmt.Println("NILAI = ", X)

	switch {
	case X == 1:
		value = "NILAI D"
	case X == 5:
		value = "NILAI C"
	case X == 8:
		value = "NILAI B"
	case X == 10:
		value = "NILAI A"
	default:
		value = "TIDAK DITEMUKAN NILAI"
	}
	return value
}
