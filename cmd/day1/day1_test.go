package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElves(t *testing.T) {

	prepareArrayofElves()
	sortElves()

	// assert most cals on top elf is 74394
	assert.True(t, elves[len(elves)-1].TotalCals == 74394, "The elf with the most total calories has: ", elves[len(elves)-1].TotalCals)

	// assert top three are 212836
	topElves := elves[len(elves)-3:]

	cals := 0
	for _, elf := range topElves {
		cals += elf.TotalCals
	}

	assert.True(t, cals == 212836, "The elves with the most total calories have: ", cals)

}
