package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundScoreB(t *testing.T) {

	tests := []struct {
		description string
		opponent    string
		you         string
		expected    int
	}{
		{
			description: "Rock vs You must draw",
			opponent:    opponentRock,
			you:         youMustDraw,
			expected:    4,
		},
		{
			description: "Paper vs You must lose",
			opponent:    opponentPaper,
			you:         youMustLose,
			expected:    1,
		},
		{
			description: "Scissors vs You must win",
			opponent:    opponentScissors,
			you:         youMustWin,
			expected:    7,
		},
		{description: "Rock vs You must win",
			opponent: opponentRock,
			you:      youMustWin,
			expected: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, roundScoreB(tc.opponent, tc.you))
		})
	}

}
