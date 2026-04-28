package main

import (
	"fmt"
	"time"
)

/*

worker → tries to receive
→ no sender
→ runtime.gopark()  (goroutine sleeps)

main → sends data
→ runtime.goready() (wake worker)

*/

func worker(ch chan string) {
	fmt.Println("worker: waiting for data (will park)")
	msg := <-ch // ❗ BLOCK → internally runtime.gopark()
	fmt.Println("worker received:", msg)
}

func main() {
	ch := make(chan string)

	go worker(ch)

	time.Sleep(2 * time.Second) // give time to park

	fmt.Println("main: sending data")
	ch <- "hello" // ❗ unpark worker

	time.Sleep(1 * time.Second)
}
