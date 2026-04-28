package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------------- TASK ----------------

type Task struct {
	ID int
}

// ---------------- GLOBAL QUEUE ----------------

type GlobalQueue struct {
	mu    sync.Mutex
	queue []Task
}

func (g *GlobalQueue) Pop() (Task, bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.queue) == 0 {
		return Task{}, false
	}

	t := g.queue[0]
	g.queue = g.queue[1:]
	return t, true
}

// ---------------- LOCAL QUEUE ----------------

type LocalQueue struct {
	mu    sync.Mutex
	queue []Task
}

func (l *LocalQueue) Push(t Task) {
	l.mu.Lock()
	l.queue = append(l.queue, t)
	l.mu.Unlock()
}

func (l *LocalQueue) Pop() (Task, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.queue) == 0 {
		return Task{}, false
	}

	t := l.queue[0]
	l.queue = l.queue[1:]
	return t, true
}

// steal half (like Go runtime)
func (l *LocalQueue) StealHalf() []Task {
	l.mu.Lock()
	defer l.mu.Unlock()

	n := len(l.queue)
	if n <= 1 {
		return nil
	}

	half := n / 2
	stolen := make([]Task, half)
	copy(stolen, l.queue[:half])
	l.queue = l.queue[half:]
	return stolen
}

// ---------------- WORKER ----------------

func worker(id int, local *LocalQueue, global *GlobalQueue, allLocals []*LocalQueue, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// 1. Try local queue
		if task, ok := local.Pop(); ok {
			process(id, task)
			continue
		}

		// 2. Try global queue
		if task, ok := global.Pop(); ok {
			process(id, task)
			continue
		}

		// 3. Work stealing
		stole := false
		for _, other := range allLocals {
			if other == local {
				continue
			}

			tasks := other.StealHalf()
			if len(tasks) > 0 {
				fmt.Printf("🔥 Worker %d stole %d tasks\n", id, len(tasks))

				for _, t := range tasks {
					local.Push(t)
				}
				stole = true
				break
			}
		}

		if !stole {
			// nothing left anywhere → exit
			return
		}
	}
}

// ---------------- PROCESS ----------------

func process(workerID int, task Task) {
	fmt.Printf("Worker %d processing task %d\n", workerID, task.ID)
	time.Sleep(50 * time.Millisecond)
}

// ---------------- MAIN ----------------

func main() {
	numWorkers := 4

	global := &GlobalQueue{}
	locals := make([]*LocalQueue, numWorkers)

	for i := 0; i < numWorkers; i++ {
		locals[i] = &LocalQueue{}
	}

	// ---------------- PRODUCER ----------------
	// INTENTIONALLY UNBALANCED → all tasks to worker 0
	for i := 1; i <= 20; i++ {
		locals[0].Push(Task{ID: i})
	}

	// ---------------- WORKERS ----------------
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, locals[i], global, locals, &wg)
	}

	wg.Wait()

	fmt.Println("\n✅ All tasks completed")
}
