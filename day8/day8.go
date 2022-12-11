package main

import (
	"aoc-22/pkg/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("day8/input.txt")
	treeGrid, isTreeVisible := initializeGrid(lines)
	iterateInwardsFromAllAnglesAndCheckIfVisible(treeGrid, isTreeVisible)
	fmt.Println(countTrues(isTreeVisible))
}

func countTrues(isTreeVisible [][]bool) int {
	count := 0
	for _, i := range isTreeVisible {
		for _, j := range i {
			if j {
				count += 1
			}
		}
	}
	return count
}
func iterateInwardsFromAllAnglesAndCheckIfVisible(treeGrid [][]uint8, isTreeVisible [][]bool) {
	//iterate from left
	for i := 1; i < len(treeGrid)-1; i++ {
		largestTreeSeenSoFar := treeGrid[i][0]
		for j := 1; j < len(treeGrid[i])-1; j++ {
			if treeGrid[i][j] > largestTreeSeenSoFar {
				isTreeVisible[i][j] = true
				largestTreeSeenSoFar = treeGrid[i][j]
			}
			if treeGrid[i][j] == 9 {
				break
			}
		}
	}

	//iterate from right
	for i := 1; i < len(treeGrid); i++ {
		largestTreeSeenSoFar := treeGrid[i][len(treeGrid[0])-1]
		for j := len(treeGrid[i]) - 2; j > 0; j-- {
			if treeGrid[i][j] > largestTreeSeenSoFar {
				isTreeVisible[i][j] = true
				largestTreeSeenSoFar = treeGrid[i][j]
			}
			if treeGrid[i][j] == 9 {
				break
			}
		}
	}
	//iterate from top
	for j := 1; j < len(treeGrid[0])-1; j++ {
		largestTreeSeenSoFar := treeGrid[0][j]
		for i := 1; i < len(treeGrid)-1; i++ {
			if treeGrid[i][j] > largestTreeSeenSoFar {
				isTreeVisible[i][j] = true
				largestTreeSeenSoFar = treeGrid[i][j]
			}
			if treeGrid[i][j] == 9 {
				break
			}
		}
	}

	//iterate from bottom
	for j := 1; j < len(treeGrid[0])-1; j++ {
		largestTreeSeenSoFar := treeGrid[len(treeGrid)-1][j]
		for i := len(treeGrid) - 2; i > 0; i-- {
			if treeGrid[i][j] > largestTreeSeenSoFar {
				isTreeVisible[i][j] = true
				largestTreeSeenSoFar = treeGrid[i][j]
			}
			if treeGrid[i][j] == 9 {
				break
			}
		}
	}
}

func initializeGrid(lines []string) ([][]uint8, [][]bool) {
	treeGrid := make([][]uint8, len(lines))
	isTreeVisible := make([][]bool, len(lines))
	for i, row := range lines {
		treeGrid[i] = []uint8(row)
		isTreeVisible[i] = make([]bool, len(row))
	}
	for i, row := range isTreeVisible {
		for j := range row {
			if i == 0 || i == len(isTreeVisible)-1 || j == 0 || j == len(row)-1 {
				isTreeVisible[i][j] = true
			}
		}
	}
	return treeGrid, isTreeVisible
}
