package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("sample.in")
	file, err := os.Open("dec7.in")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()

	noOfSplits := part1(input)
	fmt.Println("Part 1: Number of beam splits: " + strconv.Itoa(noOfSplits))

	noOfTimelines := part2(input)
	fmt.Println("Part 2: Number of timelines: " + strconv.Itoa(noOfTimelines))
}

func generateGraph(input []string) ([][]string, [][]bool, int, int) {
	var graph [][]string
	var visited [][]bool
	var startX, startY int
	for i, v := range input {
		row := strings.Split(v, "")
		visitedRow := make([]bool, len(row))
		for j := range visitedRow {
			visitedRow[j] = false
		}
		for j := range row {
			if row[j] == "S" {
				startX, startY = i, j
			}
		}
		graph = append(graph, row)
		visited = append(visited, visitedRow)
	}
	return graph, visited, startX, startY
}

func followBeam(graph [][]string, visited [][]bool, x int, y int) int {
	noOfSplits := 0
	if x >= len(graph) || y >= len(graph[0]) || y < 0 {
		// Out of bounds
		return noOfSplits
	}
	if visited[x][y] {
		// Already visited
		return noOfSplits
	}
	visited[x][y] = true
	if graph[x][y] == "." || graph[x][y] == "S" {
		// Beam can pass through here
		noOfSplits += followBeam(graph, visited, x+1, y)
	}

	if graph[x][y] == "^" {
		// Beam splits here
		noOfSplits++
		noOfSplits += followBeam(graph, visited, x, y-1)
		noOfSplits += followBeam(graph, visited, x, y+1)
	}

	return noOfSplits
}

func printGraph(graph [][]string, visited [][]bool) {
	for _, v := range graph {
		fmt.Println(strings.Join(v, ""))
	}
	fmt.Println("")
	for _, v := range visited {
		for _, b := range v {
			if b {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func part1(input []string) int {
	noOfSplits := 0

	graph, visited, startX, startY := generateGraph(input)

	noOfSplits = followBeam(graph, visited, startX, startY)
	printGraph(graph, visited)

	return noOfSplits
}

func followBeamAltTimelines(graph [][]string, visitedTimelineCache [][]int, x int, y int) int {
	noOfSplits := 0 // original timeline
	if x >= len(graph) || y >= len(graph[0]) || y < 0 {
		// Out of bounds
		return noOfSplits
	}

	if visitedTimelineCache[x][y] >= 0 {
		// Already calculated timelines from here
		return visitedTimelineCache[x][y]
	}

	if graph[x][y] == "." || graph[x][y] == "S" {
		// Beam can pass through here
		noOfSplits += followBeamAltTimelines(graph, visitedTimelineCache, x+1, y)
	}

	if graph[x][y] == "^" {
		// Beam splits here
		noOfSplits++
		noOfSplits += followBeamAltTimelines(graph, visitedTimelineCache, x, y-1)
		noOfSplits += followBeamAltTimelines(graph, visitedTimelineCache, x, y+1)
	}

	visitedTimelineCache[x][y] = noOfSplits
	return noOfSplits
}

func part2(input []string) int {
	noOfSplits := 0

	graph, _, startX, startY := generateGraph(input)
	visitedTimelineCache := make([][]int, len(graph))
	for i := range visitedTimelineCache {
		visitedTimelineCache[i] = make([]int, len(graph[0]))
		for j := range visitedTimelineCache[i] {
			visitedTimelineCache[i][j] = -1
		}
	}

	noOfSplits = followBeamAltTimelines(graph, visitedTimelineCache, startX, startY)

	return noOfSplits + 1 // including original timeline
}
