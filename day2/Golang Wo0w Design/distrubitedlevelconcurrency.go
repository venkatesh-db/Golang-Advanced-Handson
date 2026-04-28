package main

import (
	"fmt"
	"time"
)

func collectlogs(processing string) {

	fmt.Println("Hi collectlogs start ")

	fmt.Println("Hi collectlogs ends ")
}

func producelogs(processing int) {

	fmt.Println("Hi producelogs start ")

	fmt.Println("Hi producelogs end")

}

func main() {

	for i := 0; i < 5; i++ {

		go func(processing int) {

			fmt.Println("Hi collectlogs start ", i)

			fmt.Println("Hi collectlogs ends ", i)
		}(i)

	}

	for j := 0; j < 3; j++ {

		go func(processing int) {

			fmt.Println("Hi producelogs start ", j)

			fmt.Println("Hi producelogs end", j)

		}(j)

	}

	time.Sleep(2 * time.Second)
	fmt.Println(" main ends")
}




