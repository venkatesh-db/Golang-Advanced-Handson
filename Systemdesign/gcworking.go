package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMem(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%s → Alloc = %v MB | TotalAlloc = %v MB | NumGC = %v\n",
		tag,
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.NumGC,
	)
}

func main() {

	printMem("START")

	for i := 1; i <= 3; i++ {

		// allocate large memory
		data := make([][]byte, 1000)
		for j := range data {
			data[j] = make([]byte, 1024*50) // 50KB
		}

		printMem(fmt.Sprintf("After Allocation %d", i))

		// remove reference → eligible for GC
		data = nil

		// force GC (for learning)
		runtime.GC()

		printMem(fmt.Sprintf("After GC %d", i))

		time.Sleep(1 * time.Second)
	}
}
