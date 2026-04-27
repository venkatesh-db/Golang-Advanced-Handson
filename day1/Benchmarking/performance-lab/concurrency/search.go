package concurrency

import (
	"context"
	"sync"
	"time"
)

type Result struct {
	ID int
}

func process(ctx context.Context, id int) (Result, error) {
	select {
	case <-time.After(5 * time.Millisecond):
		return Result{ID: id}, nil
	case <-ctx.Done():
		return Result{}, ctx.Err()
	}
}

// unbounded concurrency (not recommended)
func ProcessUnbounded(ctx context.Context, ids []int) []Result {
	var wg sync.WaitGroup
	results := make([]Result, len(ids))

	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			r, _ := process(ctx, id)
			results[i] = r
		}(i, id)
	}

	wg.Wait()
	return results
}

// bounded worker pool
func ProcessWithPool(ctx context.Context, ids []int, workers int) []Result {
	jobs := make(chan int)
	results := make([]Result, len(ids))

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs {
				r, _ := process(ctx, id)
				results[id] = r
			}
		}()
	}

	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	wg.Wait()
	return results
}
