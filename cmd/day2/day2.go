package main

import (
	_ "embed"
	"log"
	"strings"
)

const (
	// roundDelim is the delimiter between rounds
	roundDelim = "\n"

	//choiceDelime is the delimiter between choices
	choiceDelim = " "

	// yourRock is your code for your rock
	yourRock = "X"

	// yourPaper is your code for your paper
	yourPaper = "Y"

	// yourScissors is your code for your scissors
	yourScissors = "Z"

	// opponentRock is your code for opponent's rock
	opponentRock = "A"

	// opponentPaper is your code for opponent's paper
	opponentPaper = "B"

	// opponentScissors is your code for opponent's scissors
	opponentScissors = "C"
)

var (
	//go:embed input_day02.txt
	input string
)

// TODO come up with some structs and functions

func main() {
	total := prepareInput()
	log.Print("The total is: ", total)
}

func prepareInput() int {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatal("Input for Day 2 is empty")
	}

	// rounds are each round of the game
	rounds := strings.Split(input, roundDelim)
	totalScore := 0
	for _, round := range rounds {
		choices := strings.Split(round, choiceDelim)
		opponent := choices[0]
		you := choices[1]
		score := roundScore(opponent, you)
		totalScore += score

	}

	return totalScore
}

// roundSCore returns the score for a round
func roundScore(opponent string, you string) int {

	// yourChoiceScore looks up the score for your choice
	yourChoiceScore := map[string]int{yourRock: 1, yourPaper: 2, yourScissors: 3}

	// outcomeMap looks up the score for your round versus your opponent
	// 0 for loss, 3 for draw, 6 for win

	outcomeMap := map[string]map[string]int{
		opponentRock: map[string]int{
			yourRock:     3,
			yourPaper:    6,
			yourScissors: 0,
		},
		opponentPaper: map[string]int{
			yourRock:     0,
			yourPaper:    3,
			yourScissors: 6,
		},
		opponentScissors: map[string]int{
			yourRock:     6,
			yourPaper:    0,
			yourScissors: 3,
		},
	}

	return outcomeMap[opponent][you] + yourChoiceScore[you]
}
