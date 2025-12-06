package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"chwojkofrank.com/interval"
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

func getRange(s string) interval.Interval {
	var r interval.Interval
	fmt.Sscanf(s, "%d-%d", &r.Start, &r.End)
	return r
}

func addRange(ranges []interval.Interval, newRange interval.Interval) []interval.Interval {
	// take the intersection of newRange with existing ranges
	for _, r := range ranges {
		if newRange.Intersects(r) {
			minus := newRange.Minus(r)
			// add the non-intersecting parts back to ranges
			for _, xr := range minus {
				ranges = addRange(ranges, xr)
			}
			return ranges
		}
	}
	ranges = append(ranges, newRange)

	return ranges
}

func getRanges(lines []string) []interval.Interval {
	ranges := make([]interval.Interval, 0)
	for _, line := range lines {
		ranges = append(ranges, getRange(line))
	}
	// sort ranges by Start
	// simple insertion sort since the number of ranges is small
	for i := 1; i < len(ranges); i++ {
		j := i
		for j > 0 && ranges[j-1].Start > ranges[j].Start {
			ranges[j-1], ranges[j] = ranges[j], ranges[j-1]
			j--
		}
	}
	return ranges
}

func getRanges2(lines []string) []interval.Interval {
	ranges := make([]interval.Interval, 0)
	for _, line := range lines {
		r := getRange(line)
		ranges = addRange(ranges, r)
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

func isFresh(ranges []interval.Interval, ingredient int) bool {
	for _, r := range ranges {
		if ingredient >= r.Start && ingredient <= r.End {
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
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	lines := strings.Split(parts[0], "\n")
	ranges := getRanges2(lines)
	// count how many integers are in the ranges
	count := 0
	for _, r := range ranges {
		count += r.End - r.Start + 1
	}
	fmt.Printf("Total fresh ingredients: %d\n", count)
	return strconv.Itoa(count)
}
