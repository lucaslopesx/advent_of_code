package main

import "testing"

func TestCalculateReports(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		result := runReports(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
