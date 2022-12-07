package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputLines := readFile("day2/input.txt")
	playsToPoints := createPlayCombinationToPointsMap()
	points := calculateTotalPoints(inputLines, playsToPoints)
	fmt.Println(points)
}

func calculateTotalPoints(lines []string, playCombinationsToPoints map[string]map[string]int) int {
	runningPoints := 0
	for _, line := range lines {
		p1Move, p2Move := getMovesFromLine(line)
		runningPoints += playCombinationsToPoints[p1Move][p2Move]
	}
	return runningPoints
}

func getMovesFromLine(line string) (string, string) {
	moves := strings.Split(line, " ")
	if len(moves) != 2 {
		panic("more than two types of moves in a line")
	}
	return moves[0], moves[1]
}

//rock -- p1 A p2 x
//paper -- p1 B p2 y
//scissor -- p1 C p2 Z
func createPlayCombinationToPointsMap() map[string]map[string]int {
	return map[string]map[string]int{"A": {"X": 4, "Y": 8, "Z": 3}, "B": {"X": 1, "Y": 5, "Z": 9}, "C": {"X": 7, "Y": 2, "Z": 6}}
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines
}
