package main

import "fmt"

func main() {
	benar := true
	salah := false

	AND := benar && salah
	fmt.Println("benar && salah = ", AND)

	OR := benar || salah
	fmt.Println("benar || salah", OR)

	NEGASIBENAR := !benar
	fmt.Println("! benar = ", NEGASIBENAR)

	NEGASISALAH := !salah
	fmt.Println("! salah = ", NEGASISALAH)
}
