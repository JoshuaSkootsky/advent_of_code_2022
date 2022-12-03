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
	priority := calculatePriority()
	log.Printf("The priority is %d", priority)
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
func findDoubledLetter(firstHalf string, secondHalf string) rune {
	for _, letter := range firstHalf {
		for _, letter2 := range secondHalf {
			if letter == letter2 {
				return letter
			}
		}
	}
	return 0
}

// calculatePriority
func calculatePriority() int {
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
