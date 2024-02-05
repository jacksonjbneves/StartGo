package main

import (
	"fmt"
	"time"
)

func MultThread(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Line[%s] Thread %d \n", msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("---[Start Mult-Thread's]---")
	go MultThread("Thread One") //Thread One
	go MultThread("Thread Two") //Thread Two

	fmt.Println("---[Thread Single]---")
	MultThread("Thread Single")
}
