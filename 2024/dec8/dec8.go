package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type loc struct {
	x, y int
}

type antenna struct {
	name     string
	location loc
}

type set map[loc]bool

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	file, err := os.Open("dec8.in")

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

	antennaMap, antennaList := parseInput(input)

	noOfUniqueAntinodeLocations := part1(antennaMap, antennaList)
	fmt.Println("Part 1: Number of Unique Antinode Positions: " + strconv.Itoa(noOfUniqueAntinodeLocations))

	noOfUniqueAntinodeNoDistLocations := part2(antennaMap, antennaList)
	fmt.Println("Part 2: Number of Unique Antinode Positions Without Distance restriction: " + strconv.Itoa(noOfUniqueAntinodeNoDistLocations))
}

func parseInput(input []string) ([][]string, map[string][]antenna) {
	antennaMap := [][]string{}
	antennaList := map[string][]antenna{}
	for i, line := range input {
		antennaMap = append(antennaMap, []string{})
		for j, s := range line {
			antennaMap[i] = append(antennaMap[i], string(s))

			if string(s) != "." {
				a := antenna{
					name:     string(s),
					location: loc{i, j},
				}

				if _, ok := antennaList[a.name]; !ok {
					antennaList[a.name] = []antenna{}
				}
				antennaList[a.name] = append(antennaList[a.name], a)
			}
		}
	}

	return antennaMap, antennaList
}

func generatePairs(locs []antenna) [][]antenna {
	antennaPairs := [][]antenna{}

	for i, loc := range locs {
		for j, loc2 := range locs {
			if j <= i {
				continue
			}
			pair := []antenna{loc, loc2}
			antennaPairs = append(antennaPairs, pair)
		}
	}

	return antennaPairs
}

func getInLinePoints(antennaMap [][]string, pair []antenna) []loc {
	distLoc := loc{
		x: pair[0].location.x - pair[1].location.x,
		y: pair[0].location.y - pair[1].location.y,
	}

	antiNodes := []loc{}

	// get 1 before and 1 after
	beforePoint := loc{}
	beforePoint.x = pair[1].location.x - distLoc.x
	beforePoint.y = pair[1].location.y - distLoc.y

	if !(beforePoint.x < 0 || beforePoint.y < 0 || beforePoint.x >= len(antennaMap) || beforePoint.y >= len(antennaMap[0])) {
		antiNodes = append(antiNodes, beforePoint)
	}

	afterPoint := loc{}
	afterPoint.x = pair[0].location.x + distLoc.x
	afterPoint.y = pair[0].location.y + distLoc.y

	if !(afterPoint.x < 0 || afterPoint.y < 0 || afterPoint.x >= len(antennaMap) || afterPoint.y >= len(antennaMap[0])) {
		antiNodes = append(antiNodes, afterPoint)
	}

	return antiNodes
}

func findAntiNode(antennaMap [][]string, pair []antenna, foundNodes set) {
	nodes := getInLinePoints(antennaMap, pair)
	for _, node := range nodes {
		foundNodes[node] = true
	}
}

func part1(antennaMap [][]string, antennaList map[string][]antenna) int {
	antiNodeSet := set{}
	for _, antennas := range antennaList {
		pairs := generatePairs(antennas)
		for _, pair := range pairs {
			findAntiNode(antennaMap, pair, antiNodeSet)
		}
	}

	return len(antiNodeSet)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return int(math.Abs(float64(a)))
}

func getAllInLinePoints(antennaMap [][]string, pair []antenna) []loc {
	distLoc := loc{
		x: pair[0].location.x - pair[1].location.x,
		y: pair[0].location.y - pair[1].location.y,
	}

	denom := gcd(distLoc.x, distLoc.y)

	distLoc.x /= denom
	distLoc.y /= denom

	antiNodes := []loc{pair[0].location, pair[1].location}

	// get all before and all after
	mult := 1
	for {
		beforePoint := loc{}
		beforePoint.x = pair[1].location.x - mult*distLoc.x
		beforePoint.y = pair[1].location.y - mult*distLoc.y

		if !(beforePoint.x < 0 || beforePoint.y < 0 || beforePoint.x >= len(antennaMap) || beforePoint.y >= len(antennaMap[0])) {
			antiNodes = append(antiNodes, beforePoint)
			mult++
		} else {
			break
		}
	}

	mult = 1
	for {
		afterPoint := loc{}
		afterPoint.x = pair[0].location.x + mult*distLoc.x
		afterPoint.y = pair[0].location.y + mult*distLoc.y

		if !(afterPoint.x < 0 || afterPoint.y < 0 || afterPoint.x >= len(antennaMap) || afterPoint.y >= len(antennaMap[0])) {
			antiNodes = append(antiNodes, afterPoint)
			mult++
		} else {
			break
		}
	}

	return antiNodes
}

func findAntiNodeP2(antennaMap [][]string, pair []antenna, foundNodes set) {
	nodes := getAllInLinePoints(antennaMap, pair)
	for _, node := range nodes {
		foundNodes[node] = true
	}
}

func part2(antennaMap [][]string, antennaList map[string][]antenna) int {
	antiNodeSet := set{}
	for _, antennas := range antennaList {
		pairs := generatePairs(antennas)
		for _, pair := range pairs {
			findAntiNodeP2(antennaMap, pair, antiNodeSet)
		}
	}

	return len(antiNodeSet)
}
