package main

import (
	"bytes"
	"testing"
)

/*

go test -bench=. -benchmem

 Meaning (industry interpretation)
80180 iterations
→ Go auto-tuned iterations for stable measurement
13781 ns/op (~13.7 µs)
→ Each request takes ~13 microseconds
✔️ Fast (CPU not your problem)
65599 B/op (~64 KB per call)
→  This is the red flag
10 allocs/op
→ Moderate allocations (can be optimized)

*/

/*

go test -bench=. -benchmem -memprofile=mem.out

go tool pprof mem.out

top

   flat  flat%   sum%        cum   cum%
 6293.19MB 99.67% 99.67%  6293.19MB 99.67%  bytes.growSlice
   12.90MB   0.2% 99.87%  6311.09MB   100%  go-bench.(*Processor).Process
       5MB 0.079%   100%  6298.19MB 99.75%  bytes.(*Buffer).grow
         0     0%   100%  6298.19MB 99.75%  bytes.(*Buffer).WriteString
         0     0%   100%  6311.09MB   100%  go-bench.BenchmarkProcessor
         0     0%   100%  6311.09MB   100%  testing.(*B).launch
         0     0%   100%  6311.60MB   100%  testing.(*B).runN


6293.19MB (99.67%) → bytes.growSlice
12.90MB           → Processor.Process
Translation (very important)

👉 6.2 GB memory allocated
👉 Almost 100% from bytes.growSlice


*/

// ====== PRODUCTION STYLE STRUCT ======
type Processor struct {
	cache [][]byte
}

func NewProcessor() *Processor {
	return &Processor{
		cache: make([][]byte, 0),
	}
}

// INTENTIONAL MEMORY LEAK
func (p *Processor) Process(data string) []byte {
	buf := bytes.NewBuffer(nil)

	for i := 0; i < 1000; i++ {
		buf.WriteString(data)

		/*

			bytes.Buffer keeps growing
			Each growth triggers:
			→ bytes.growSlice (huge allocations)
			Happens 1000 times per request
			Multiplied by benchmark iterations → 💥 6GB
		*/
	}

	result := buf.Bytes()

	//  Leak: keeps growing forever
	p.cache = append(p.cache, result)

	return result
}

// ====== BENCHMARK ======
func BenchmarkProcessor(b *testing.B) {
	processor := NewProcessor()
	data := "golang-performance-test"

	b.ReportAllocs() // 👈 important (shows memory usage)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		processor.Process(data)
	}
}

/* Good solution

func (p *Processor) Process(data string) []byte {
	var buf bytes.Buffer

	// 🔥 Preallocate exact capacity
	buf.Grow(len(data) * 1000)

	for i := 0; i < 1000; i++ {
		buf.WriteString(data)
	}

	result := buf.Bytes()

	// optional cache control
	if len(p.cache) < 1000 {
		p.cache = append(p.cache, result)
	}

	return result
}
*/
