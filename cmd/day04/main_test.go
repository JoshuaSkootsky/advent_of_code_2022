package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlap(t *testing.T) {

	tests := []struct {
		description string
		elves       string
		expected    bool
	}{
		{
			description: "Elves overlap but not completely",
			elves:       "1-3,2-4",
			expected:    false,
		},
		{
			description: "Elves completely overlap",
			elves:       "1-3,1-3",
			expected:    true,
		},
		{
			description: "Elves completely overlap",
			elves:       "1-4,2-3",
			expected:    true,
		},
		{
			description: "Elves do not overlap",
			elves:       "1-2,3-4",
			expected:    false,
		},
		{
			description: "Elves do overlap",
			elves:       "20-82,19-87",
			expected:    true,
		},
		{
			description: "Elves do overlap",
			elves:       "98-98,40-98",
			expected:    true,
		},
		{
			description: "Elves do overlap",
			elves:       "40-40,40-98",
			expected:    true,
		},
		{
			description: "Elves do not overlap",
			elves:       "7-89,6-8",
			expected:    false,
		},
		{
			description: "Elves do not overlap",
			elves:       "4-6,6-7",
			expected:    false,
		},
		{
			description: "Elves do overlap",
			elves:       "4-6,6-6",
			expected:    true,
		},
		{
			description: "Elves do not overlap",
			elves:       "2-3,4-29",
			expected:    false,
		},
		{
			description: "Elves do not overlap",
			elves:       "10-42,9-11",
			expected:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, doesPairOverlap(tc.elves))
		})
	}

}

func TestPartialOverlap(t *testing.T) {
	tests := []struct {
		description string
		elves       string
		expected    bool
	}{
		{
			description: "Elves overlap but not completely",
			elves:       "1-3,2-4",
			expected:    true,
		},
		{
			description: "Elves do not overlap",
			elves:       "2-4,6-8",
			expected:    false,
		},
		{
			description: "Elves do  overlap",
			elves:       "2-8,6-7",
			expected:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, doesPairPartial(tc.elves))
		})
	}
}
