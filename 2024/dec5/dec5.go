package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type set map[int]bool

func main() {
	//file, err := os.Open("sample.in")
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

	ordering, updates := parseInput(input)

	sumValidMiddlePageNum := part1(ordering, updates)
	fmt.Println("Part 1: Sum of ordered middle page nums: " + strconv.Itoa(sumValidMiddlePageNum))

	sumCorrectedMiddlePageNum := part2(ordering, updates)
	fmt.Println("Part 2: Sum of corrected unordered middle page nums: " + strconv.Itoa(sumCorrectedMiddlePageNum))
}

func parseInput(input []string) (map[int]set, [][]int) {
	ordering := map[int]set{}
	updates := [][]int{}

	index := 0
	for i, s := range input {
		index = i
		if s == "" {
			break
		}

		lineSplit := strings.Split(s, "|")
		p1, _ := strconv.Atoi(lineSplit[0])
		p2, _ := strconv.Atoi(lineSplit[1])

		if ordering[p1] == nil {
			ordering[p1] = set{}
		}

		ordering[p1][p2] = true
	}

	for i := index + 1; i < len(input); i++ {
		line := input[i]
		lineSplit := strings.Split(line, ",")
		update := []int{}
		for _, s := range lineSplit {
			v, _ := strconv.Atoi(s)
			update = append(update, v)
		}
		updates = append(updates, update)
	}

	return ordering, updates
}

func isBefore(ordering map[int]set, x, y int) bool {
	_, found := ordering[x][y]
	if !found {
		//for k, _ := range ordering[x] {
		//	//subsetFound := isBefore(ordering, k, y)
		//	//if subsetFound {
		//	//	return true
		//	//}
		//}
	} else {
		return true
	}

	return false
}

func checkValid(ordering map[int]set, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		j := i + 1
		if !isBefore(ordering, update[i], update[j]) {
			return false
		}
	}
	return true
}

func part1(ordering map[int]set, updates [][]int) int {
	sumOfValidMiddlePageNum := 0

	for _, update := range updates {
		if checkValid(ordering, update) {
			sumOfValidMiddlePageNum += update[len(update)/2]
			//fmt.Printf("Valid line: %v, num: %d \n", update, update[len(update)/2])
		}
	}

	return sumOfValidMiddlePageNum
}

func part2(ordering map[int]set, updates [][]int) int {
	sumCorrectedMiddlePageNum := 0

	for _, update := range updates {
		if !checkValid(ordering, update) {
			// "sort" pages
			sort.Slice(update, func(i, j int) bool {
				_, ok := ordering[update[i]][update[j]]
				return ok
			})
			//fmt.Printf("Valid line: %v, num: %d \n", update, update[len(update)/2])
			sumCorrectedMiddlePageNum += update[len(update)/2]
		}
	}

	return sumCorrectedMiddlePageNum
}
