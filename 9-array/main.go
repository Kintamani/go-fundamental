package main

import "fmt"

func main() {
	// INTEGER ARRAY
	var valueSatu [4]int
	valueSatu[0] = 1
	valueSatu[1] = 2
	valueSatu[2] = 3
	valueSatu[3] = 4
	for _, data := range valueSatu {
		fmt.Println(data)
	}

	fmt.Println("=========================")
	// STRING ARRAY
	valueDua := [...]string{"name", "address", "age"}
	for _, data := range valueDua {
		fmt.Println(data)
	}

	fmt.Println("=========================")
	// MULTI ARRAY
	valueTiga := [2][4]int{{2, 3, 4, 1}, {2, 3, 4, 5}}
	fmt.Println(valueTiga[0][2])
}
