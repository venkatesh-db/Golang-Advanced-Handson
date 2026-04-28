package main

import (
	"fmt"
	"time"

	//"os"
	"sync"
)

func cloud(wg *sync.WaitGroup) {

	fmt.Println("cloud server")
	wg.Done()
}

func server(wg *sync.WaitGroup) {

	fmt.Println("server")
	wg.Done()
}

func nowait() {

	fmt.Println("entry point nowait")
	time.Sleep(2 * time.Second)
	fmt.Println("exit point nowait")

}

func main2() {

	var wg sync.WaitGroup

	ch := make(chan int, 2)
	fmt.Println("created channel")
	//close(ch)
	fmt.Println("closed channel")
	<-ch
	fmt.Println("reading channel")

	go cloud(&wg)
	go server(&wg)

	for i := 0; i < 5; i++ {

		//go nowait()

		go func() {

			fmt.Println("entry point nowait")
			time.Sleep(2 * time.Second)
			fmt.Println("exit point nowait")

		}()
	}

	wg.Add(2)

	fmt.Println("hello server")

	//os.Open("file.txt")

	wg.Wait()

	fmt.Println("end of main")

}
