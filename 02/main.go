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
	run(text)
	end := time.Now()
	fmt.Printf("Running time: %v\n", end.Sub(start))
	start = time.Now()
	run2(text)
	end = time.Now()
	fmt.Printf("Running time: %v\n", end.Sub(start))
}

func getDoubles(bounds []string) ([]int, int) {
	doubles := []int{}
	sum := 0
	startDigits := len(bounds[0]) / 2
	endDigits := len(bounds[1])/2 + len(bounds[1])%2
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	startHalf := bounds[0][:startDigits]
	endHalf := bounds[1][:endDigits]
	startHalfInt, _ := strconv.Atoi(startHalf)
	endHalfInt, _ := strconv.Atoi(endHalf)
	for i := startHalfInt; i <= endHalfInt; i++ {
		doubleStr := fmt.Sprintf("%d%d", i, i)
		doubleInt, _ := strconv.Atoi(doubleStr)
		if doubleInt >= start && doubleInt <= end {
			doubles = append(doubles, doubleInt)
			sum += doubleInt
		}
	}
	return doubles, sum
}

func run(input string) string {
	ranges := strings.Split(input, ",")
	total := 0
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		_, sum := getDoubles(bounds)
		total += sum
	}
	fmt.Printf("Total: %d\n", total)
	return fmt.Sprintf("%d", total)
}

func isRepeated(n int) bool {
	digits := strconv.Itoa(n)
	for length := 1; length <= len(digits)/2; length++ {
		// abc..nabc..nabc..n...
		pattern := digits[0:length]
		index := length
		for ; index+length <= len(digits); index += length {
			if digits[index:index+length] != pattern {
				break
			}
		}
		if index == len(digits) {
			return true
		}
	}
	return false
}

func getRepeats(bounds []string) ([]int, int) {
	doubles := []int{}
	sum := 0
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	for i := start; i <= end; i++ {
		if isRepeated(i) {
			doubles = append(doubles, i)
			sum += i
		}
	}
	return doubles, sum
}

func run2(input string) string {
	ranges := strings.Split(input, ",")
	total := 0
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		_, sum := getRepeats(bounds)
		total += sum
	}
	fmt.Printf("Total: %d\n", total)
	return fmt.Sprintf("%d", total)
}
