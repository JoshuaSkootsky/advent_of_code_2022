package main

// https://adventofcode.com/2022/day/4

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

const (
	// pairDelimiter is the pairDelimiter
	pairDelimiter = "\n"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		log.Fatal("Input for Day 4 is empty")
	}

	count := 0

	for _, pair := range strings.Split(input, pairDelimiter) {
		pairOverlaps := doesPairOverlap(pair)
		if pairOverlaps {
			count++
		}
	}

	log.Print("The number of pairs that overlap is: ", count, " solving part A")
}

// doesPairOverlap returns true if the pair overlaps
func doesPairOverlap(pair string) bool {

	elves := strings.Split(pair, ",")
	elfA := strings.Split(elves[0], "-")
	elfB := strings.Split(elves[1], "-")

	// very important to convert to int - comparisons of the string values of numbers will not work
	s1, _ := strconv.Atoi(elfA[0])
	e1, _ := strconv.Atoi(elfA[1])
	s2, _ := strconv.Atoi(elfB[0])
	e2, _ := strconv.Atoi(elfB[1])

	if s1 >= s2 && e1 <= e2 {
		return true
	}

	if s1 <= s2 && e1 >= e2 {
		return true
	}

	return false
}
