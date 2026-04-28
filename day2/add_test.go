package main

import "testing"

// go test -v

// ---------------- CODE UNDER TEST ----------------

func Add(a, b int) int {
	return a + b
}

// ---------------- TEST ----------------

func TestAdd(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero case", 0, 5, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
	}

	for _, tc := range tests {
		tc := tc // important for subtests

		t.Run(tc.name, func(t *testing.T) {

			result := Add(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("expected %d, got %d",
					tc.expected, result)
			}
		})
	}
}
