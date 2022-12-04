package main

// https://adventofcode.com/2022/day/3

import (
	_ "embed"
	"log"
	"strings"
)

const (
	// packDelimiter is the delimiter between packs
	packDelimiter = "\n"
)

var (
	//go:embed input_day03.txt
	input string
)

func main() {
	prepareInput()
	priority := calculatePriorityOne()
	log.Printf("The priority is %d", priority)
	// Part 2
	priority = calculatePriorityTwo()
	log.Printf("For part two, the priority is %d", priority)
}

// prepareInput prepares the input for processing
func prepareInput() {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatal("Input for Day 3 is empty")
	}
}

// calculate the letter value
func letterValue(letter rune) int {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i, l := range alphabet {
		if l == letter {
			return i + 1
		}
	}
	return 0
}

// find the doubled letter in two strings
func findDoubledLetter(inputA string, inputB string) rune {
	// seen is a map of seen characters
	seen := make(map[rune]bool)

	for _, letter := range inputA {
		seen[letter] = true
	}

	for _, letter := range inputB {
		if seen[letter] {
			return letter
		}
	}
	return 0
}

// calculatePriorityOne
func calculatePriorityOne() int {
	priority := 0
	packs := strings.Split(input, packDelimiter)
	for _, pack := range packs {
		pack = strings.TrimSpace(pack)
		mid := len(pack) / 2
		firsHalf := pack[:mid]
		secondHalf := pack[mid:]
		doubledLetter := findDoubledLetter(firsHalf, secondHalf)
		value := letterValue(doubledLetter)
		priority += value
	}

	return priority
}

// calculatePriorityTwo calculates the priority of the item that is common to three packs
func calculatePriorityTwo() int {

	rucksacks := strings.Split(input, packDelimiter)
	priority := 0

	for i := 0; i < len(rucksacks); i += 3 {
		firstRuck := rucksacks[i]
		secondRuck := rucksacks[i+1]
		thirdRuck := rucksacks[i+2]

		commonLetter := findCommonLetterFromThree(firstRuck, secondRuck, thirdRuck)
		priority += letterValue(commonLetter)
	}

	return priority
}

func findCommonLetterFromThree(first string, second string, third string) rune {
	seenOne := make(map[rune]bool)
	for _, letter := range first {
		seenOne[letter] = true
	}

	seenTwo := make(map[rune]bool)
	for _, l2 := range second {
		seenTwo[l2] = true
	}

	for _, l3 := range third {
		if seenOne[l3] && seenTwo[l3] {
			return l3
		}
	}

	return 0
}
