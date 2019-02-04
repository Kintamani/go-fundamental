package main

import (
	"fmt"
	"sync"
	"time"
)

func Proses(id int, delay time.Duration, waitGroup *sync.WaitGroup) <-chan string {
	output := make(chan string)
	go func() {
		defer waitGroup.Done()

		time.Sleep(time.Second * delay)
		output <- fmt.Sprintln("Process ", id, ",done")
	}()
	return output
}

func main() {
	var waitGroup sync.WaitGroup
	// valueDone := make(chan bool)

	waitGroup.Add(3)

	// go func() {
	// 	fmt.Println("Gproutine")
	// 	// valueDone <- true
	// 	waitGroup.Done()
	// }()

	res1 := <-Proses(1, 3, &waitGroup)
	res2 := <-Proses(2, 2, &waitGroup)
	res3 := <-Proses(3, 4, &waitGroup)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// <-valueDone
	waitGroup.Done()
}
