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

type Range struct {
	start int
	end   int
}

func getRange(s string) Range {
	var r Range
	fmt.Sscanf(s, "%d-%d", &r.start, &r.end)
	return r
}

func getRanges(lines []string) []Range {
	ranges := make([]Range, 0)
	for _, line := range lines {
		ranges = append(ranges, getRange(line))
	}
	return ranges
}

func getIngredient(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getIngredients(lines []string) []int {
	ingredients := make([]int, 0)
	for _, line := range lines {
		ingredients = append(ingredients, getIngredient(line))
	}
	return ingredients
}

func isFresh(ranges []Range, ingredient int) bool {
	for _, r := range ranges {
		if ingredient >= r.start && ingredient <= r.end {
			return true
		}
	}
	return false
}

func run(input string) string {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	lines := strings.Split(parts[0], "\n")
	ranges := getRanges(lines)
	ingredients := getIngredients(strings.Split(parts[1], "\n"))
	count := 0
	for _, ingredient := range ingredients {
		if isFresh(ranges, ingredient) {
			count++
		}
	}
	fmt.Printf("Fresh count: %d\n", count)
	return strconv.Itoa(count)
}

func run2(input string) string {
	return ""
}
