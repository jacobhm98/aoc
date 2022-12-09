package main

import (
	"aoc-22/pkg/utils"
	"fmt"
)

func main() {
	line := utils.ReadFile("day6/input.txt")[0]
	byteLine := []byte(line)
	var window []byte
	for i := 0; i < len(byteLine)-13; i++ {
		window = byteLine[i : i+14]
		if !containsDuplicates(window){
			fmt.Println(i + 14)
			break
		}
	}
}

func  containsDuplicates(window []byte) bool {
	chars := make(map[byte]bool)
	for _, byte := range window{
		_, contains := chars[byte]
		if contains {return true}
		chars[byte] = true
	}
	return false
}
