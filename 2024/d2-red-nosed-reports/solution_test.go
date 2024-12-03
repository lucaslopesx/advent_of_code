package main

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected: 4,
		},
	}

	for _, test := range tests {
		result := calc(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
