package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	pipe = string("|")
	dash = string("-")
	_L   = string("L")
	_J   = string("J")
	_7   = string("7")
	_F   = string("F")
	dot  = string(".")
	_S   = string("S")

	left  = 0
	up    = 1
	down  = 2
	right = 3
)

type dirPipe struct {
	dir  int
	pipe string
}

type node struct {
	i    int
	j    int
	pipe string
	next *node
	prev *node
}

var (
	moveMap = map[dirPipe][]int{
		// left
		dirPipe{dir: left, pipe: _L}:   {-1, 0, up},
		dirPipe{dir: left, pipe: dash}: {0, -1, left},
		dirPipe{dir: left, pipe: _F}:   {1, 0, down},
		// right
		dirPipe{dir: right, pipe: _J}:   {-1, 0, up},
		dirPipe{dir: right, pipe: dash}: {0, 1, right},
		dirPipe{dir: right, pipe: _7}:   {1, 0, down},
		// up
		dirPipe{dir: up, pipe: pipe}: {-1, 0, up},
		dirPipe{dir: up, pipe: _F}:   {0, 1, right},
		dirPipe{dir: up, pipe: _7}:   {0, -1, left},
		// down
		dirPipe{dir: down, pipe: pipe}: {1, 0, down},
		dirPipe{dir: down, pipe: _L}:   {0, 1, right},
		dirPipe{dir: down, pipe: _J}:   {0, -1, left},
	}

	loopLine = map[string]interface{}{}
)

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	file, err := os.Open("sample3.in")
	//file, err := os.Open("sample4.in")
	//file, err := os.Open("dec10.in")

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

	pipeMap, s_i, s_j := buildMap(input)
	//fmt.Printf("%v\n", pipeMap)
	//fmt.Printf("%d, %d", s_i, s_j)"github.com/open-telemetry/opamp-go/client"

	part1(pipeMap, s_i, s_j)
	part2(pipeMap)
}

func buildMap(inputs []string) ([][]string, int, int) {
	length := len(inputs)
	width := len(inputs[0])

	pipeMap := make([][]string, length)
	s_i, s_j := 0, 0

	for i, line := range inputs {
		pipeMap[i] = make([]string, width)
		for j, pipe := range line {
			pipeMap[i][j] = string(pipe)
			if pipeMap[i][j] == _S {
				s_i = i
				s_j = j
			}
		}
	}

	return pipeMap, s_i, s_j
}

func traverseMap(dir int, i int, j int, length int, prevNode *node, pipeMap [][]string) int {
	if i < 0 || j < 0 || i >= len(pipeMap) || j >= len(pipeMap[0]) {
		return -1
	}

	length++

	if pipeMap[i][j] == _S {
		// found
		return length
	}

	curNode := &node{
		i:    i,
		j:    j,
		pipe: pipeMap[i][j],
		next: nil,
		prev: nil,
	}

	coord, movePossible := moveMap[dirPipe{dir: dir, pipe: pipeMap[i][j]}]
	if movePossible {
		newLength := traverseMap(coord[2], i+coord[0], j+coord[1], length, curNode, pipeMap)
		if newLength != -1 {
			prevNode.next = curNode
			loopLine[fmt.Sprintf("%d_%d", i, j)] = nil
			//fmt.Print(pipeMap[i][j])
			return newLength
		}
	}

	return -1
}

func startMapTraversal(sI int, sJ int, startNode *node, pipeMap [][]string) int {
	loopLine[fmt.Sprintf("%d_%d", sI, sJ)] = nil
	cycleLength := traverseMap(left, sI, sJ-1, 0, startNode, pipeMap)
	if cycleLength != -1 {
		return cycleLength
	}

	cycleLength = traverseMap(right, sI, sJ+1, 0, startNode, pipeMap)
	if cycleLength != -1 {
		return cycleLength
	}

	cycleLength = traverseMap(up, sI-1, sJ, 0, startNode, pipeMap)
	if cycleLength != -1 {
		return cycleLength
	}

	cycleLength = traverseMap(down, sI+1, sJ, 0, startNode, pipeMap)
	if cycleLength != -1 {
		return cycleLength
	}

	return -1
}

func part1(pipeMap [][]string, sI int, sJ int) {
	startNode := &node{
		i:    sI,
		j:    sJ,
		pipe: pipeMap[sI][sJ],
		next: nil,
		prev: nil,
	}

	length := startMapTraversal(sI, sJ, startNode, pipeMap)
	fmt.Printf("\nLength: %d\n", length)

	if length%2 == 0 {
		fmt.Printf("Part 1: max distance: %d\n", length/2)
	} else {
		fmt.Printf("Part 1: max distance: %d\n", (length/2)+1)
	}
}

func part2(pipeMap [][]string) {
	// apply scan lines to check if in loop
	// Every time you cross a loopLine switch 0->1->0
	tileCount := 0
	for i, line := range pipeMap {
		isInside := false
		for j, pipe := range line {
			_, isOnLine := loopLine[fmt.Sprintf("%d_%d", i, j)]

			if isOnLine {
				if pipeMap[i][j] == _S || pipeMap[i][j] == _L || pipeMap[i][j] == _J || pipeMap[i][j] == _F || pipeMap[i][j] == _7 || pipeMap[i][j] == pipe {
					isInside = !isInside
				}
				continue
			}

			//if pipeMap[i][j] == dot {
			if isInside {
				tileCount++
			}
			//}
		}
	}

	for j := 0; j < len(pipeMap[0]); j++ {
		isInside := false
		for i := 0; i < len(pipeMap); i++ {

		}

	}

	tileCount
	for i, line := range pipeMap {
		isInside := false
		for j, pipe := range line {
			_, isOnLine := loopLine[fmt.Sprintf("%d_%d", i, j)]

			if isOnLine {
				if pipeMap[i][j] == pipe {

				}
				if pipeMap[i][j] == _S || pipeMap[i][j] == _L || pipeMap[i][j] == _J || pipeMap[i][j] == _F || pipeMap[i][j] == _7 || pipeMap[i][j] == pipe {
					isInside = !isInside
				}
				continue
			}

			//if pipeMap[i][j] == dot {
			if isInside {
				tileCount++
			}
			//}
		}
	}

	fmt.Printf("Part 2: Tiles enclosed: %d\n", tileCount)
}
