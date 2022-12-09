package main

import (
	"aoc-22/pkg/stack"
	"aoc-22/pkg/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := utils.ReadFile("day5/input.txt")
	stacks, moves := createStacksAndGetListOfMoves(input)
	stacks = executeMovesV2(stacks, moves)
	fmt.Println(getFinishingState(stacks))
}

func getFinishingState(stacks []stack.StringStack) string {
	var sb strings.Builder
	for _, stack := range stacks {
		if stack.IsEmpty() {
			continue
		}
		sb.WriteString(stack.Pop())
	}
	return sb.String()
}

type move struct {
	from   int
	to     int
	amount int
}

func executeMovesV2(stacks []stack.StringStack, moves []string) []stack.StringStack {
	for _, line := range moves {
		move := getMove(line)
		stacks[move.to].PushList(stacks[move.from].PopList(move.amount))
	}
	return stacks
}

func executeMoves(stacks []stack.StringStack, moves []string) []stack.StringStack {
	for _, line := range moves {
		move := getMove(line)
		for i := 0; i < move.amount; i++ {
			stacks[move.to].Push(stacks[move.from].Pop())
		}
	}
	return stacks
}

func getMove(line string) move {
	words := strings.Split(line, " ")
	from, _ := strconv.Atoi(words[3])
	to, _ := strconv.Atoi(words[5])
	amount, _ := strconv.Atoi(words[1])
	return move{
		from:   from - 1,
		to:     to - 1,
		amount: amount,
	}
}

func createStacksAndGetListOfMoves(input []string) ([]stack.StringStack, []string) {
	var puzzleLines []string
	var instructions []string
	for i, line := range input {
		if line == "" {
			puzzleLines = input[:i-1]
			instructions = input[i+1:]
			break
		}
	}
	return createStacks(puzzleLines), instructions
}

func createStacks(puzzleLines []string) []stack.StringStack {
	puzzleBoard := make([][]rune, len(puzzleLines))
	for _, line := range puzzleLines {
		puzzleBoard = append(puzzleBoard, []rune(line))
	}

	var stacks []stack.StringStack
	for i, char := range puzzleBoard[len(puzzleBoard)-1] {
		if unicode.IsLetter(char) {
			currStack := stack.New()
			populateStackUpwards(&currStack, i, puzzleBoard)
			stacks = append(stacks, currStack)
		}
	}
	return stacks
}

func populateStackUpwards(stack *stack.StringStack, i int, puzzleBoard [][]rune) {
	for j := len(puzzleBoard) - 1; j >= 0; j-- {
		if i >= len(puzzleBoard[j]) {
			break
		}
		if unicode.IsLetter(puzzleBoard[j][i]) {
			stack.Push(string(puzzleBoard[j][i]))
		}
	}
}
