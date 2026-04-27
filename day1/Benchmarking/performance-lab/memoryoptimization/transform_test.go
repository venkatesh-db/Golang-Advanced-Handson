package memoryoptimization

import (
	"math/rand"
	"testing"
)

func makeData(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Intn(100000)
	}
	return data
}

func BenchmarkBuildCSVNaive(b *testing.B) {
	data := makeData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = BuildCSVNaive(data)
	}
}

func BenchmarkBuildCSVOptimized(b *testing.B) {
	data := makeData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = BuildCSVOptimized(data)
	}
}
