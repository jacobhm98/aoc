package main

import (
	"aoc-22/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("day3/input.txt")
	runningTotalPriority := 0
	for i := 0; i < len(lines); i += 3 {
		duplicate_char := findDuplicateCharInLines(lines[i : i+3])
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

func findDuplicateCharInLines(lines []string) byte {
	seenOnce := make(map[byte]bool)
	seenTwice := make(map[byte]bool)
	if len(lines) != 3 {
		panic("this shit whack yo")
	}
	for _, char := range []byte(lines[0]) {
		seenOnce[char] = true
	}
	for _, char := range []byte(lines[1]) {
		if seenOnce[char] {
			seenTwice[char] = true
		}
	}
	for _, char := range []byte(lines[2]) {
		if seenTwice[char] {
			return char
		}
	}
	panic("no triple occuring character found")
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
