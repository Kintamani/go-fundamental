package main

import "fmt"

func main() {
	var mapValue map[int]string
	mapValue = make(map[int]string)

	mapValue[0] = "nama1"
	mapValue[1] = "nama2"
	mapValue[2] = "nama3"
	mapValue[3] = "nama4"
	mapValue[4] = "nama5"

	for key, data := range mapValue {
		fmt.Println("KEY : ", key, ",DATA : ", data)
	}

	// CHECK MAP
	stringnya, boolean := mapValue[4]
	if !boolean {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println("ADA DENGAN VALUE : ", stringnya)
	}
}
