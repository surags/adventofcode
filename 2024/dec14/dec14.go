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
	//file, err := os.Open("sample.in")
	//mapSizeX, mapSizeY := 11, 7
	file, err := os.Open("dec14.in")
	mapSizeX, mapSizeY := 101, 103

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

	bathroomMap, robots := parseInput(input, mapSizeX, mapSizeY)

	safetyFactor := part1(robots, bathroomMap)
	fmt.Println("Part 1: Robot Safety Factor: " + strconv.Itoa(safetyFactor))

	part2(robots, bathroomMap)
}

type robot struct {
	posX, posY int
	velX, velY int
}

func parseInput(input []string, mapSizeX int, mapSizeY int) ([][]string, []robot) {
	bathroomMap := make([][]string, mapSizeY)
	for i := range bathroomMap {
		bathroomMap[i] = make([]string, mapSizeX)
	}

	robots := []robot{}
	for _, line := range input {
		r := robot{}
		s := strings.Split(line, " ")

		p := strings.Split(strings.Split(s[0], "=")[1], ",")
		v := strings.Split(strings.Split(s[1], "=")[1], ",")

		r.posX, _ = strconv.Atoi(p[0])
		r.posY, _ = strconv.Atoi(p[1])

		r.velX, _ = strconv.Atoi(v[0])
		r.velY, _ = strconv.Atoi(v[1])

		robots = append(robots, r)
	}

	return bathroomMap, robots
}

func moveRobot(r robot, seconds int, bathroomMap [][]string) robot {
	for i := 0; i < seconds; i++ {
		r.posX += r.velX
		r.posY += r.velY

		if r.posX < 0 {
			r.posX = len(bathroomMap[0]) + r.posX
		}

		if r.posY < 0 {
			r.posY = len(bathroomMap) + r.posY
		}

		if r.posX >= len(bathroomMap[0]) {
			r.posX = r.posX - len(bathroomMap[0])
		}

		if r.posY >= len(bathroomMap) {
			r.posY = r.posY - len(bathroomMap)
		}
	}
	return r
}

// 0 top left, 1 top right, 2 bottom left 3 bottom right
func getQuadrant(r robot, bathroomMap [][]string) int {
	x_mid, y_mid := len(bathroomMap[0])/2, len(bathroomMap)/2
	if r.posX < x_mid && r.posY < y_mid {
		return 0
	}

	if r.posX > x_mid && r.posY < y_mid {
		return 1
	}

	if r.posX < x_mid && r.posY > y_mid {
		return 2
	}

	if r.posX > x_mid && r.posY > y_mid {
		return 3
	}

	return -1
}

func part1(robots []robot, bathroomMap [][]string) int {
	quadrants := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}

	res := []robot{}
	for _, r := range robots {
		r_res := moveRobot(r, 100, bathroomMap)
		res = append(res, r_res)
	}

	for _, re := range res {
		q := getQuadrant(re, bathroomMap)
		quadrants[q]++
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func updateMap(r robot, bathroomMap [][]string) {
	bathroomMap[r.posY][r.posX] = "o"
}

func part2(robots []robot, bathroomMap [][]string) int {
	secondTime := -1
	// seemed to be closest at a diff of 101
	for i := 538; i < 10000; i += 101 {
		res := []robot{}
		for _, r := range robots {
			r_res := moveRobot(r, i, bathroomMap)
			res = append(res, r_res)
		}

		iterMap := make([][]string, len(bathroomMap))
		for i := range iterMap {
			iterMap[i] = make([]string, len(bathroomMap[0]))
			for j := 0; j < len(iterMap[i]); j++ {
				iterMap[i][j] = "_"
			}
		}
		for _, re := range res {
			updateMap(re, iterMap)
		}

		fmt.Printf("Iteration: %d\n\n", i)
		for i := range iterMap {
			for j := 0; j < len(iterMap[i]); j++ {
				fmt.Printf("%s", iterMap[i][j])
			}
			fmt.Printf("\n")
		}
	}

	return secondTime
}
