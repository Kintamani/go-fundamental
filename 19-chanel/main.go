package main

import "fmt"

func main() {
	// CHANNEL BUFFER (TIDAK HARUS MENUNGGU SAMPAI PENERIMANYA SIAP)
	value := make(chan string, 2)
	value <- "STRING SATU"
	value <- "STRING DUA"
	close(value)

	// fmt.Println(<-value)
	for data := range value {
		fmt.Println(data)
	}
}
