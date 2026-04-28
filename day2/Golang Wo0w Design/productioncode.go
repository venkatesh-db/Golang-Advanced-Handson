package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Panic counter metric for Prometheus
var panicCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "panic_total",
		Help: "Total number of panics recovered",
	},
)

func init() {
	// Register Prometheus metric
	prometheus.MustRegister(panicCount)
}

// Middleware for panic recovery
func withRecovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				panicCount.Inc() // Prometheus counter
				log.Printf("🔥 Panic Recovered: %v\n%s", rec, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

// Handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "✅ Home — healthy and happy!")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("💥 Boom! Something broke.")
}

func goroutineHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				panicCount.Inc()
				log.Printf("⚠️ Panic in Goroutine: %v\n%s", rec, debug.Stack())
			}
		}()
		time.Sleep(500 * time.Millisecond)
		panic("💣 Goroutine exploded!")
	}()
	fmt.Fprintln(w, "🌀 Started goroutine that will panic in 0.5s.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", withRecovery(homeHandler))
	mux.HandleFunc("/panic", withRecovery(panicHandler))
	mux.HandleFunc("/goroutine", withRecovery(goroutineHandler))
	mux.Handle("/metrics", promhttp.Handler()) // Prometheus endpoint

	log.Println("🚀 Server started at http://localhost:7080")
	log.Fatal(http.ListenAndServe(":7080", mux))
}
