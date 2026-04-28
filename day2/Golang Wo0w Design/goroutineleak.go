package main

import (
	"fmt"
	"log"
	"net/http"
)

// ⚠️ This simulates a goroutine leak on each request
func leakyHandler(w http.ResponseWriter, r *http.Request) {

	// Spawning a goroutine per request without cleanup
	
	go func() {
		select {} // Block forever — leak!
	}()
	
	fmt.Fprintln(w, "💧 Goroutine leaked! (intentionally)")
}

/*
🧠 What Happens?
Every request to /leak creates a goroutine that:

Calls select {} (blocks forever)

Never dies

Keeps memory, stack, and scheduler time

After 1000 requests, you have 1000 leaking goroutines.

*/

func main() {
	http.HandleFunc("/leak", leakyHandler)

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
