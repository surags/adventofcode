package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	testValue int
	numbers   []int
}

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec7.in")

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

	eqs := parseInput(input)

	totalCalibrationReqsWith2Ops := part1(eqs)
	fmt.Println("Part 1: Total calibration requirements with 2 ops: " + strconv.Itoa(totalCalibrationReqsWith2Ops))

	totalCalibrationReqsWith3Ops := part2(eqs)
	fmt.Println("Part 2: Total calibration requirements with 3 ops: " + strconv.Itoa(totalCalibrationReqsWith3Ops))
}

func parseInput(input []string) []equation {
	var eqs []equation
	for _, s := range input {
		eq := equation{}
		eqS := strings.Split(s, ":")
		eq.testValue, _ = strconv.Atoi(eqS[0])
		eq.numbers = []int{}
		numS := strings.Split(eqS[1], " ")
		for _, n := range numS {
			if n != "" {
				num, _ := strconv.Atoi(n)
				eq.numbers = append(eq.numbers, num)
			}
		}

		eqs = append(eqs, eq)
	}

	return eqs
}

func doesOpWork(op string, curTestValue, index int, eq equation) bool {
	if index == len(eq.numbers) {
		if curTestValue != eq.testValue {
			return false
		} else {
			return true
		}
	}
	if op == "+" {
		curTestValue += eq.numbers[index]
		return doesOpWork("+", curTestValue, index+1, eq) || doesOpWork("*", curTestValue, index+1, eq)
	} else {
		// op == "*"
		curTestValue *= eq.numbers[index]
		return doesOpWork("+", curTestValue, index+1, eq) || doesOpWork("*", curTestValue, index+1, eq)
	}
}

func isEqPossible(eq equation) bool {
	return doesOpWork("+", eq.numbers[0], 1, eq) || doesOpWork("*", eq.numbers[0], 1, eq)
}

func part1(eqs []equation) int {
	totalCalibrationReqs := 0
	for _, eq := range eqs {
		if isEqPossible(eq) {
			totalCalibrationReqs += eq.testValue
		}
	}
	return totalCalibrationReqs
}

func doesOpWorkP2(op string, curTestValue, index int, eq equation) bool {
	if index == len(eq.numbers) {
		if curTestValue != eq.testValue {
			return false
		} else {
			return true
		}
	}
	if op == "+" {
		curTestValue += eq.numbers[index]
		return doesOpWorkP2("+", curTestValue, index+1, eq) || doesOpWorkP2("*", curTestValue, index+1, eq) || doesOpWorkP2("||", curTestValue, index+1, eq)
	} else if op == "*" {
		curTestValue *= eq.numbers[index]
		return doesOpWorkP2("+", curTestValue, index+1, eq) || doesOpWorkP2("*", curTestValue, index+1, eq) || doesOpWorkP2("||", curTestValue, index+1, eq)
	} else {
		// op == "||"
		s := strconv.Itoa(curTestValue)
		s = s + strconv.Itoa(eq.numbers[index])
		curTestValue, _ = strconv.Atoi(s)
		return doesOpWorkP2("+", curTestValue, index+1, eq) || doesOpWorkP2("*", curTestValue, index+1, eq) || doesOpWorkP2("||", curTestValue, index+1, eq)
	}
}

func isEqPossibleP2(eq equation) bool {
	return doesOpWorkP2("+", eq.numbers[0], 1, eq) || doesOpWorkP2("*", eq.numbers[0], 1, eq) || doesOpWorkP2("||", eq.numbers[0], 1, eq)
}

func part2(eqs []equation) int {
	totalCalibrationReqs := 0
	for _, eq := range eqs {
		if isEqPossibleP2(eq) {
			totalCalibrationReqs += eq.testValue
		}
	}
	return totalCalibrationReqs
}
