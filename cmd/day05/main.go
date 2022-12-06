package main

// https://adventofcode.com/2022/day/4

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gammazero/deque"
)

const (
	// stackNumber is the number of stacks
	stackNumber = 9
)

var (
	//go:embed crates.txt
	cratesInput string

	//go:embed moves.txt
	movesInput string
)

type Stack struct {
	crates deque.Deque[Crate]
}

type Crate struct {
	name string
}

func (stack *Stack) push(crate Crate) {
	stack.crates.PushBack(crate)
}

func (stack *Stack) peek() string {
	if stack.crates.Len() == 0 {
		return ""
	}

	top := stack.crates.Back()

	return top.name
}

func (stack *Stack) pop() (Crate, error) {
	if stack.crates.Len() == 0 {
		return Crate{}, errors.New("Stack is empty")
	}

	top := stack.crates.PopBack()

	return top, nil
}

// prepend is the "addFirst" or PushFront method
func (stack *Stack) prepend(crate Crate) {
	stack.crates.PushFront(crate)
}

func main() {
	stacks := parseCrates()

	stacks = makeMoves(stacks)

	peekAll(stacks)

}

// peekAll the stacks, logging the value
func peekAll(stacks []Stack) {
	message := ""
	for _, stack := range stacks {
		message += stack.peek()
	}

	log.Print(message)
}

func parseCrates() []Stack {

	// stacks is a slice of crates
	stacks := make([]Stack, stackNumber)

	if len(cratesInput) == 0 {
		log.Fatal("Crates for Day 4 is empty")
	}

	crates := strings.Split(cratesInput, "\n")

	for _, crateRow := range crates {
		log.Print(crateRow)
		for i := 1; i < len(crateRow); i += 4 {
			crateName := string(crateRow[i])

			if crateName == " " {
				continue
			} else {
				crate := Crate{
					name: crateName,
				}
				// which stack is it?
				stackNumber := (i - 1) / 4
				stacks[stackNumber].prepend(crate)
			}
		}
	}

	return stacks
}

// make moves takes a stack state and makes a series of moves
func makeMoves(stacks []Stack) []Stack {
	if len(movesInput) == 0 {
		log.Fatal("Moves for Day 5 is empty")
	}

	moves := strings.Split(movesInput, "\n")

	for i, move := range moves {
		log.Print("Move ", i, " is ", move)
		var toMove, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &toMove, &from, &to)
		stacks = moveCrate(stacks, toMove, from, to)
	}

	return stacks
}

// moveCrate moves a crate from one stack to another
func moveCrate(stacks []Stack, toMove, from, to int) []Stack {

	log.Print("Moving to ", to)
	// Convert to 0-based indexing
	to = to - 1
	from = from - 1

	peekAll(stacks)

	// pop the crate from the stack
	crate, err := stacks[from].pop()
	if err != nil {
		// in this case, no op
		log.Print("No crate to move")
		return stacks
	}

	// push the crate to the stack
	stacks[to].push(crate)
	return stacks
}
