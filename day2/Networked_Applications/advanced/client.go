
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
	if err != nil {
		panic(err)
	}

	// Add auth header
	req.Header.Set("Authorization", "Bearer production-token")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", resp.Status)
	fmt.Println("response:", string(body))
}