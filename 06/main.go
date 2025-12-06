package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func getArgs(s string) [][]string {
	var args [][]string
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, line := range lines {
		args = append(args, strings.Fields(line))
	}
	// transpose args
	numArgs := len(args[0])
	transposed := make([][]string, numArgs)
	for i := 0; i < numArgs; i++ {
		for j := 0; j < len(args); j++ {
			transposed[i] = append(transposed[i], args[j][i])
		}
	}
	return transposed
}

func evaluate(row []string) int {
	op := row[len(row)-1]
	total, _ := strconv.Atoi(row[0])
	for i := 1; i < len(row)-1; i++ {
		val, _ := strconv.Atoi(row[i])
		switch op {
		case "+":
			total += val
		case "*":
			total *= val
		}
	}
	return total
}

func run(input string) string {
	args := getArgs(input)
	total := 0
	for _, row := range args {
		total += evaluate(row)
	}
	fmt.Printf("Total: %d\n", total)
	return fmt.Sprintf("%d", total)
}

func getGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	maxCols := 0
	for i, line := range lines {
		grid[i] = []rune(line)
		if len(line) > maxCols {
			maxCols = len(line)
		}
	}
	// transpose
	numRows := len(grid)
	newGrid := make([][]rune, maxCols)
	for i := 0; i < maxCols; i++ {
		newGrid[i] = make([]rune, numRows)
		for j := 0; j < numRows; j++ {
			newGrid[i][j] = grid[j][maxCols-1-i]
		}
	}
	return newGrid
}

func getValueFromRunes(runes []rune) (int, rune) {
	val := 0
	for _, r := range runes {
		if r >= '0' && r <= '9' {
			val = val*10 + int(r-'0')
		}
	}
	return val, runes[len(runes)-1]
}

func evaluateGrid(grid [][]rune) int {
	total := 0
	argStack := []int{}
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		val, op := getValueFromRunes(row)
		argStack = append(argStack, val)
		if op == '+' {
			result := argStack[0]
			for j := 1; j < len(argStack); j++ {
				result += argStack[j]
			}
			argStack = []int{}
			total += result
			i++ // skip the blank row
		} else if op == '*' {
			result := argStack[0]
			for j := 1; j < len(argStack); j++ {
				result *= argStack[j]
			}
			argStack = []int{}
			total += result
			i++ // skip the blank row
		}
	}

	return total
}

func run2(input string) string {
	grid := getGrid(input)
	result := evaluateGrid(grid)
	fmt.Printf("Grid evaluation result: %d\n", result)
	return fmt.Sprintf("%d", result)
}
