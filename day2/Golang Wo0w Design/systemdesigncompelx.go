package main

import (
	"fmt"
	"sync"
)

func collectlogs(processing string, channel chan string, p *sync.WaitGroup) {

	fmt.Println("Hi collectlogs")

	fmt.Println("processing", <-channel)

	p.Done()

}

func producelogs(processing string, channel chan string, p *sync.WaitGroup) {

	fmt.Println("Hi producelogs start ")

	channel <- " failed to connect db"

	fmt.Println("Hi producelogs end")

	p.Done()
}

func main() {

	var wg sync.WaitGroup

	var channel chan string = make(chan string)
	// log --> tell logs to store in a  filename

	// How to create goroutine
	go collectlogs("serivce1", channel, &wg) // register to  gororutine schduler
	go producelogs("serivce1", channel, &wg) // register to  gororutine schduler

	wg.Add(1)
	wg.Add(1)

	// How to communciate gororutines - channel
	wg.Wait() // main is waiting communciation two groroutines

	fmt.Println(" main ends")
}
