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
	answer := run(text)
	end := time.Now()
	fmt.Printf("Answer 1: %v\n", answer)
	fmt.Printf("Running time: %v\n", end.Sub(start))
	start = time.Now()
	answer = run2(text)
	fmt.Printf("Answer 2: %v\n", answer)
	end = time.Now()
	fmt.Printf("Running time: %v\n", end.Sub(start))
}

func run(input string) string {
	instructions := strings.Split(strings.TrimSpace(input), "\n")
	value := 50
	count := 0
	for _, instr := range instructions {
		distance, err := strconv.Atoi(instr[1:])
		if err != nil {
			log.Fatalf("Invalid number in instruction: %s", instr[1:])
		}
		if instr[0] == 'L' {
			value = (100 + value - distance) % 100
		} else {
			value = (value + distance) % 100
		}
		if value == 0 {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

func run2(input string) string {
	instructions := strings.Split(strings.TrimSpace(input), "\n")
	value := 50
	count := 0
	for _, instr := range instructions {
		distance, err := strconv.Atoi(instr[1:])
		count += distance / 100
		distance = distance % 100
		if distance == 0 {
			continue
		}
		if err != nil {
			log.Fatalf("Invalid number in instruction: %s", instr[1:])
		}
		isZero := value == 0
		if instr[0] == 'L' {
			newValue := value - distance
			if newValue <= 0 {
				if !isZero {
					count++
				}
			}
			value = (newValue + 100) % 100
		} else if instr[0] == 'R' {
			newValue := value + distance
			if newValue >= 100 {
				count++
			}
			value = newValue % 100
		}
	}
	return fmt.Sprintf("%d", count)
}
