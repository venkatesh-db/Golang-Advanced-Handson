

func safeHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})

	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("⏱ Done sleeping")
		case <-done:
			fmt.Println("✅ Clean exit")
		}
	}()

	// Respond to HTTP
	fmt.Fprintln(w, "✅ Goroutine will exit after timeout or cancel")
	close(done) // clean exit signal
}
