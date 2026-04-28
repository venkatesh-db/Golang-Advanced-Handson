
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Question struct {
	ID      int
	Content string
}

type Result struct {
	QuestionID int
	Status     string
}

// ---------------- CONFIG ----------------

type Config struct {
	WorkerCount int
	QueueSize   int
	TotalJobs   int
}

func LoadConfig() Config {
	return Config{
		WorkerCount: 200,     // tuned based on CPU
		QueueSize:   5000,    // backpressure buffer
		TotalJobs:   1000000, // 1 million
	}
}

// ---------------- PROCESSOR ----------------

type QuestionProcessor struct{}

func (p *QuestionProcessor) Process(ctx context.Context, q Question) Result {
	// Simulate real work:
	// - NLP scoring
	// - tagging
	// - DB insert
	time.Sleep(1 * time.Millisecond)

	return Result{
		QuestionID: q.ID,
		Status:     "processed",
	}
}

// ---------------- WORKER ----------------

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Question,
	results chan<- Result,
	processor *QuestionProcessor,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker-%d shutting down", id)
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			result := processor.Process(ctx, job)

			select {
			case results <- result:
			case <-ctx.Done():
				return
			}
		}
	}
}

// ---------------- PRODUCER ----------------

func produceJobs(ctx context.Context, jobs chan<- Question, total int) {
	defer close(jobs)

	for i := 1; i <= total; i++ {
		select {
		case jobs <- Question{
			ID:      i,
			Content: "System Design Question",
		}:
		case <-ctx.Done():
			return
		}
	}
}

// ---------------- RESULT HANDLER ----------------

func consumeResults(ctx context.Context, results <-chan Result) {
	count := 0
	start := time.Now()

	for {
		select {
		case res, ok := <-results:
			if !ok {
				log.Printf("all results processed: %d in %v", count, time.Since(start))
				return
			}

			count++
			if count%100000 == 0 {
				log.Printf("progress: %d processed", count)
				_ = res // in real system → DB write / analytics
			}

		case <-ctx.Done():
			return
		}
	}
}

// ---------------- MAIN ----------------

func main() {

	cfg := LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	jobs := make(chan Question, cfg.QueueSize)
	results := make(chan Result, cfg.QueueSize)

	processor := &QuestionProcessor{}

	var workerWG sync.WaitGroup

	// Start workers
	for i := 1; i <= cfg.WorkerCount; i++ {
		workerWG.Add(1)
		go worker(ctx, i, jobs, results, processor, &workerWG)
	}

	// Start producer
	go produceJobs(ctx, jobs, cfg.TotalJobs)

	// Close results after workers done
	go func() {
		workerWG.Wait()
		close(results)
	}()

	// Start result consumer
	go consumeResults(ctx, results)

	// Wait for shutdown signal
	select {
	case sig := <-sigChan:
		log.Println("shutdown signal received:", sig)
		cancel()

	case <-ctx.Done():
	}

	time.Sleep(2 * time.Second) // allow graceful drain
	fmt.Println("system exited cleanly")
}