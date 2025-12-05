package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("sample.in")
	// file, err := os.Open("dec4.in")

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

	noOfPaperRolls := part1(input)
	fmt.Println("Part 1: Total Number of paper rolls accessible by forklift: " + strconv.Itoa(noOfPaperRolls))

	noOfPaperRollsRemoved := part2(input)
	fmt.Println("Part 2: Total Number of paper rolls removed by forklift: " + strconv.Itoa(noOfPaperRollsRemoved))
}

func part1(input []string) int {
	noOfPaperRolls := 0

	for i, v := range input {
		for j, w := range v {
			if w == '@' {
				// I am a paper roll. Check 8 rolls around me
				rollsAround := 0

				// check up left
				if i > 0 && j > 0 && input[i-1][j-1] == '@' {
					rollsAround++
				}
				// check up middle
				if i > 0 && input[i-1][j] == '@' {
					rollsAround++
				}
				// check up right
				if i > 0 && j < len(v)-1 && input[i-1][j+1] == '@' {
					rollsAround++
				}
				// check left
				if j > 0 && input[i][j-1] == '@' {
					rollsAround++
				}
				// check right
				if j < len(v)-1 && input[i][j+1] == '@' {
					rollsAround++
				}
				// check down left
				if i < len(input)-1 && j > 0 && input[i+1][j-1] == '@' {
					rollsAround++
				}
				// check down middle
				if i < len(input)-1 && input[i+1][j] == '@' {
					rollsAround++
				}
				// check down right
				if i < len(input)-1 && j < len(v)-1 && input[i+1][j+1] == '@' {
					rollsAround++
				}

				if rollsAround < 4 {
					noOfPaperRolls++
				}
			}
		}
	}

	return noOfPaperRolls
}

func part2(input []string) int {
	noOfPaperRollsRemoved := 0

	noOfRemovablePaperRolls := -1
	for noOfRemovablePaperRolls != 0 {
		noOfRemovablePaperRolls = 0
		toRemove := [][]int{}
		for i, v := range input {
			for j, w := range v {
				if w == '@' {
					// I am a paper roll. Check 8 rolls around me
					rollsAround := 0

					// check up left
					if i > 0 && j > 0 && input[i-1][j-1] == '@' {
						rollsAround++
					}
					// check up middle
					if i > 0 && input[i-1][j] == '@' {
						rollsAround++
					}
					// check up right
					if i > 0 && j < len(v)-1 && input[i-1][j+1] == '@' {
						rollsAround++
					}
					// check left
					if j > 0 && input[i][j-1] == '@' {
						rollsAround++
					}
					// check right
					if j < len(v)-1 && input[i][j+1] == '@' {
						rollsAround++
					}
					// check down left
					if i < len(input)-1 && j > 0 && input[i+1][j-1] == '@' {
						rollsAround++
					}
					// check down middle
					if i < len(input)-1 && input[i+1][j] == '@' {
						rollsAround++
					}
					// check down right
					if i < len(input)-1 && j < len(v)-1 && input[i+1][j+1] == '@' {
						rollsAround++
					}

					if rollsAround < 4 {
						noOfRemovablePaperRolls++
						toRemove = append(toRemove, []int{i, j})
					}
				}
			}
		}

		// Remove paper rolls for next round
		for _, v := range toRemove {
			i, j := v[0], v[1]
			line := []rune(input[i])
			line[j] = '.'
			input[i] = string(line)
		}
		noOfPaperRollsRemoved += noOfRemovablePaperRolls
	}

	return noOfPaperRollsRemoved
}
