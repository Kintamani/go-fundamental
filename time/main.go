package main

import (
	"fmt"
	"time"
)

func main() {
	sekarang := time.Now()
	tahun := sekarang.Year()
	bulan := sekarang.Month()
	hari := sekarang.Day()

	fmt.Println(sekarang)
	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(hari)

	fmt.Println("================================")

	// KONFERSI STRING TO TIME
	stringTime := "January 12, 2019"
	firstStringTime := "January 1, 2000"
	konfToString, err := time.Parse(firstStringTime, stringTime)

	if err != nil {
		err.Error()
	}
	fmt.Println(konfToString.Year())

	fmt.Println("================================")

	// MEMBUAT WAKTU SECARA MANUAL
	makeTime1 := time.Date(2019, 2, 4, 1, 0, 0, 0, time.UTC)

	fmt.Println(makeTime1)

	fmt.Println("================================")
	// MENGECEK APAKAH TIME 1 SAM DENGAN TIME 2 (HASIL BOOLEAN)
	makeTime2 := time.Date(2019, 2, 4, 1, 0, 0, 0, time.UTC)
	isEquals := makeTime1.Equal(makeTime2)

	fmt.Println(isEquals)

	fmt.Println("================================")

	// KONFERSI TIME TO STRING
	stringDate := "2004-02-01"
	resString := sekarang.Format(stringDate)

	fmt.Println(resString)
}
