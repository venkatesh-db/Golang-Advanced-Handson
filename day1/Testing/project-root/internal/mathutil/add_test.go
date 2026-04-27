package mathutil

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negative", -2, -3, -5},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
