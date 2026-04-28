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
	Workers       int
	QueueSize     int
	BatchSize     int
	TotalJobs     int
	WriterCount   int
	DBConcurrency int
}

func loadConfig() Config {
	return Config{
		Workers:       runtime.NumCPU() * 4,
		QueueSize:     20000,
		BatchSize:     500,
		TotalJobs:     1_000_000,
		WriterCount:   4,
		DBConcurrency: 8, // simulate DB connection pool
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

	// DB semaphore (critical for real systems)
	dbSem := make(chan struct{}, cfg.DBConcurrency)

	// ---------------- WORKERS ----------------

	for i := 0; i < cfg.Workers; i++ {
		workerWG.Add(1)
		go func() {
			defer workerWG.Done()
			for job := range jobs {

				// simulate CPU work
				process(job.ID)

				results <- Result{ID: job.ID}
			}
		}()
	}

	// ---------------- WRITERS (BATCH + DB CONTROL) ----------------

	for i := 0; i < cfg.WriterCount; i++ {
		writerWG.Add(1)
		go func() {
			defer writerWG.Done()

			batch := make([]Result, 0, cfg.BatchSize)

			for res := range results {
				batch = append(batch, res)

				if len(batch) >= cfg.BatchSize {
					writeBatch(batch, dbSem)
					batch = batch[:0]
				}
			}

			if len(batch) > 0 {
				writeBatch(batch, dbSem)
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

// ---------------- CPU WORK ----------------

func process(id int) {
	// simulate CPU-bound work
	for i := 0; i < 50; i++ {
		_ = id * i
	}
}

// ---------------- DB WRITE ----------------

func writeBatch(batch []Result, sem chan struct{}) {
	// acquire DB connection
	sem <- struct{}{}

	// simulate DB latency (realistic)
	time.Sleep(2 * time.Millisecond)

	// release connection
	<-sem
}
