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
	file, err := os.Open("dec9.in")

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
	sequences := generateSequences(input)
	part1(sequences)
	part2(sequences)
}

func generateSequences(input []string) [][]int {
	sequences := [][]int{}

	for _, line := range input {
		sequence := []int{}
		numStrs := strings.Split(line, " ")

		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			sequence = append(sequence, num)
		}

		sequences = append(sequences, sequence)
	}

	return sequences
}

func parseLevelDiff(sequence []int) int {
	isZero := true
	diffArr := []int{}

	for i, num := range sequence {
		if num != 0 {
			isZero = false
		}

		if i != 0 {
			diffArr = append(diffArr, sequence[i]-sequence[i-1])
		}
	}

	if !isZero {
		return sequence[len(sequence)-1] + parseLevelDiff(diffArr)
	}

	return 0
}

func parseLevelDiff0(sequence []int) int {
	isZero := true
	diffArr := []int{}

	for i, num := range sequence {
		if num != 0 {
			isZero = false
		}

		if i != 0 {
			diffArr = append(diffArr, sequence[i]-sequence[i-1])
		}
	}

	if !isZero {
		return sequence[0] - parseLevelDiff0(diffArr)
	}

	return 0
}

func part1(sequences [][]int) {
	sum := 0
	for _, sequence := range sequences {
		sum += parseLevelDiff(sequence)
	}

	fmt.Printf("Part 1: Sum of diffs: %d\n", sum)
}

func part2(sequences [][]int) {
	sum := 0
	for _, sequence := range sequences {
		sum += parseLevelDiff0(sequence)
	}

	fmt.Printf("Part 2: Sum of diffs: %d\n", sum)
}
