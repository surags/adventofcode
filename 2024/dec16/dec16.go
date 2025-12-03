package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	file, err := os.Open("dec16.in")

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

	maze, start, end := parseInput(input)

	lowestScore := part1(maze, start, end)
	fmt.Println("Part 1: Lowest Score S->E: " + strconv.Itoa(lowestScore))

	//sumGPSCoords2 := part2(wareHouse2, moves, r2)
	//fmt.Println("Part 2: Sum of all GPS Coordinates: " + strconv.Itoa(sumGPSCoords2))
}

type loc struct {
	x, y int
}

func parseInput(input []string) ([][]string, loc, loc) {
	reindeerMaze := [][]string{}
	start := loc{}
	end := loc{}
	for i, s := range input {
		r := []string{}
		for j, m := range s {
			r = append(r, string(m))
			if string(m) == "S" {
				start.x = i
				start.y = j
			}

			if string(m) == "E" {
				end.x = i
				end.y = j
			}
		}
		reindeerMaze = append(reindeerMaze, r)
	}

	return reindeerMaze, start, end
}

func getSetKey(cur loc, dir int) string {
	return fmt.Sprintf("%d-%d-%d", cur.x, cur.y, dir)
}

func search(maze [][]string, start, end loc) int {
	cur := loc{start.x, start.y}
	visitedSet := map[string]interface{}{}
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	heap.Push(&q, &op{cur: cur, direction: 2, distance: 0})

	for {
		h := heap.Pop(&q)
		if h == nil {
			break
		}
		o := h.(*op)
		fmt.Println(o)
		if o.cur.x < 0 || o.cur.y < 0 || o.cur.x >= len(maze) || o.cur.y >= len(maze[0]) {
			// out of bounds...shouldn't happen
			continue
		}
		if _, f := visitedSet[getSetKey(o.cur, o.direction)]; f {
			continue
		} else {
			visitedSet[getSetKey(o.cur, o.direction)] = struct{}{}
		}
		if maze[o.cur.x][o.cur.y] == "E" {
			return o.distance
		}
		if maze[o.cur.x][o.cur.y] == "#" {
			// can't move here
			continue
		}

		// add moves
		if o.direction == 1 {
			// move N
			heap.Push(&q, &op{cur: loc{o.cur.x - 1, o.cur.y}, direction: o.direction, distance: o.distance + 1})
		} else if o.direction == 2 {
			// move E
			heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y + 1}, direction: o.direction, distance: o.distance + 1})
		} else if o.direction == 2 {
			// move S
			heap.Push(&q, &op{cur: loc{o.cur.x + 1, o.cur.y}, direction: o.direction, distance: o.distance + 1})
		} else {
			heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y - 1}, direction: o.direction, distance: o.distance + 1})
		}

		newDir := rotate90(o.direction)
		heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y}, direction: newDir, distance: o.distance + 1000})
		newDir = rotate90(newDir)
		heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y}, direction: newDir, distance: o.distance + 1000 + 1000})
		// can technically rotate counter-clockwise so only 1000 extra score
		newDir = rotate90(newDir)
		heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y}, direction: newDir, distance: o.distance + 1000})
	}

	return -1
}

func rotate90(dir int) int {
	newDir := dir + 1
	if newDir == 5 {
		newDir = 1
	}
	return newDir
}

func part1(maze [][]string, start, end loc) int {
	bestScore := search(maze, start, end)

	return bestScore
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*op

type op struct {
	cur       loc
	direction int // 1 N, 2 E, 3 S, 4 W
	distance  int
	i         int // internal
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	if i == -1 || j == -1 {
		// notjing to swap
		return
	}
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].i = i
	pq[j].i = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*op)
	item.i = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	item.i = -1    // for safety
	*pq = old[0 : n-1]
	return item
}
