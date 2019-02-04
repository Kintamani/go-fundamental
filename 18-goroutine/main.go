package main

import (
	"fmt"
	"time"
)

func fungsiGoroutine() {
	fmt.Println("- INI FUNGSI GORUTINE")
}

func main() {
	go fungsiGoroutine()        // FUNGSI INI AKAN DIABAIKAN BILA TIDAK ADA FUNGSI SLEEP
	time.Sleep(1 * time.Second) // AKAN MENIDURKAN FUNGSI MAIN DALAM 1 DETIK
	fmt.Println("- INI FUNGSI MAIN")
}
