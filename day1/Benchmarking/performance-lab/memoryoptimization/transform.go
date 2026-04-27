package memoryoptimization

import (
	"strings"
)

// naive version: repeated allocations
func BuildCSVNaive(values []int) string {
	out := ""
	for _, v := range values {
		out += strconv(v) + ","
	}
	return out
}

// optimized version: fewer allocations
func BuildCSVOptimized(values []int) string {
	var b strings.Builder
	b.Grow(len(values) * 6)

	for _, v := range values {
		b.WriteString(strconv(v))
		b.WriteByte(',')
	}
	return b.String()
}

func strconv(v int) string {
	// minimal local conversion to avoid fmt
	return itoa(v)
}

func itoa(v int) string {
	// simple conversion for demo; in real code use strconv.Itoa
	if v == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	n := v
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
