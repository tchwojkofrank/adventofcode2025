package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readInput(fname string) string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	return string(content)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Enter input file name.\n")
		return
	}
	params := os.Args[1]
	inputName := strings.Split(params, " ")[0]
	text := readInput(inputName)
	start := time.Now()
	result := run(text)
	end := time.Now()
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Running time: %v\n", end.Sub(start))
	start = time.Now()
	result = run2(text)
	end = time.Now()
	fmt.Printf("Result 2: %v\n", result)
	fmt.Printf("Running time: %v\n", end.Sub(start))
}

func split(lines []string, n int) (string, int) {
	splitCount := 0
	input := []byte(lines[n])
	output := []byte(lines[n+1])
	if n == 0 {
		for i := 0; i < len(input); i++ {
			switch input[i] {
			case 'S':
				output[i] = '|'
			}
		}
		return string(output), splitCount
	}
	prev := []byte(lines[n-1])
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '^':
			switch prev[i] {
			case '|':
				if i > 0 {
					output[i-1] = '|'
				}
				if i < len(input)-1 {
					output[i+1] = '|'
				}
				splitCount++
			}
		default:
			switch prev[i] {
			case '|':
				output[i] = '|'
			}
		}
	}
	return string(output), splitCount
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}

func run(input string) string {
	lines := strings.Split(input, "\n")
	totalSplitterCount := 0
	splitCount := 0
	// printLines(lines)
	// fmt.Printf("Splits: %d\n\n", totalSplitterCount)
	for i := 0; i < len(lines)-1; i += 2 {
		lines[i+1], splitCount = split(lines, i)
		totalSplitterCount += splitCount
		// printLines(lines)
		// fmt.Printf("Splits: %d\n\n", totalSplitterCount)
	}
	// fmt.Printf("Splitter count: %d\n", totalSplitterCount)
	return fmt.Sprintf("%d", totalSplitterCount)
}

type Position struct {
	line int
	pos  int
}

type Cache map[Position]int

func travel(lines []string, lineIndex int, position int, cache Cache) int {
	if val, ok := cache[Position{lineIndex, position}]; ok {
		return val
	}
	if lineIndex >= len(lines) {
		cache[Position{lineIndex, position}] = 1
		return 1
	}
	if position < 0 || position >= len(lines[0]) {
		cache[Position{lineIndex, position}] = 0
		return 0
	}
	pathCount := 0

	switch lines[lineIndex][position] {
	case '^':
		pathCount += travel(lines, lineIndex+2, position-1, cache)
		pathCount += travel(lines, lineIndex+2, position+1, cache)
	default:
		pathCount += travel(lines, lineIndex+1, position, cache)
	}
	cache[Position{lineIndex, position}] = pathCount
	return pathCount
}

func run2(input string) string {
	lines := strings.Split(input, "\n")
	pos := strings.Index(lines[0], "S")
	cache := make(Cache)
	pathCount := travel(lines, 2, pos, cache)
	// fmt.Printf("Travel count: %d\n", pathCount)
	return fmt.Sprintf("%d", pathCount)
}
