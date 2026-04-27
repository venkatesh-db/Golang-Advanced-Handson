package mathutil

import "testing"

func Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		reversed := Reverse(input)
		double := Reverse(reversed)

		if input != double {
			t.Fatalf("mismatch: %s vs %s", input, double)
		}
	})
}
