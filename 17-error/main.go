package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Fungsi(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("PEMBAGI TIDAK BOLEH = 0")
	} else {
		hasil := x / y
		return hasil, nil
	}
}

func main() {
	// HANDLE ERROR PERTAMA
	valueString := "ADA"

	i, err := strconv.Atoi(valueString)
	if err != nil {
		fmt.Println("ERROR !,", err.Error())
	} else {
		fmt.Println("STRING TELAH DIRUBAH MENJADI INT", i)
	}

	// HANDLE ERROR KEDUA
	hasil, err := Fungsi(1, 0)

	if err != nil {
		fmt.Println("ERROR !,", err.Error())
	} else {
		fmt.Println("BERHASIL !, HASILNYA = ", hasil)
	}

	// HANDLE ERROR KETIGA
	if hasil2, err := Fungsi(2, 2); err != nil {
		fmt.Println("ERROR !,", err.Error())
	} else {
		fmt.Println("BERHASIL !, HASILNYA = ", hasil2)
	}
}
