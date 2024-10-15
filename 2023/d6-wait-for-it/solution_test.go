package main

import "testing"

func TestCalculateWaysOfBeatingRaces(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 288,
		},
	}

	for _, test := range tests {
		result := calculateWaysOfBeatingRaces(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
