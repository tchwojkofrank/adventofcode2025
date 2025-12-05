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

func getMap(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func isAccessible(r int, c int, grid [][]rune) bool {
	if grid[r][c] != '@' {
		return false
	}
	neighbors := 0
	for dr := -1; dr <= 1; dr++ {
		if r+dr < 0 || r+dr >= len(grid) {
			continue
		}
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if c+dc < 0 || c+dc >= len(grid[0]) {
				continue
			}
			if grid[r+dr][c+dc] == '@' {
				neighbors++
			}
		}
	}
	return neighbors < 4
}

func run(input string) string {
	count := 0
	grid := getMap(input)
	for r, row := range grid {
		for c := range row {
			if isAccessible(r, c, grid) {
				count++
			}
		}
	}
	return fmt.Sprintf("%d", count)
}

func removeAccessible(grid [][]rune) ([][]rune, int) {
	removed := 0
	newGrid := make([][]rune, len(grid))
	for r := range grid {
		newGrid[r] = make([]rune, len(grid[0]))
	}
	for r, row := range grid {
		for c, val := range row {
			if val == '@' && isAccessible(r, c, grid) {
				newGrid[r][c] = '.'
				removed++
			} else {
				newGrid[r][c] = val
			}
		}
	}
	return newGrid, removed
}

func run2(input string) string {
	count := 0
	removed := -1
	grid := getMap(input)
	for removed != 0 {
		grid, removed = removeAccessible(grid)
		count += removed
	}
	return fmt.Sprintf("%d", count)
}
