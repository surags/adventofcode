package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	part1(input)
	part2(input)
}

var cardMap = map[int]int{}

func part1(input []string) {
	points := 0
	for i, card := range input {
		cardMap[i+1] = 1
		nos := strings.Split(strings.Split(card, ":")[1], "|")
		winningNosS := strings.Split(strings.TrimSpace(nos[0]), " ")
		winMap := map[string]interface{}{}
		for _, wNo := range winningNosS {
			if wNo == "" {
				continue
			}
			winMap[wNo] = nil
		}

		roundPoints := 0
		roundNos := strings.Split(strings.TrimSpace(nos[1]), " ")
		for _, rNo := range roundNos {
			_, isAWin := winMap[rNo]
			if isAWin {
				if roundPoints == 0 {
					roundPoints = 1
				} else {
					roundPoints *= 2
				}
			}
		}

		points += roundPoints
	}

	fmt.Printf("Part 1: Points total: %d\n", points)
}

func part2(input []string) {
	totalNoOfCards := 0
	for i, card := range input {
		nos := strings.Split(strings.Split(card, ":")[1], "|")
		winningNosS := strings.Split(strings.TrimSpace(nos[0]), " ")
		winMap := map[string]interface{}{}
		for _, wNo := range winningNosS {
			if wNo == "" {
				continue
			}
			winMap[wNo] = nil
		}

		roundNos := strings.Split(strings.TrimSpace(nos[1]), " ")
		for a := 0; a < cardMap[i+1]; a++ {
			winId := i + 1
			for _, rNo := range roundNos {
				_, isAWin := winMap[rNo]
				if isAWin {
					winId++
					cardMap[winId]++
				}
			}
		}
	}

	for _, v := range cardMap {
		totalNoOfCards += v
	}

	fmt.Printf("Part 2: Total scratchcards: %d\n", totalNoOfCards)
}
