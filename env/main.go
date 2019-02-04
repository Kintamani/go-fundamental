package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("MYSQL", "mysql.two.host")

	mysql := os.Getenv("MYSQL")
	fmt.Println("MySQL CONFIG : ", mysql)
}
