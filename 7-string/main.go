package main

import (
	"fmt"
	"strconv"
	"strings"
)

// MANIPULATION STRING
func main() {
	valueSatu := "Value String"
	valueDua := "10"
	fmt.Println("- ", valueSatu)

	// MEMBESARKAN STRING
	stringUpper := strings.ToUpper(valueSatu)
	fmt.Println("- ", stringUpper)

	// MENGECILKAN STRING
	stringLower := strings.ToLower(valueSatu)
	fmt.Println("- ", stringLower)

	// MENGECEK APAKAH ADA STRING DENGAN ISI "golang"
	resultContain := strings.Contains(valueSatu, "golang")
	fmt.Println("- ", resultContain)

	// MENGUBAH DARI STRING KE ARRAY
	resultSplit := strings.Split(valueSatu, " ")
	for _, data := range resultSplit {
		fmt.Println("- ARRAY : ", data)
	}

	// MENGUBAH DARI STRING KE INT
	resultConvInt, err := strconv.Atoi(valueDua)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("- NILAI VALUE DUA : ", resultConvInt)
	fmt.Println("- 10 * 5 =", resultConvInt*5)

	// MENGUBAH DARI INT KE STRING
	resultContStr := strconv.Itoa(10)
	fmt.Println("- NILAI 10 YANG TELAH DI UBAH MENJADI STRING : ", resultContStr)
}
