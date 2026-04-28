package main

import (
	"fmt"
//	"time"
)

func server1() {
	fmt.Println("server")
	var daysmiles int = 5
	var smile *int = new(int)

	fmt.Printf("daysmiles is %d\n", daysmiles)
	fmt.Printf("address\n", smile, &smile)
}

func main2() {

	fmt.Println("entry main")

	ch := make(chan string)

	go server1()

	fmt.Println("before smile")

	//time.Sleep(1 * time.Second)

	 ch <- "smile"

	fmt.Println("end smile")

}
