package main

import "fmt"

type users struct {
	name string
	age  int
}

func main() {
	var struk []users
	struk = []users{
		users{
			name: "JAKA",
			age:  12,
		},
		users{
			name: "PRABAWA",
			age:  12,
		},
	}

	for _, data := range struk {
		fmt.Println(data)
	}

	fmt.Println("----------------==PROSES APPEND==----------------")
	newStruk := append(struk, users{
		name: "DILDO",
		age:  12,
	})
	for _, data := range newStruk {
		fmt.Println(data)
	}
}
