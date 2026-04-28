package main

import (
	"fmt"
	"time"
)

func consumer(id int, ch chan int) {
	for {
		val := <-ch // ❗ blocks if empty
		fmt.Printf("consumer %d got %d\n", id, val)
	}
}

func main() {
	ch := make(chan int)

	go consumer(1, ch)
	go consumer(2, ch)

	time.Sleep(2 * time.Second)

	for i := 1; i <= 5; i++ {
		ch <- i // wakes one waiting goroutine
	}

	time.Sleep(2 * time.Second)
}
