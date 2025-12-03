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
	run(text)
	end := time.Now()
	fmt.Printf("Running time: %v\n", end.Sub(start))
	start = time.Now()
	run2(text)
	end = time.Now()
	fmt.Printf("Running time: %v\n", end.Sub(start))
}

func bankJoltage(bank []int) int {
	joltage := 0
	max10 := 0
	for high := 0; high < len(bank)-1; high++ {
		if bank[high] > max10 {
			max10 = bank[high]
			max1 := 0
			for low := high + 1; low < len(bank); low++ {
				if bank[low] > max1 {
					max1 = bank[low]
					joltage = 10*max10 + max1
				}
			}
		}
	}
	return joltage
}

func bankOverrideJoltage(bank []int) int {
	joltage := 0
	exp := 100000000000
	index := 0
	for power := 11; power >= 0; power-- {
		maxDigit := 0
		for i := index; i < len(bank)-power; i++ {
			if bank[i] > maxDigit {
				maxDigit = bank[i]
				index = i + 1
			}
		}
		joltage += maxDigit * exp
		exp /= 10
	}

	return joltage
}

func getBankValues(bank string) []int {
	values := []int{}
	for _, ch := range bank {
		values = append(values, int(ch-'0'))
	}
	return values
}

func run(input string) string {
	banks := strings.Split(input, "\n")
	joltage := 0
	for _, bank := range banks {
		bankValues := getBankValues(bank)
		joltage += bankJoltage(bankValues)
	}
	fmt.Printf("Total joltage: %d\n", joltage)
	return fmt.Sprintf("%d", joltage)
}

func run2(input string) string {
	banks := strings.Split(input, "\n")
	joltage := 0
	for _, bank := range banks {
		bankValues := getBankValues(bank)
		joltage += bankOverrideJoltage(bankValues)
	}
	fmt.Printf("Total override joltage: %d\n", joltage)
	return fmt.Sprintf("%d", joltage)
}
