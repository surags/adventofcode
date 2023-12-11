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

	part1(input)
	part2(input)
}

func isRoundPossible(r, g, b, maxR, maxG, maxB int) bool {
	if r <= maxR && g <= maxG && b <= maxB {
		return true
	}
	return false
}

func isGamePossible(game string, id int) bool {
	rounds := strings.Split(game, ";")

	for _, round := range rounds {
		// r := regexp.MustCompile(`(?P<Red>[])`)
		dice := strings.Split(round, ",")
		r, g, b := 0, 0, 0

		for _, d := range dice {
			dCleaned := strings.TrimSpace(d)
			val := strings.Split(dCleaned, " ")
			if val[1] == "red" {
				r, _ = strconv.Atoi(val[0])
			}

			if val[1] == "green" {
				g, _ = strconv.Atoi(val[0])
			}

			if val[1] == "blue" {
				b, _ = strconv.Atoi(val[0])
			}
		}

		res := isRoundPossible(r, g, b, 12, 13, 14)
		if res == false {
			return false
		}
	}

	return true
}

func part1(input []string) {
	sumPossible := 0
	for _, game := range input {
		split := strings.Split(game, ":")
		gameTitle := strings.Split(split[0], " ")
		gameId, _ := strconv.Atoi(gameTitle[1])
		gameRounds := split[1]

		res := isGamePossible(gameRounds, gameId)
		// fmt.Printf("GAME: %d Result: %v\n", gameId, res)

		if res == true {
			sumPossible = sumPossible + gameId
		}
	}

	fmt.Printf("Part 1: %d\n", sumPossible)
}

func getPower(game string, id int) int {
	minR, minG, minB := 0, 0, 0
	rounds := strings.Split(game, ";")

	for _, round := range rounds {
		dice := strings.Split(round, ",")
		r, g, b := 0, 0, 0

		for _, d := range dice {
			dCleaned := strings.TrimSpace(d)
			val := strings.Split(dCleaned, " ")
			if val[1] == "red" {
				r, _ = strconv.Atoi(val[0])
			}

			if val[1] == "green" {
				g, _ = strconv.Atoi(val[0])
			}

			if val[1] == "blue" {
				b, _ = strconv.Atoi(val[0])
			}
		}

		minR = max(minR, r)
		minG = max(minG, g)
		minB = max(minB, b)
	}

	return minR * minG * minB
}

func part2(input []string) {
	powerSum := 0
	for _, game := range input {
		split := strings.Split(game, ":")
		gameTitle := strings.Split(split[0], " ")
		gameId, _ := strconv.Atoi(gameTitle[1])
		gameRounds := split[1]

		res := getPower(gameRounds, gameId)
		// fmt.Printf("GAME: %d Result: %v\n", gameId, res)

		powerSum += res
	}

	fmt.Printf("Part 2: %d\n", powerSum)
}
