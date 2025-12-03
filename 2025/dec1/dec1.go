package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("sample.in")
	file, err := os.Open("dec1.in")

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

	noOfTimesAtZero := part1(input)
	fmt.Println("Part 1: Password: " + strconv.Itoa(noOfTimesAtZero))

	noOfTimesAtZeroWithIntermediate := part2(input)
	fmt.Println("Part 2: Password: " + strconv.Itoa(noOfTimesAtZeroWithIntermediate))
}

// calculate no of times at zero
func part1(input []string) int {
	dialPosition := 50
	dir := 1
	noOfTimesAtZero := 0
	for _, v := range input {
		line := v[1:]
		if v[0] == 'L' {
			dir = -1
		} else {
			dir = 1
		}
		rot, _ := strconv.Atoi(line)
		dialPosition += dir * (rot % 100)
		dialPosition = dialPosition % 100
		if dialPosition < 0 {
			dialPosition += 100
		}
		fmt.Println("Dial position: ", dialPosition)
		if dialPosition == 0 {
			noOfTimesAtZero++
		}
	}
	return noOfTimesAtZero
}

// calculate no of times at zero
func part2(input []string) int {
	dialPosition := 50
	dir := 1
	noOfTimesAtZero := 0
	for _, v := range input {
		line := v[1:]
		if v[0] == 'L' {
			dir = -1
		} else {
			dir = 1
		}
		rot, _ := strconv.Atoi(line)
		for rot > 0 {
			dialPosition += dir
			rot--
			if dialPosition == 100 {
				dialPosition = 0
			} else if dialPosition < 0 {
				dialPosition = 99
			}
			// fmt.Println("Dial position: ", dialPosition)
			if dialPosition == 0 {
				noOfTimesAtZero++
			}
		}
		fmt.Println("Dial position: ", dialPosition)
	}
	return noOfTimesAtZero
}
