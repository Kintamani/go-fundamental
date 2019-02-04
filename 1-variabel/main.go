package main

import (
	"fmt"
)

func main() {
	//static type declaration
	var x int
	x = 10

	var y float64
	y = 5.5

	fmt.Println(x)
	fmt.Printf("x = %T\n", x)
	fmt.Println(y)
	fmt.Printf("y = %T\n", y)

	fmt.Println("===========")
	//dynamic type declaration
	z1 := 10
	z2 := 5.5

	fmt.Println(z1)
	fmt.Printf("z1 = %T\n", z1)
	fmt.Println(z2)
	fmt.Printf("z2 = %T\n", z2)
}
