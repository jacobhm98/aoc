package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	elfTotalCalories := readFile("day1/input.txt")
	sort.Slice(elfTotalCalories, func(i, j int) bool { return elfTotalCalories[i] < elfTotalCalories[j] })
	topThreeElves := elfTotalCalories[len(elfTotalCalories)-3:]
	topThreeTotal := 0
	for _, elf := range topThreeElves {
		topThreeTotal += elf
	}
	fmt.Println(topThreeTotal)

}

func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var calories []int
	nextElfCurrSum := 0
	for scanner.Scan() {
		currLine := scanner.Text()
		if currLine == "" {
			calories = append(calories, nextElfCurrSum)
			nextElfCurrSum = 0
			continue
		}
		currItem, _ := strconv.Atoi(currLine)
		nextElfCurrSum += currItem

	}
	return calories
}
