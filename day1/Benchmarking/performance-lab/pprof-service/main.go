/*

1️⃣ Profiling (pprof service)
cd pprof-service
go run main.go

Then in another terminal:

go tool pprof http://localhost:8080/debug/pprof/profile?seconds=10
2️⃣ Memory Optimization (benchmark)
cd memory-optimization
go test -bench=. -benchmem

👉 This compares:

naive vs optimized
allocations and speed
3️⃣ Concurrency (benchmark)
cd concurrency
go test -bench=.

👉 This compares:
unbounded goroutines vs worker pool

*/

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		nStr := r.URL.Query().Get("n")
		if nStr == "" {
			nStr = "50000"
		}
		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			http.Error(w, "invalid n", http.StatusBadRequest)
			return
		}

		result := compute(n)
		w.Write([]byte(result))
	})

	// pprof endpoints are available under /debug/pprof/
	// e.g., /debug/pprof/profile?seconds=10, /debug/pprof/heap
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("service listening on :8080")
	log.Println("pprof available at http://localhost:8080/debug/pprof/")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func compute(n int) string {
	// intentionally allocation-heavy work to make profiles visible
	var builder strings.Builder
	builder.Grow(n * 6)

	for i := 0; i < n; i++ {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteByte(',')
	}
	return builder.String()
}
