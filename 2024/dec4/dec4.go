package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec4.in")

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

	noOfXmas := part1(input)
	fmt.Println("Part 1: No of XMAS: " + strconv.Itoa(noOfXmas))

	noOfX_MAS := part2(input)
	fmt.Println("Part 2: No of X-MASs: " + strconv.Itoa(noOfX_MAS))
}

func checkRight(wordSearch [][]string, x, y int) bool {
	rightBound := len(wordSearch[x])

	if wordSearch[x][y] != "X" {
		return false
	}
	if y+1 >= rightBound || wordSearch[x][y+1] != "M" {
		return false
	}
	if y+2 >= rightBound || wordSearch[x][y+2] != "A" {
		return false
	}
	if y+3 >= rightBound || wordSearch[x][y+3] != "S" {
		return false
	}
	return true
}

func checkLeft(wordSearch [][]string, x, y int) bool {
	leftBound := 0

	if wordSearch[x][y] != "X" {
		return false
	}
	if y-1 < leftBound || wordSearch[x][y-1] != "M" {
		return false
	}
	if y-2 < leftBound || wordSearch[x][y-2] != "A" {
		return false
	}
	if y-3 < leftBound || wordSearch[x][y-3] != "S" {
		return false
	}
	return true
}

func checkUp(wordSearch [][]string, x, y int) bool {
	upperBound := 0

	if wordSearch[x][y] != "X" {
		return false
	}
	if x-1 < upperBound || wordSearch[x-1][y] != "M" {
		return false
	}
	if x-2 < upperBound || wordSearch[x-2][y] != "A" {
		return false
	}
	if x-3 < upperBound || wordSearch[x-3][y] != "S" {
		return false
	}
	return true
}

func checkDown(wordSearch [][]string, x, y int) bool {
	lowerBound := len(wordSearch)

	if wordSearch[x][y] != "X" {
		return false
	}
	if x+1 >= lowerBound || wordSearch[x+1][y] != "M" {
		return false
	}
	if x+2 >= lowerBound || wordSearch[x+2][y] != "A" {
		return false
	}
	if x+3 >= lowerBound || wordSearch[x+3][y] != "S" {
		return false
	}
	return true
}

func checkUpRight(wordSearch [][]string, x, y int) bool {
	upperBound := 0
	rightBound := len(wordSearch[x])

	if wordSearch[x][y] != "X" {
		return false
	}
	if x-1 < upperBound || y+1 >= rightBound || wordSearch[x-1][y+1] != "M" {
		return false
	}
	if x-2 < upperBound || y+2 >= rightBound || wordSearch[x-2][y+2] != "A" {
		return false
	}
	if x-3 < upperBound || y+3 >= rightBound || wordSearch[x-3][y+3] != "S" {
		return false
	}
	return true
}

func checkDownRight(wordSearch [][]string, x, y int) bool {
	lowerBound := len(wordSearch)
	rightBound := len(wordSearch[x])

	if wordSearch[x][y] != "X" {
		return false
	}
	if x+1 >= lowerBound || y+1 >= rightBound || wordSearch[x+1][y+1] != "M" {
		return false
	}
	if x+2 >= lowerBound || y+2 >= rightBound || wordSearch[x+2][y+2] != "A" {
		return false
	}
	if x+3 >= lowerBound || y+3 >= rightBound || wordSearch[x+3][y+3] != "S" {
		return false
	}
	return true
}

func checkUpLeft(wordSearch [][]string, x, y int) bool {
	leftBound := 0
	upperBound := 0

	if wordSearch[x][y] != "X" {
		return false
	}
	if x-1 < upperBound || y-1 < leftBound || wordSearch[x-1][y-1] != "M" {
		return false
	}
	if x-2 < upperBound || y-2 < leftBound || wordSearch[x-2][y-2] != "A" {
		return false
	}
	if x-3 < upperBound || y-3 < leftBound || wordSearch[x-3][y-3] != "S" {
		return false
	}
	return true
}

func checkDownLeft(wordSearch [][]string, x, y int) bool {
	leftBound := 0
	lowerBound := len(wordSearch)

	if wordSearch[x][y] != "X" {
		return false
	}
	if x+1 >= lowerBound || y-1 < leftBound || wordSearch[x+1][y-1] != "M" {
		return false
	}
	if x+2 >= lowerBound || y-2 < leftBound || wordSearch[x+2][y-2] != "A" {
		return false
	}
	if x+3 >= lowerBound || y-3 < leftBound || wordSearch[x+3][y-3] != "S" {
		return false
	}
	return true
}

func getNumberOfXmasAtIndex(wordSearch [][]string, x, y int) int {
	numberOfXmas := 0

	if checkRight(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkLeft(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkUp(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkDown(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkUpRight(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkDownRight(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkUpLeft(wordSearch, x, y) {
		numberOfXmas++
	}
	if checkDownLeft(wordSearch, x, y) {
		numberOfXmas++
	}

	return numberOfXmas
}

func part1(input []string) int {
	noOfXmas := 0
	wordSearch := [][]string{}

	for _, word := range input {
		wordline := []string{}
		for i := 0; i < len(word); i++ {
			wordline = append(wordline, string(word[i]))
		}
		wordSearch = append(wordSearch, wordline)
	}

	for i, word := range wordSearch {
		for j := 0; j < len(word); j++ {
			noOfXmas += getNumberOfXmasAtIndex(wordSearch, i, j)
		}
	}

	return noOfXmas
}

func checkDownRightMAS(wordSearch [][]string, x, y int) bool {
	leftBound := 0
	upperBound := 0
	rightBound := len(wordSearch[x])
	lowerBound := len(wordSearch)

	// check center A
	if wordSearch[x][y] != "A" {
		return false
	}

	// check MAS
	if x-1 >= upperBound && y-1 >= leftBound && wordSearch[x-1][y-1] == "M" {
		if x+1 < lowerBound && y+1 < rightBound && wordSearch[x+1][y+1] == "S" {
			return true
		}
	}

	// check SAM
	if x-1 >= upperBound && y-1 >= leftBound && wordSearch[x-1][y-1] == "S" {
		if x+1 < lowerBound && y+1 < rightBound && wordSearch[x+1][y+1] == "M" {
			return true
		}
	}

	return false
}

func checkUpRightMAS(wordSearch [][]string, x, y int) bool {
	leftBound := 0
	upperBound := 0
	rightBound := len(wordSearch[x])
	lowerBound := len(wordSearch)

	// check center A
	if wordSearch[x][y] != "A" {
		return false
	}

	// check MAS
	if x+1 < lowerBound && y-1 >= leftBound && wordSearch[x+1][y-1] == "M" {
		if x-1 >= upperBound && y+1 < rightBound && wordSearch[x-1][y+1] == "S" {
			return true
		}
	}

	// check SAM
	if x+1 < lowerBound && y-1 >= leftBound && wordSearch[x+1][y-1] == "S" {
		if x-1 >= upperBound && y+1 < rightBound && wordSearch[x-1][y+1] == "M" {
			return true
		}
	}

	return false
}

func isXMASAtIndex(wordSearch [][]string, x, y int) bool {
	// start at A in center and look for M-S / S-M
	if checkUpRightMAS(wordSearch, x, y) && checkDownRightMAS(wordSearch, x, y) {
		return true
	}

	return false
}

func part2(input []string) int {
	noOfXMAS := 0
	wordSearch := [][]string{}

	for _, word := range input {
		wordline := []string{}
		for i := 0; i < len(word); i++ {
			wordline = append(wordline, string(word[i]))
		}
		wordSearch = append(wordSearch, wordline)
	}

	for i, word := range wordSearch {
		for j := 0; j < len(word); j++ {
			if isXMASAtIndex(wordSearch, i, j) {
				noOfXMAS++
			}
		}
	}

	return noOfXMAS
}
