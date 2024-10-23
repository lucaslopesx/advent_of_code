package main

import "testing"

func TestCalculateCamelCardsGameWinnings(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 6440,
		},
	}

	for _, test := range tests {
		result := calculateCamelCardsGameWinnings(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
