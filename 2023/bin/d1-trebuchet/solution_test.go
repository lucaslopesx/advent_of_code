package main

import "testing"

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		input []string
		expected int
	}{
		{
			input: []string{"3218093", "312uhnduij24hn29adawdwadaw", "te7te"},
			expected: 33 + 39 + 77,
		},
		{
			input:    []string{"102", "305", "807"},
			expected: 12 + 35 + 87,
		},
		{
			input:    []string{"abc", "1def9", "12a3"},
			expected: 0 + 19 + 13,
		},
		{
			input:    []string{"", " ", "0"},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := calculateSum(test.input)

		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}