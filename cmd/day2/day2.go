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

	// youMustLose is your code for you must lose
	youMustLose = "X"

	// youMustDraw is your code for you must draw
	youMustDraw = "Y"

	// youMustWin is your code for you must win
	youMustWin = "Z"
)

var (
	//go:embed input_day02.txt
	input string
)

func main() {
	prepareInput()
	totalA := scoreInput(roundScoreA)
	log.Print("The total for Part A is: ", totalA)
	totalB := scoreInput(roundScoreB)
	log.Print("The total for Part B is: ", totalB)
}

func prepareInput() {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatal("Input for Day 2 is empty")
	}
}

func scoreInput(roundScore func(opponent string, you string) int) int {
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

// roundScoreA returns the score for a round in Part A
func roundScoreA(opponent string, you string) int {

	// yourChoiceScore looks up the score for your choice
	yourChoiceScore := map[string]int{yourRock: 1, yourPaper: 2, yourScissors: 3}

	// outcomeMap looks up the score for your round versus your opponent
	// 0 for loss, 3 for draw, 6 for win

	outcomeMap := map[string]map[string]int{
		opponentRock: {
			yourRock:     3,
			yourPaper:    6,
			yourScissors: 0,
		},
		opponentPaper: {
			yourRock:     0,
			yourPaper:    3,
			yourScissors: 6,
		},
		opponentScissors: {
			yourRock:     6,
			yourPaper:    0,
			yourScissors: 3,
		},
	}

	return outcomeMap[opponent][you] + yourChoiceScore[you]
}

func roundScoreB(opponent string, you string) int {

	// yourChoiceScore looks up the score for your choice
	yourChoiceScore := map[string]int{yourRock: 1, yourPaper: 2, yourScissors: 3}

	// now X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
	// depending on what result you need, you will need to choose a different move.

	winScore := map[string]int{youMustLose: 0, youMustDraw: 3, youMustWin: 6}

	// choiceMap tells you what move to make
	choiceMap := map[string]map[string]string{
		opponentRock: {
			youMustLose: yourScissors,
			youMustDraw: yourRock,
			youMustWin:  yourPaper,
		},
		opponentPaper: {
			youMustLose: yourRock,
			youMustDraw: yourPaper,
			youMustWin:  yourScissors,
		},
		opponentScissors: {
			youMustLose: yourPaper,
			youMustDraw: yourScissors,
			youMustWin:  yourRock,
		},
	}

	yourChoice := choiceMap[opponent][you]

	return yourChoiceScore[yourChoice] + winScore[you]
}
