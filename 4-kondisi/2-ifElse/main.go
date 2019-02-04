package main

import "fmt"

func main() {
	X := 5
	Y := 5

	fmt.Println("X = ", X)
	fmt.Println("Y = ", Y)

	if X > Y {
		fmt.Println("X > Y")
	} else if X < Y {
		fmt.Println("X < Y")
	} else {
		fmt.Println("X = Y")
	}
}
