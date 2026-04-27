package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	jobs := []int{1, 2, 3}

	for _, job := range jobs {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			time.Sleep(500 * time.Millisecond)
			fmt.Println("processed job", j)
		}(job)
	}

	wg.Wait()
	fmt.Println("all jobs completed")
}
