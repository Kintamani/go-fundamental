package main

import "fmt"

func main() {
	x := 10
	y := 5

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("FUNGSI IRATIONAL ADALAH :")

	samaDengan := x == y
	fmt.Println("x == y = ", samaDengan)

	tidak := x != y
	fmt.Println("x != y = ", tidak)

	kurang := x < y
	fmt.Println("x < y = ", kurang)

	kurangSama := x <= y
	fmt.Println("x <= y = ", kurangSama)

	lebih := x > y
	fmt.Println("x > y = ", lebih)

	lebihSama := x >= y
	fmt.Println("x >= y = ", lebihSama)
}
