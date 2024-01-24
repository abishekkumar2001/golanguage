package main

import (
	"fmt"
	"time"
)

func messages1(msg1 string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg1, "=", i)
		time.Sleep(time.Second)
	}
}

func main() {
	//Goroutines are called as light weight threads
	go messages("Threads...") //The prefix go indicates Multithreading/Goroutines
	go messages1("Threads1...")
	time.Sleep(6 * time.Second)
	go messages("Threads...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!!")
}

func messages(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
		time.Sleep(time.Second)
	}
}
