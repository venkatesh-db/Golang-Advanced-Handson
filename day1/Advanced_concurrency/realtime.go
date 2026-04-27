package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SeatResult struct {
	Service string
	Data    string
	Err     error
}

func fetchFromService(ctx context.Context, service string, delay time.Duration) SeatResult {
	select {
	case <-time.After(delay):
		return SeatResult{
			Service: service,
			Data:    "available",
			Err:     nil,
		}
	case <-ctx.Done():
		return SeatResult{
			Service: service,
			Err:     ctx.Err(),
		}
	}
}

func worker(ctx context.Context, jobs <-chan string, results chan<- SeatResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for service := range jobs {
		result := fetchFromService(ctx, service, 800*time.Millisecond)
		results <- result
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	services := []string{"seat-service", "pricing-service", "quota-service"}

	jobs := make(chan string, len(services))
	results := make(chan SeatResult, len(services))

	var wg sync.WaitGroup

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}

	for _, svc := range services {
		jobs <- svc
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Println("error from", res.Service, ":", res.Err)
			continue
		}
		fmt.Println("response from", res.Service, ":", res.Data)
	}
}
