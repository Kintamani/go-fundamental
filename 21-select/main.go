package main

import (
	"fmt"
	"time"
)

// memilih proses yang tercepat
func FungsiProses(id int, delay time.Duration) <-chan string {
	output := make(chan string)
	go func() {
		time.Sleep(delay * time.Second)
		output <- fmt.Sprintln("Process : ", id)
	}()
	return output
}

func main() {
	httpRequest1 := FungsiProses(1, 13)
	httpRequest2 := FungsiProses(2, 1)
	httpRequest3 := FungsiProses(3, 4)

	for index := 1; index < 3; index++ {
		select {
		case p1 := <-httpRequest1:
			fmt.Print(p1)
		case p2 := <-httpRequest2:
			fmt.Print(p2)
		case p3 := <-httpRequest3:
			fmt.Print(p3)
		}
	}
}
