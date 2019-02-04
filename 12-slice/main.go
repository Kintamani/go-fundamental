package main

import "fmt"

func main() {
	value := []int{2, 3, 2, 4, 1}
	for i, data := range value {
		fmt.Println("DATA KEY = ", i, ",DENGAN VALUE = ", data)
	}

	fmt.Println("----------------==PROSES APPEND==----------------")
	// MENAMBAHKAN SLICE
	diAppend := append(value, 100)
	for i, data := range diAppend {
		fmt.Println("DATA KEY = ", i, ",DENGAN VALUE = ", data)
	}
}
