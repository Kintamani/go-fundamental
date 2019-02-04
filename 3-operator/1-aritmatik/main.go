package main

import "fmt"

func main() {
	x := 10
	y := 5

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("FUNGSI ARITMATIKANYA ADALAH :")

	tambah := x + y
	fmt.Println("x + y = ", tambah)

	kurang := x - y
	fmt.Println("x - y = ", kurang)

	kali := x * y
	fmt.Println("x * y = ", kali)

	bagi := x / y
	fmt.Println("x / y = ", bagi)

	modulus := x % y
	fmt.Println("x % y = ", modulus)
}
