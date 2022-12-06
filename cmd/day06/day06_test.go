package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Never go full TDD
// TestFirstMarkerIndex tests the firstMarkerIndex function
func TestFirstMarkerIndex(t *testing.T) {

	tests := []struct {
		description string
		input       string
		expected    int
		actual      int
	}{
		{
			description: "First marker is at index 5",
			input:       "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:    5,
		},
		{
			description: "First marker is at index 6",
			input:       "nppdvjthqldpwncqszvftbrmjlhg",
			expected:    6,
		},
		{
			description: "First marker is at index 10",
			input:       "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:    10,
		},
		{
			description: "First marker is at index 11",
			input:       "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:    11,
		},
		{
			description: "First marker is at index 7",
			input:       "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected:    7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, firstMarkerIndex(tc.input, windowSize4))
		})
	}

}
