package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// ---------------- DOMAIN ----------------

type Question struct {
	ID int
}

type Result struct {
	ID int
}

// ---------------- CONFIG ----------------

type Config struct {
	Workers     int
	QueueSize   int
	BatchSize   int
	TotalJobs   int
	WriterCount int
}

func loadConfig() Config {
	return Config{
		Workers:     runtime.NumCPU() * 4,
		QueueSize:   10000,
		BatchSize:   1000,
		TotalJobs:   1_000_000,
		WriterCount: 4,
	}
}

// ---------------- MAIN ----------------

func main() {
	cfg := loadConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	jobs := make(chan Question, cfg.QueueSize)
	results := make(chan Result, cfg.QueueSize)

	var workerWG sync.WaitGroup
	var writerWG sync.WaitGroup

	// ---------------- WORKERS ----------------

	for i := 0; i < cfg.Workers; i++ {
		workerWG.Add(1)
		go func() {
			defer workerWG.Done()
			for job := range jobs {
				results <- Result{ID: job.ID}
			}
		}()
	}

	// ---------------- WRITERS (BATCH) ----------------

	for i := 0; i < cfg.WriterCount; i++ {
		writerWG.Add(1)
		go func() {
			defer writerWG.Done()

			batch := make([]Result, 0, cfg.BatchSize)

			for res := range results {
				batch = append(batch, res)

				if len(batch) == cfg.BatchSize {
					flush(batch)
					batch = batch[:0]
				}
			}

			// flush remaining
			if len(batch) > 0 {
				flush(batch)
			}
		}()
	}

	// ---------------- PRODUCER ----------------

	go func() {
		for i := 1; i <= cfg.TotalJobs; i++ {
			jobs <- Question{ID: i}
		}
		close(jobs)
	}()

	// ---------------- CLOSE RESULTS ----------------

	go func() {
		workerWG.Wait()
		close(results)
	}()

	// ---------------- WAIT ----------------

	writerWG.Wait()

	log.Printf("✅ processed %d jobs in %v\n", cfg.TotalJobs, time.Since(start))
}

// ---------------- STORAGE (SIMULATION) ----------------

func flush(batch []Result) {
	// Replace with:
	// - DB batch insert
	// - Kafka publish
	// - File write
}
