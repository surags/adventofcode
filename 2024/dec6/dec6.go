package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type set map[string]interface{}
type setDirection map[string]string

func main() {
	//file, err := os.Open("sample.in")
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

	guardMap, startX, startY := parseInput(input)

	noOfDistinctPositions := part1(guardMap, startX, startY)
	fmt.Println("Part 1: Number of Distinct Positions: " + strconv.Itoa(noOfDistinctPositions))

	noOfDistinctPositionsForObstruction := part2(guardMap, startX, startY)
	fmt.Println("Part 2: Number of Distinct Positions For Obstruction: " + strconv.Itoa(noOfDistinctPositionsForObstruction))
}

func parseInput(input []string) ([][]string, int, int) {
	guardMap := [][]string{}
	x, y := 0, 0
	for i, line := range input {
		guardMap = append(guardMap, []string{})
		for j, s := range line {
			guardMap[i] = append(guardMap[i], string(s))
			if string(s) == "^" {
				x, y = i, j
			}
		}
	}

	return guardMap, x, y
}

func positionKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func navigate(guardMap [][]string, x, y int, direction int, distinctPositions set) bool {
	if x < 0 || y < 0 || x >= len(guardMap) || y >= len(guardMap[x]) {
		// done
		return true
	}

	if guardMap[x][y] == "#" {
		// can't move
		return false
	}

	distinctPositions[positionKey(x, y)] = nil

	// directions
	// 1 up, 2 right, 3 down, 4 left
	move := false
	if direction == 1 {
		move = navigate(guardMap, x-1, y, direction, distinctPositions)
	} else if direction == 2 {
		move = navigate(guardMap, x, y+1, direction, distinctPositions)
	} else if direction == 3 {
		move = navigate(guardMap, x+1, y, direction, distinctPositions)
	} else if direction == 4 {
		move = navigate(guardMap, x, y-1, direction, distinctPositions)
	}

	if move {
		return true
	} else {
		direction++
		if direction == 5 {
			direction = 1
		}
		return navigate(guardMap, x, y, direction, distinctPositions)
	}
}

func part1(guardMap [][]string, x, y int) int {
	distinctPositions := set{}

	navigate(guardMap, x, y, 1, distinctPositions)

	return len(distinctPositions)
}

func rotate90(direction int) int {
	direction++
	if direction == 5 {
		direction = 1
	}
	return direction
}

// int 1 reached end, 2 blocked, 3 loop detected
func navigateP2(guardMap [][]string, x, y int, direction int, distinctPositions setDirection, blockX, blockY int, directionChanged bool) int {
	if x < 0 || y < 0 || x >= len(guardMap) || y >= len(guardMap[x]) {
		// done
		return 1
	}

	if guardMap[x][y] == "#" || (x == blockX && y == blockY) {
		// can't move
		return 2
	}

	dir, isLoop := distinctPositions[positionKey(x, y)]
	dirS := strconv.Itoa(direction)
	if isLoop && (dir == dirS) {
		return 3
	} else if isLoop && len(dir) > 1 {
		// check if direction is in dirS
		for _, s := range dir {
			if dirS == string(s) {
				return 3
			}
		}
	} else if isLoop {
		// capture this cross road in set
		distinctPositions[positionKey(x, y)] = distinctPositions[positionKey(x, y)] + strconv.Itoa(direction)
	} else {
		distinctPositions[positionKey(x, y)] = strconv.Itoa(direction)
	}

	// directions
	// 1 up, 2 right, 3 down, 4 left
	move := 0
	if direction == 1 {
		move = navigateP2(guardMap, x-1, y, direction, distinctPositions, blockX, blockY, false)
	} else if direction == 2 {
		move = navigateP2(guardMap, x, y+1, direction, distinctPositions, blockX, blockY, false)
	} else if direction == 3 {
		move = navigateP2(guardMap, x+1, y, direction, distinctPositions, blockX, blockY, false)
	} else if direction == 4 {
		move = navigateP2(guardMap, x, y-1, direction, distinctPositions, blockX, blockY, false)
	}

	if move == 1 {
		return move
	} else if move == 2 {
		// capture all directions in set
		distinctPositions[positionKey(x, y)] = distinctPositions[positionKey(x, y)] + strconv.Itoa(direction)
		direction = rotate90(direction)
		return navigateP2(guardMap, x, y, direction, distinctPositions, blockX, blockY, true)
	} else {
		// move == 3 i.e. loop detected
		return move
	}
}

func keyPosition(key string) (int, int) {
	pos := strings.Split(key, "-")
	posX, _ := strconv.Atoi(pos[0])
	posY, _ := strconv.Atoi(pos[1])

	return posX, posY
}

func part2(guardMap [][]string, x, y int) int {
	noOfPossibleObstructions := 0

	path := set{}
	navigate(guardMap, x, y, 1, path)

	for key, _ := range path {
		blockX, blockY := keyPosition(key)

		distinctPositions := setDirection{}
		// block position i, j
		if x == blockX && y == blockY {
			continue
		}
		res := navigateP2(guardMap, x, y, 1, distinctPositions, blockX, blockY, false)
		if res == 3 {
			noOfPossibleObstructions++
		} else if res == 2 {
			fmt.Println("i returned blocked")
		} else if res == 1 {
			// do nothing
		} else {
			fmt.Println("i don't know what to do anything")
		}
	}

	return noOfPossibleObstructions
}
