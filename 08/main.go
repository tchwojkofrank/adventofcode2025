package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

type Point struct {
	X int
	Y int
	Z int
}

type Pair struct {
	A Point
	B Point
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p Point) Distance(q Point) int {
	dx := p.X - q.X
	dy := p.Y - q.Y
	dz := p.Z - q.Z
	return dx*dx + dy*dy + dz*dz
}

func getPoints(input string) []Point {
	var points []Point
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var p Point
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		points = append(points, p)
	}
	return points
}

func getDistances(points []Point) (map[Pair]int, []Pair) {
	pairs := make([]Pair, 0)
	distances := make(map[Pair]int)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pair := Pair{A: points[i], B: points[j]}
			distances[pair] = points[i].Distance(points[j])
			pairs = append(pairs, pair)
		}
	}

	// sort the pairs by distance
	sort.Slice(pairs, func(a, b int) bool {
		return distances[pairs[a]] < distances[pairs[b]]
	})

	return distances, pairs
}

func appendConnection(connections map[Point][]Point, a Point, b Point) map[Point][]Point {
	// check if a is already in connections
	if _, ok := connections[a]; !ok {
		connections[a] = make([]Point, 0)
	}
	// check if b is already in connections[a]
	found := false
	for _, p := range connections[a] {
		if p == b {
			found = true
			break
		}
	}
	if !found {
		connections[a] = append(connections[a], b)
	}
	return connections
}

func getCiruitSize(connections map[Point][]Point, start Point, visited map[Point]struct{}) int {
	if _, ok := visited[start]; ok {
		return 0
	}
	visited[start] = struct{}{}
	size := 1
	for _, neighbor := range connections[start] {
		size += getCiruitSize(connections, neighbor, visited)
	}
	return size
}

func run(input string) string {
	points := getPoints(input)
	limit := 10
	if len(points) > 40 {
		limit = 1000
	}

	_, sortedPairs := getDistances(points)
	connections := make(map[Point][]Point)
	for i := 0; i < limit; i++ {
		pair := sortedPairs[i]
		connections = appendConnection(connections, pair.A, pair.B)
		connections = appendConnection(connections, pair.B, pair.A)
	}
	visited := make(map[Point]struct{})
	sizes := [3]int{}
	for _, point := range points {
		size := getCiruitSize(connections, point, visited)
		if size > sizes[0] {
			sizes[2] = sizes[1]
			sizes[1] = sizes[0]
			sizes[0] = size
		} else if size > sizes[1] {
			sizes[2] = sizes[1]
			sizes[1] = size
		} else if size > sizes[2] {
			sizes[2] = size
		}
	}
	return fmt.Sprintf("%d", sizes[0]*sizes[1]*sizes[2])
}

func run2(input string) string {
	points := getPoints(input)

	_, sortedPairs := getDistances(points)
	connections := make(map[Point][]Point)
	size := 0
	lastPair := Pair{}
	for i := 0; i < len(sortedPairs) && size < len(sortedPairs); i++ {
		pair := sortedPairs[i]
		connections = appendConnection(connections, pair.A, pair.B)
		connections = appendConnection(connections, pair.B, pair.A)
		visited := make(map[Point]struct{})
		size = getCiruitSize(connections, pair.A, visited)
		if size >= len(points) {
			lastPair = pair
			break
		}
	}
	fmt.Printf("Last pair: %v Value: %v\n", lastPair, lastPair.A.X*lastPair.B.X)
	return fmt.Sprintf("%d", lastPair.A.X*lastPair.B.X)
}
