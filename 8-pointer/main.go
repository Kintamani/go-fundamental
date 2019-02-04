package main

import "fmt"

func main() {
	value := "HELLO"
	var stringPointer *string
	stringPointer = &value

	fmt.Println("- ALAMAT RAM : ", stringPointer)
	fmt.Println("- ISI :", value)
	fmt.Println("====================================")

	tanpaPointer(value)
	fmt.Println("- TANPA POINTER : ", value)
	denganPointer(&value)
	fmt.Println("- DENGAN POINTER : ", value)
}

func tanpaPointer(val string) {
	val = val + " DUNIA"
}

func denganPointer(val *string) {
	*val = *val + " DUNIA"
}
