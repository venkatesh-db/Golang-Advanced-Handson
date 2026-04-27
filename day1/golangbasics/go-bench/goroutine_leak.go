package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 👈 IMPORTANT
	"runtime"
	"time"
)

/*

pprof shows where goroutines are stuck — that is your leak source


go run goroutine_leak.go

http://localhost:6060/debug/pprof/goroutine?debug=2

go tool pprof http://localhost:6060/debug/pprof/goroutine

top

list leakyWorker



*/

func leakyWorker(ch chan int) {
	for {
		<-ch // 🚨 leak (blocked forever)
	}
}

func main() {
	ch := make(chan int)

	// start leak
	for i := 0; i < 1000; i++ {
		go leakyWorker(ch)
	}

	// start pprof server
	go func() {
		fmt.Println("pprof running at http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// monitor
	for {
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
