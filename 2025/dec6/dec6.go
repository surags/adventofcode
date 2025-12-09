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
	file, err := os.Open("dec6.in")

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

	homeworkAnswer := part1(input)
	fmt.Println("Part 1: Sum of homework answers: " + strconv.Itoa(homeworkAnswer))

	homeworkAnswer = part2(input)
	fmt.Println("Part 2: Sum of homework answers: " + strconv.Itoa(homeworkAnswer))
}

func part1(input []string) int {
	var operations []string
	var curValue []int

	operationsS := input[len(input)-1]
	operationSplit := strings.Split(operationsS, " ")
	for _, v := range operationSplit {
		// v = strings.Trim(v, " ")
		if v == "" {
			continue
		}
		operations = append(operations, v)
	}

	fmt.Println("Operations: ", operations)
	curValue = make([]int, len(operations))
	for i := 0; i < len(input)-1; i++ {
		line := input[i]
		lineSplit := strings.Split(line, " ")
		index := 0
		for _, v := range lineSplit {
			if v == "" {
				continue
			}
			if i != 0 {
				if operations[index] == "+" {
					val, _ := strconv.Atoi(v)
					curValue[index] += val
				} else if operations[index] == "*" {
					val, _ := strconv.Atoi(v)
					curValue[index] *= val
				}
			} else {
				val, _ := strconv.Atoi(v)
				curValue[index] = val
			}
			index++
		}
	}

	sum := 0
	for _, v := range curValue {
		sum += v
	}

	return sum
}

func part2(input []string) int {
	// Read right to left for each line till you see an operation then perform
	pos := len(input[0]) - 1
	operationIndex := len(input) - 1
	homeworkSum := 0

	nums := []string{}
	for pos >= 0 {

		sb := strings.Builder{}
		for i := 0; i < operationIndex; i++ {
			c := input[i][pos]
			if c == ' ' {
				continue
			}
			sb.WriteByte(c)
		}
		nums = append(nums, sb.String())

		if input[operationIndex][pos] == '+' || input[operationIndex][pos] == '*' {
			// perform operation
			answer, _ := strconv.Atoi(nums[0])
			for i, v := range nums {
				if i == 0 {
					continue
				}
				if input[operationIndex][pos] == '+' {
					val, _ := strconv.Atoi(v)
					answer += val
				} else {
					val, _ := strconv.Atoi(v)
					answer *= val
				}
			}
			// fmt.Println("Operands: ", nums, " Answer: ", answer)
			homeworkSum += answer
			// reset num
			nums = []string{}
			// Skip one for space
			pos--
		}

		pos--
	}

	return homeworkSum
}
