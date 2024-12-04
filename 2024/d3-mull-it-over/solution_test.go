package main

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: 48,
		},
	}

	for _, test := range tests {
		result := calc(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
