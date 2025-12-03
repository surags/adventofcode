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
	file, err := os.Open("dec3.in")

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

	totalOutputJoltage := part1(input)
	fmt.Println("Part 1: Total output joltage: " + strconv.Itoa(totalOutputJoltage))

	totalOutputJoltage2 := part2(input)
	fmt.Println("Part 2: Total output joltage: " + strconv.Itoa(totalOutputJoltage2))
}

func calculateOutputJoltage(joltageRatings string) int {
	// implement logic here
	joltageRatingsInt := []int{}
	for _, v := range joltageRatings {
		j, _ := strconv.Atoi(string(v))
		joltageRatingsInt = append(joltageRatingsInt, j)
	}

	// joltage is the max sum of any 2 cells
	// using brute force
	maxJoltage := 0
	for i, v := range joltageRatingsInt {
		for j, w := range joltageRatingsInt {
			if i < j {
				if v*10+w > maxJoltage {
					maxJoltage = v*10 + w
				}
			}
		}
	}
	return maxJoltage
}

func part1(input []string) int {
	totalOutputJoltage := 0
	for _, v := range input {
		outputJoltage := calculateOutputJoltage(v)
		fmt.Println("Output joltage for ", v, " is ", outputJoltage)
		totalOutputJoltage += outputJoltage
	}
	return totalOutputJoltage
}

func convertSliceToNumber(slice []int) int {
	var numstrb strings.Builder
	for _, v := range slice {
		numstrb.WriteString(strconv.Itoa(v))
	}
	num, _ := strconv.Atoi(numstrb.String())
	return num
}

func convertNumberToSlice(num int) []int {
	numstr := strconv.Itoa(num)
	slice := []int{}
	for _, v := range numstr {
		digit, _ := strconv.Atoi(string(v))
		slice = append(slice, digit)
	}
	return slice
}

func dfs(joltageRatings []int, n int, currentPos int, currentNum []int, maxJoltage int) int {
	if len(currentNum) == n {
		// terminal case
		num := convertSliceToNumber(currentNum)
		if num > maxJoltage {
			// fmt.Println("New max found: ", num)
			return num
		} else {
			return maxJoltage
		}
	}

	if maxJoltage != 0 {
		// Attempt to terminate early if currentNum can never exceed maxJoltage
		num := convertSliceToNumber(currentNum)
		maxJoltageSub := convertSliceToNumber(convertNumberToSlice(maxJoltage)[0:len(currentNum)])
		if num < maxJoltageSub {
			// fmt.Println("Terminated early ", num, "current max ", maxJoltageSub)
			return maxJoltage
		}
	}

	// Old suboptimal logic
	// for i := currentPos + 1; i < len(joltageRatings); i++ {
	// 	newNum := currentNum[:]
	// 	subMaxJoltage := dfs(joltageRatings, n, i, append(newNum, joltageRatings[i]), maxJoltage)
	// 	if subMaxJoltage > maxJoltage {
	// 		maxJoltage = subMaxJoltage
	// 	}
	// }

	// Non terminal case. Try building number till n digits
	// As an optimization, try larger digits first so we can quickly prune small trees
	targetDig := 9
	for targetDig >= 0 {
		for i := currentPos + 1; i < len(joltageRatings); i++ {
			if joltageRatings[i] == targetDig {
				newNum := currentNum[:]
				subMaxJoltage := dfs(joltageRatings, n, i, append(newNum, joltageRatings[i]), maxJoltage)
				if subMaxJoltage > maxJoltage {
					maxJoltage = subMaxJoltage
				}
				if maxJoltage == 999999999999 {
					// max possible reached don't try harder
					return maxJoltage
				}
			}
		}
		targetDig--
	}
	return maxJoltage
}

func calculateOutputJoltageWithNCells(joltageRatings string, n int) int {
	// implement logic here
	joltageRatingsInt := []int{}
	for _, v := range joltageRatings {
		j, _ := strconv.Atoi(string(v))
		joltageRatingsInt = append(joltageRatingsInt, j)
	}

	// joltage is the max sum of any n cells
	// DFS to get all possible combinations and update maxJoltage
	maxJoltage := dfs(joltageRatingsInt, n, -1, []int{}, 0)
	return maxJoltage
}

func part2(input []string) int {
	totalOutputJoltage := 0
	for _, v := range input {
		outputJoltage := calculateOutputJoltageWithNCells(v, 12)
		fmt.Println("Output joltage for ", v, " is ", outputJoltage)
		totalOutputJoltage += outputJoltage
	}
	return totalOutputJoltage
}
