package main

import (
	_ "embed"
	"log"
)

const (
	// windowSize4 is the size of the window to check for unique characters
	windowSize4 = 4
)

var (
	//go:embed input_day06.txt
	input string
)

// firstMarkerIndex returns the first index of the first marker
func firstMarkerIndex(input string, windowSize int) int {

	for i := 0; i < len(input)-4; i++ {
		seq := input[i : i+windowSize]
		if isSeqUnique(seq) {
			return i + windowSize
		}

	}
	return -1
}

// isSeqUnique returns true if the sequence is unique
func isSeqUnique(seq string) bool {
	seen := make(map[rune]bool)
	for _, letter := range seq {
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}

func main() {

	partA := firstMarkerIndex(input, windowSize4)

	log.Print("The first marker index is: ", partA, " solving part A")
}
