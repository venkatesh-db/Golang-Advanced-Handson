package main

import "tests/internal/logging"

/*


Run all tests
go test ./...

Run specific package
go test ./internal/mathutil

Run fuzz test
go test -fuzz=FuzzReverse ./internal/mathutil

*/

func main() {
	logger := logging.NewLogger()

	logger.Info("application started",
		"service", "basic-example",
	)
}
