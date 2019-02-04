package jenis

import "fmt"

func Pertama(x int) string {
	for index := 0; index < x; index++ {
		fmt.Println("NILAI X = ", x)
	}
	xx := "---------------------------------------"
	return xx
}

func Kedua(x int) string {
	index := 0
	for index < x {
		index++
		fmt.Println("halo")
	}
	xx := "---------------------------------------"
	return xx
}

func Ketiga(x int) string {
	index := 0

	for index < x {
		index++

		if index%2 == 0 {
			fmt.Println("GANJIL")
		} else {
			fmt.Println("GENAP")
		}
	}
	xx := "---------------------------------------"
	return xx
}
