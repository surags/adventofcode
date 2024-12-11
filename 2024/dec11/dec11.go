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
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec11.in")

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

	stones := parseInput(input)

	noOfStones := part1(stones)
	fmt.Println("Part 1: Number of Stones (25 blinks): " + strconv.Itoa(noOfStones))

	noOfStonesP2 := part2(stones)
	fmt.Println("Part 2: Number of Stones (75 blinks): " + strconv.Itoa(noOfStonesP2))
}

func parseInput(input []string) []int {
	stones := []int{}
	line := strings.Split(input[0], " ")
	for _, stoneS := range line {
		stone, _ := strconv.Atoi(string(stoneS))
		stones = append(stones, stone)
	}

	return stones
}

func blinkNTimes(stones []int, times int) []int {
	stonesIter := stones
	for i := 0; i < times; i++ {
		newStones := []int{}
		for _, stone := range stonesIter {

			// Stone 0
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			// Even digits stone
			stoneS := strconv.Itoa(stone)
			if len(stoneS)%2 == 0 {
				a := stoneS[0 : len(stoneS)/2]
				aI, _ := strconv.Atoi(a)
				b := stoneS[len(stoneS)/2:]
				bI, _ := strconv.Atoi(b)
				newStones = append(newStones, aI, bI)
				continue
			}

			// stone * 2024
			newStones = append(newStones, stone*2024)
		}
		stonesIter = newStones
	}

	return stonesIter
}

func part1(stones []int) int {
	stoneIter := []int{}
	for _, stone := range stones {
		stoneIter = append(stoneIter, stone)
	}

	newStones := blinkNTimes(stoneIter, 25)

	return len(newStones)
}

type node struct {
	i    int
	next *node
}

func blinkNTimesRecurse(stone int, iterNum int, limit int, countCache map[string]int) int {
	if iterNum == limit {
		return 1
	}

	if c, found := countCache[getCacheKey(stone, limit-iterNum)]; found {
		return c
	}

	if stone == 0 {
		stone = 1
		count := blinkNTimesRecurse(stone, iterNum+1, limit, countCache)
		countCache[getCacheKey(stone, limit-(iterNum+1))] = count
		return count
	}

	stoneS := strconv.Itoa(stone)
	if len(stoneS)%2 == 0 {
		a := stoneS[0 : len(stoneS)/2]
		aI, _ := strconv.Atoi(a)
		b := stoneS[len(stoneS)/2:]
		bI, _ := strconv.Atoi(b)
		stone = aI
		node1Count := blinkNTimesRecurse(stone, iterNum+1, limit, countCache)
		countCache[getCacheKey(stone, limit-(iterNum+1))] = node1Count
		stone = bI
		node2Count := blinkNTimesRecurse(stone, iterNum+1, limit, countCache)
		countCache[getCacheKey(stone, limit-(iterNum+1))] = node2Count
		return node1Count + node2Count
	}

	stone = stone * 2024
	count := blinkNTimesRecurse(stone, iterNum+1, limit, countCache)
	countCache[getCacheKey(stone, limit-(iterNum+1))] = count
	return count
}

func getCacheKey(nodeI, iterCount int) string {
	return fmt.Sprintf("%d-%d", nodeI, iterCount)
}

func part2(stones []int) int {
	countCache := map[string]int{}
	count := 0
	for _, stone := range stones {
		count += blinkNTimesRecurse(stone, 0, 75, countCache)
	}

	return count
}
