package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	// elfDelim is the delimiter between elves, each seperated by two new lines
	elfDelim = "\n\n"

	// calorieDelim is the delimiter between the calories packs on each elf
	calorieDelim = "\n"
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
	printResultB()
}

func printResult() {
	// Assume the elves are sorted by total calories
	log.Print("The elf with the most total calories has: ", elves[len(elves)-1].TotalCals)
}

// print the three elves with the most calories
func printResultB() {
	// Assume the elves are sorted by total calories, get the last three
	topElves := elves[len(elves)-3:]

	cals := 0
	for _, elf := range topElves {
		cals += elf.TotalCals
	}

	log.Printf("The three elves with the most total calories have: %d", cals)
}

// sortElves in ascending order by total calories
func sortElves() {
	// sort the elves by total calories
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCals < elves[j].TotalCals
	})
}

// prepareArrayofElves prepares the elves array, reading from input
func prepareArrayofElves() {
	// Prepare the input data
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatal("Input file for Day 1 is empty")
	}

	// parse the input
	for _, elfData := range strings.Split(input, elfDelim) {
		var elf Elf
		for _, cals := range strings.Split(elfData, calorieDelim) {
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
