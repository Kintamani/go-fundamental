package main

import "fmt"

func loop(mulai, input int) {
	for index := mulai; index < input; index++ {
		defer fmt.Println("VALUE KE-", index)
	}
	fmt.Println("==============-TANDPA DEFER-==============")
	for index := mulai; index < input; index++ {
		fmt.Println("VALUE KE-", index)
	}
	fmt.Println("==============-DENGAN DEFER-==============")
}
func main() {
	loop(1, 5)
}
