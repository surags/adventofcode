package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("sample.in")
	file, err := os.Open("dec5.in")

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

	freshIngredients := part1(input)
	fmt.Println("Part 1: Number of fresh ingredients: " + strconv.Itoa(freshIngredients))

	possibleFreshIngredients := part2(input)
	fmt.Println("Part 2: Total Number of fresh ingredients: " + strconv.Itoa(possibleFreshIngredients))

	// noOfPaperRollsRemoved := part2(input)
	// fmt.Println("Part 2: Total Number of paper rolls removed by forklift: " + strconv.Itoa(noOfPaperRollsRemoved))
}

func isFresh(ingredient int, ranges [][]int) bool {
	for _, v := range ranges {
		if ingredient >= v[0] && ingredient <= v[1] {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	freshIngredients := 0

	ranges := [][]int{}

	restart := 0
	for i, v := range input {
		if v == "" {
			restart = i
			break
		}
		rangeS := strings.Split(v, "-")
		start, _ := strconv.Atoi(rangeS[0])
		end, _ := strconv.Atoi(rangeS[1])
		ranges = append(ranges, []int{start, end})
	}

	for i := restart + 1; i < len(input); i++ {
		ingredient, _ := strconv.Atoi(input[i])
		if isFresh(ingredient, ranges) {
			freshIngredients++
		}
	}

	return freshIngredients
}

func collapseRanges(ranges [][]int) [][]int {
	// implement range collapsing logic here
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})

	// slice sorted based on ascending order. Now collapse overlapping ranges
	remove := map[int]bool{}
	for i := 0; i < len(ranges)-1; i++ {
		if remove[i] {
			continue
		}
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j][0] <= ranges[i][1] {
				// overlapping ranges
				ranges[i][1] = max(ranges[i][1], ranges[j][1])
				remove[j] = true
			} else {
				// no more overlapping ranges
				break
			}
		}
	}

	collapsedRanges := [][]int{}
	for i, v := range ranges {
		if !remove[i] {
			collapsedRanges = append(collapsedRanges, v)
		}
	}

	return collapsedRanges
}

func part2(input []string) int {
	freshIngredients := 0

	ranges := [][]int{}

	for _, v := range input {
		if v == "" {
			break
		}
		rangeS := strings.Split(v, "-")
		start, _ := strconv.Atoi(rangeS[0])
		end, _ := strconv.Atoi(rangeS[1])
		ranges = append(ranges, []int{start, end})
	}

	collapsedRanges := collapseRanges(ranges)

	for _, v := range collapsedRanges {
		freshIngredients += v[1] - v[0] + 1
	}

	return freshIngredients
}
