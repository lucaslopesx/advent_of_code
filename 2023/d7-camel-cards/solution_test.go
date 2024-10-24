package main

import "testing"

func TestCalculateCamelCardsGameWinnings(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"JKA32 28",
				"32AJJ 765",
				"3JJA2 684",
				"3333J 220",
				"JJAAA 483",
			},
			expected: 5905,
		},
	}

	for _, test := range tests {
		result := calculateCamelCardsGameWinnings(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
