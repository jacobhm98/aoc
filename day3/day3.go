package main

import (
	"aoc-22/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("day3/input.txt")
	runningTotalPriority := 0
	for _, line := range lines {
		duplicate_char := findDuplicateCharInLine(line)
		val := mapAsciiToInt(duplicate_char)
		runningTotalPriority += val
	}
	fmt.Println(runningTotalPriority)
}

func mapAsciiToInt(char byte) int {
	if 65 <= char && char < 91 {
		return int(char - 38)
	}
	if 97 <= char && char < 123 {
		return int(char - 96)
	}
	fmt.Println(char)
	panic("non alphabet byte encountered")
}

func findDuplicateCharInLine(line string) byte {
	firstCompartment := line[:len(line)/2]
	secondCompartment := line[len(line)/2:]
	seenCharacters := make(map[byte]bool)
	for _, char := range []byte(firstCompartment) {
		seenCharacters[char] = true
	}
	for _, char := range []byte(secondCompartment) {
		if seenCharacters[char] {
			return char
		}
	}
	panic(fmt.Sprintf("no duplicate chars detected in %s", line))
}
