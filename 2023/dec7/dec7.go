package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	hand string
	bid  int
}

type ByHand []hand
type ByHandP2 []hand

func (b ByHand) Len() int           { return len(b) }
func (b ByHand) Less(i, j int) bool { return isLess(b[i], b[j]) }
func (b ByHand) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func (b ByHandP2) Len() int           { return len(b) }
func (b ByHandP2) Less(i, j int) bool { return isLessP2(b[i], b[j]) }
func (b ByHandP2) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

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
	hands := generateHands(input)

	part1(hands)
	part2(hands)
}

func generateHands(input []string) []hand {
	var hands []hand
	for _, handLine := range input {
		handStr := strings.Split(handLine, " ")
		h := hand{}
		h.hand = strings.TrimSpace(handStr[0])
		h.bid, _ = strconv.Atoi(strings.TrimSpace(handStr[1]))
		hands = append(hands, h)
	}

	return hands
}

func getHandRank(h hand, subJoker bool) int {
	handMap := map[string]int{}
	rank := -1
	for _, cardR := range h.hand {
		card := string(cardR)
		_, isInMap := handMap[card]
		if !isInMap {
			handMap[card] = 1
		} else {
			handMap[card]++
		}
	}

	if subJoker && strings.Contains(h.hand, "J") {
		maxLetter := ""
		maxCount := -1
		for k, v := range handMap {
			if v > maxCount && k != "J" {
				maxLetter = k
				maxCount = v
			}
		}

		handMap[maxLetter] += handMap["J"]
		delete(handMap, "J")
	}

	noOfCards := len(handMap)
	if noOfCards == 1 {
		// five of a kind
		return 7
	}

	// 2 cards
	if noOfCards == 2 {
		// 4-1, 3-2
		for _, v := range handMap {
			if v == 1 || v == 4 {
				// 4 of a kind
				return 6
			} else {
				// full house
				return 5
			}
		}
	}

	// 3+ cards
	maxCount := 0
	pairCount := 0
	for _, v := range handMap {
		if v == 2 {
			pairCount++
		}
		if v > maxCount {
			maxCount = v
		}
	}

	switch maxCount {
	case 3:
		// 3 of a kind
		return 4
	case 2:
		if pairCount == 2 {
			// 2 pairs
			return 3
		}
		// 1 pair
		return 2
	case 1:
		return 1
	default:
	}

	return rank
}

var cardRank = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

var cardRankP2 = map[string]int{
	"A": 14, "K": 13, "Q": 12, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1,
}

func isLess(h1, h2 hand) bool {
	rankH1 := getHandRank(h1, false)
	rankH2 := getHandRank(h2, false)

	if rankH1 > rankH2 {
		return false
	} else if rankH2 > rankH1 {
		return true
	}

	// equal ranks
	for i := 0; i < len(h1.hand); i++ {
		if cardRank[string(h1.hand[i])] > cardRank[string(h2.hand[i])] {
			return false
		}

		if cardRank[string(h1.hand[i])] < cardRank[string(h2.hand[i])] {
			return true
		}
	}

	return false
}

func isLessP2(h1, h2 hand) bool {
	rankH1 := getHandRank(h1, true)
	rankH2 := getHandRank(h2, true)

	if rankH1 > rankH2 {
		return false
	} else if rankH2 > rankH1 {
		return true
	}

	// equal ranks
	for i := 0; i < len(h1.hand); i++ {
		if cardRankP2[string(h1.hand[i])] > cardRankP2[string(h2.hand[i])] {
			return false
		}

		if cardRankP2[string(h1.hand[i])] < cardRankP2[string(h2.hand[i])] {
			return true
		}
	}

	return false
}

func part1(hands []hand) {
	sort.Sort(ByHand(hands))

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("Part 1: Winnings: %d\n", winnings)
}

func part2(hands []hand) {
	sort.Sort(ByHandP2(hands))

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	fmt.Printf("Part 2: Winnings: %d\n", winnings)
}
