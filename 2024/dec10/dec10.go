package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type set map[string]bool

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec10.in")

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

	topoMap := parseInput(input)

	sumTrailScores := part1(topoMap)
	fmt.Println("Part 1: Sum of Trail Scores: " + strconv.Itoa(sumTrailScores))

	sumTrailRatings := part2(topoMap)
	fmt.Println("Part 2: Sum of Trail Ratings: " + strconv.Itoa(sumTrailRatings))
}

func parseInput(input []string) [][]int {
	topoMap := [][]int{}
	for _, line := range input {
		heights := []int{}
		for _, heightS := range line {
			height, _ := strconv.Atoi(string(heightS))
			heights = append(heights, height)
		}
		topoMap = append(topoMap, heights)
	}

	return topoMap
}

func getSetKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func traverseTrail(topoMap [][]int, destSet set, curX, curY, nextX, nextY int) int {
	if nextX < 0 || nextX >= len(topoMap) || nextY < 0 || nextY >= len(topoMap[0]) {
		// out of bounds
		return 0
	}

	if topoMap[nextX][nextY]-topoMap[curX][curY] != 1 {
		// unreachable
		return 0
	}

	if topoMap[nextX][nextY] == 9 {
		// peak
		if destSet != nil {
			// trail score
			key := getSetKey(nextX, nextY)
			if _, found := destSet[key]; !found {
				destSet[key] = true
				return 1
			}
		} else {
			// trail rating
			return 1
		}
	}

	return traverseTrail(topoMap, destSet, nextX, nextY, nextX-1, nextY) + // up
		traverseTrail(topoMap, destSet, nextX, nextY, nextX+1, nextY) + // down
		traverseTrail(topoMap, destSet, nextX, nextY, nextX, nextY-1) + // left
		traverseTrail(topoMap, destSet, nextX, nextY, nextX, nextY+1) // right
}

func calculateTrailScore(topoMap [][]int, height, startX, startY int) int {
	if height != 0 {
		// Not a trail head
		return 0
	}

	destSet := map[string]bool{}
	return traverseTrail(topoMap, destSet, startX, startY, startX-1, startY) + // up
		traverseTrail(topoMap, destSet, startX, startY, startX+1, startY) + // down
		traverseTrail(topoMap, destSet, startX, startY, startX, startY-1) + // left
		traverseTrail(topoMap, destSet, startX, startY, startX, startY+1) // right
}

func part1(topoMap [][]int) int {
	sumTrailScores := 0
	for i, lines := range topoMap {
		for j, height := range lines {
			sumTrailScores += calculateTrailScore(topoMap, height, i, j)
		}
	}

	return sumTrailScores
}

func calculateTrailRating(topoMap [][]int, height, startX, startY int) int {
	if height != 0 {
		// Not a trail head
		return 0
	}

	return traverseTrail(topoMap, nil, startX, startY, startX-1, startY) + // up
		traverseTrail(topoMap, nil, startX, startY, startX+1, startY) + // down
		traverseTrail(topoMap, nil, startX, startY, startX, startY-1) + // left
		traverseTrail(topoMap, nil, startX, startY, startX, startY+1) // right
}

func part2(topoMap [][]int) int {
	sumTrailRatings := 0
	for i, lines := range topoMap {
		for j, height := range lines {
			sumTrailRatings += calculateTrailRating(topoMap, height, i, j)
		}
	}

	return sumTrailRatings
}
