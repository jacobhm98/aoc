package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileName string) []string {
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
