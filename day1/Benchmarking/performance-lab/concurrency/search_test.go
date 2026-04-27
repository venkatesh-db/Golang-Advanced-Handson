package concurrency

import (
	"context"
	"testing"
)

func makeIDs(n int) []int {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	return ids
}

func BenchmarkProcessUnbounded(b *testing.B) {
	ids := makeIDs(1000)
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		ProcessUnbounded(ctx, ids)
	}
}

func BenchmarkProcessWithPool(b *testing.B) {
	ids := makeIDs(1000)
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		ProcessWithPool(ctx, ids, 10)
	}
}