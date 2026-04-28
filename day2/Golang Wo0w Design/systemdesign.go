package main

import (
	"fmt"
)

func collectlogs(processing string, channel chan int) {

	fmt.Println("Hi servcie1", processing, <-channel)
}

func main() {

	var channel chan int = make(chan int)

	// How to create goroutine
	go collectlogs("serivce1", channel) // register to  gororutine schduler

	// How to execute goroutines
	// time.Sleep(1)

	// How to communciate gororutines - channel
	channel <- 1

}

//System design Interview

// How to create goroutine
// How to execute goroutines
// How to communciate gororutines
