package main

import (
	"aoc-22/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sectionPairs := utils.ReadFile("day4/input.txt")
	fullyContainedSectionsWithinPairCount := 0
	for _, sectionPair := range sectionPairs {
		if existsOverlapWithinPair(sectionPair) {
			fullyContainedSectionsWithinPairCount += 1
		}
	}
	fmt.Println(fullyContainedSectionsWithinPairCount)
}

type intRange struct {
	lower      int
	upper      int
	difference int
}

func existsOverlapWithinPair(pair string) bool {
	ranges := strings.Split(pair, ",")
	if len(ranges) != 2 {
		panic("more than two ranges detected in a pair")
	}
	firstRange := getIntRange(ranges[0])
	secondRange := getIntRange(ranges[1])
	return (secondRange.lower <= firstRange.lower && secondRange.upper >= firstRange.lower) || (firstRange.lower <= secondRange.lower && firstRange.upper >= secondRange.lower)
}

func existsFullyContainedSectionWithinPair(pair string) bool {
	ranges := strings.Split(pair, ",")
	if len(ranges) != 2 {
		panic("more than two ranges detected in a pair")
	}
	firstRange := getIntRange(ranges[0])
	secondRange := getIntRange(ranges[1])
	if secondRange.difference > firstRange.difference {
		return secondRange.lower <= firstRange.lower && secondRange.upper >= firstRange.upper
	}
	return firstRange.lower <= secondRange.lower && firstRange.upper >= secondRange.upper
}

func getIntRange(stringRange string) intRange {
	valuePair := strings.Split(stringRange, "-")
	if len(valuePair) != 2 {
		panic("more than two ranges detected in a pair")
	}
	lower, lErr := strconv.Atoi(valuePair[0])
	upper, rErr := strconv.Atoi(valuePair[1])
	if lErr != nil || rErr != nil {
		panic("issue getting integer bounds of ranges")
	}
	return intRange{lower, upper, upper - lower}
}
