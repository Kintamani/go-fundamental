package main

import "fmt"

const A string = "INI KONSTANTA A (GLOBAL KONSTANTA)"

func main() {
	fmt.Println(A)

	const B int = 10
	fmt.Printf("INI KONSTANTA B (LOCAL KONSTANTA) DENGAN ISI = %v\n", B)

	z := 100 / B
	fmt.Println("HASIL Z ADALAH 100/B = ", z)
}
