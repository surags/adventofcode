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
	file, err := os.Open("dec2.in")

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

	invalidIdSum := part1(input)
	fmt.Println("Part 1: Invalid ID Sum: " + strconv.Itoa(invalidIdSum))

	invalidIdSum2 := part2(input)
	fmt.Println("Part 2: Invalid ID Sum: " + strconv.Itoa(invalidIdSum2))
}

func checkValidP1(id string) bool {
	// implement validation logic here
	if len(id)%2 != 0 {
		// odd length IDs are always valid
		return true
	}

	if id[0:len(id)/2] == id[len(id)/2:] {
		// invalid ID
		return false
	}
	return true
}

// calculate invalid ID sum
func part1(input []string) int {
	invalidIdSum := 0
	ranges := strings.Split(input[0], ",")
	for _, v := range ranges {
		r := strings.Split(v, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			if !checkValidP1(id) {
				fmt.Println("Invalid ID: ", id)
				invalidIdSum += i
			}
		}
	}
	return invalidIdSum
}

func checkValidP2(id string) bool {
	// 1234567890
	substringSizes := []int{}
	for v := range (len(id) / 2) + 1 {
		if v == 0 {
			continue
		}

		if len(id)%v == 0 {
			substringSizes = append(substringSizes, v)
		}
	}
	// fmt.Println("Substring sizes: ", substringSizes)

	// Try all substring sizes
	for _, v := range substringSizes {
		set := map[string]interface{}{}
		for i := 0; i < len(id); i += v {
			set[id[i:i+v]] = nil
		}
		// fmt.Println("Substrings: ", set)
		if len(set) == 1 {
			return false
		}
	}

	return true
}

// calculate invalid ID sum
func part2(input []string) int {
	invalidIdSum := 0
	ranges := strings.Split(input[0], ",")
	for _, v := range ranges {
		r := strings.Split(v, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		fmt.Println("Checking range: ", start, end)
		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			if !checkValidP2(id) {
				fmt.Println("Invalid ID: ", id)
				invalidIdSum += i
			}
		}
	}
	// invalidIdSum := 0
	// fmt.Println(checkValidP2("99"))
	// fmt.Println(checkValidP2("111"))
	// fmt.Println(checkValidP2("1188511885"))
	return invalidIdSum
}
