package main

import "fmt"

type users struct {
	name string
	age  int
}

func main() {
	mapValue := make(map[int]users)
	mapValue[1] = users{name: "nama satu", age: 12}
	mapValue[2] = users{name: "nama dua", age: 13}
	mapValue[3] = users{name: "nama tiga", age: 14}
	mapValue[4] = users{name: "nama empat", age: 15}

	for _, data := range mapValue {
		fmt.Println(data.name)
	}
}
