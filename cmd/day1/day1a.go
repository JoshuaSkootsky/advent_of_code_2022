package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	// ELF_DELIM is the delimiter between elves, each seperated by two new lines
	ELF_DELIM = "\n\n"

	// CAL_DELIM is the delimiter between the calories packs on each elf
	CALORIE_DELIM = "\n"
)

var (
	//go:embed input_day1.txt
	input string

	// elves is an array of elves
	elves []Elf
)

type Elf struct {
	Calories  []int
	TotalCals int
}

func (e *Elf) calculateTotalCals() int {
	// allow for this to be recalculated
	total := 0
	for _, c := range e.Calories {
		total += c
	}
	e.TotalCals = total
	return total
}

func main() {
	prepareArrayofElves()
	sortElves()
	printResult()
}

func printResult() {
	// Assume the elves are sorted by total calories
	log.Print("The elf with the most total calories has: ", elves[len(elves)-1].TotalCals)
}

func sortElves() {
	// sort the elves by total calories
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCals < elves[j].TotalCals
	})
}

func prepareArrayofElves() {
	// Prepare the input data
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatal("Input file for Day 1 A is empty")
	}

	// parse the input
	for _, elfData := range strings.Split(input, ELF_DELIM) {
		var elf Elf
		for _, cals := range strings.Split(elfData, CALORIE_DELIM) {
			calsInt, err := strconv.Atoi(cals)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			elf.Calories = append(elf.Calories, calsInt)
			elf.calculateTotalCals()
			elves = append(elves, elf)
		}
	}
}
