package main

import "fmt"

type value struct {
	name string
	age  int
}

type valueDua struct {
	name string
	age  int
}

// STRUCT SEBAGAI PARAMETER FUCNTION
func parameter(satu value, dua valueDua) {
	fmt.Printf("- NAMA SAYA %s, UMUR SAYA %d\n", satu.name, satu.age)
	fmt.Printf("- NAMA SAYA %s, UMUR SAYA %d", dua.name, dua.age)
}

func main() {
	// DYNAMIC DECLARATION

	var val value
	val.name = "JAKA WIDADA"
	val.age = 13

	// STATIC DECLARATION

	val2 := valueDua{
		name: "PRABAWA SUBIANTA",
		age:  12,
	}
	fmt.Println("----------------==DYNAMIC DECLARATIO==----------------")
	fmt.Printf("- NAMA SAYA %s, UMUR SAYA %d\n", val.name, val.age)
	fmt.Println("----------------==STATIC DECLARATION==----------------")
	fmt.Printf("- NAMA SAYA %s, UMUR SAYA %d\n", val2.name, val2.age)

	// MEMASUKAN STRUCT KE PARAMETER:
	fmt.Println("----------------==STRUCT PARAMETER==----------------")
	parameter(val, val2)
}
