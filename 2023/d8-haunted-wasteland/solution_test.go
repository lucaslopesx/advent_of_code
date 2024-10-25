package main

import "testing"

func TestCalculateCamelCardsGameWinnings(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 6,
		},
	}

	for _, test := range tests {
		result := runInstructions(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
