package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type loc struct {
	x, y int
}

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec18.in")

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
	//memSize := 7
	memSize := 71
	mem, corruptedBytes, start, end := parseInput(input, memSize)

	lowestScore := part1(mem, corruptedBytes, start, end, 1024)
	//lowestScore := part1(mem, corruptedBytes, start, end, 12)
	fmt.Println("Part 1: Min No Of Steps: " + strconv.Itoa(lowestScore))

	l := part2(corruptedBytes, start, end, memSize)
	fmt.Printf("Part 2: Location of first corrupted byte making the end unreachable: %d,%d\n", l.x, l.y)
}

func genMem(memSize int) [][]string {
	// mem size memSize x memSize
	mem := make([][]string, memSize)
	for i, _ := range mem {
		mem[i] = []string{}
		for j := 0; j < memSize; j++ {
			mem[i] = append(mem[i], ".")
		}
	}
	return mem
}

func parseInput(input []string, memSize int) ([][]string, []loc, loc, loc) {
	mem := genMem(memSize)
	start := loc{0, 0}
	end := loc{memSize - 1, memSize - 1}

	corruptedBytes := []loc{}
	for _, coordS := range input {
		cb := loc{}
		coord := strings.Split(coordS, ",")
		cb.x, _ = strconv.Atoi(coord[0])
		cb.y, _ = strconv.Atoi(coord[1])
		corruptedBytes = append(corruptedBytes, cb)
	}

	return mem, corruptedBytes, start, end
}

func getSetKey(cur loc) string {
	return fmt.Sprintf("%d-%d", cur.x, cur.y)
}

func getSetLoc(key string) loc {
	s := strings.Split(key, "-")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	return loc{x, y}
}

func search(mem [][]string, start, end loc) int {
	cur := loc{start.x, start.y}
	visitedSet := map[string]interface{}{}
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	heap.Push(&q, &op{cur: cur, distance: 0})

	for {
		h := heap.Pop(&q)
		if h == nil {
			break
		}
		o := h.(*op)
		//fmt.Println(o)
		if o.cur.x < 0 || o.cur.y < 0 || o.cur.x >= len(mem) || o.cur.y >= len(mem[0]) {
			continue
		}

		//// print mem with visited
		//fmt.Print("\033[0;0H")
		//for i, i2 := range mem {
		//	for i3, s := range i2 {
		//		if i == o.cur.y && o.cur.x == i3 {
		//			fmt.Print("@")
		//			continue
		//		}
		//		if _, f := visitedSet[s]; f {
		//			if mem[i][i3] == "." {
		//				fmt.Print("O")
		//			} else {
		//				fmt.Print(s)
		//			}
		//		} else {
		//			fmt.Print(s)
		//		}
		//	}
		//	fmt.Print("\n")
		//}
		//time.Sleep(100 * time.Millisecond)

		if o.cur.x == end.x && o.cur.y == end.y {
			// destination reached
			return o.distance
		}
		if _, f := visitedSet[getSetKey(o.cur)]; f {
			continue
		} else {
			visitedSet[getSetKey(o.cur)] = struct{}{}
		}

		if mem[o.cur.y][o.cur.x] == "#" {
			// can't move here
			continue
		}

		// add moves
		heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y - 1}, distance: o.distance + 1}) // up
		heap.Push(&q, &op{cur: loc{o.cur.x, o.cur.y + 1}, distance: o.distance + 1}) // down
		heap.Push(&q, &op{cur: loc{o.cur.x - 1, o.cur.y}, distance: o.distance + 1}) // left
		heap.Push(&q, &op{cur: loc{o.cur.x + 1, o.cur.y}, distance: o.distance + 1}) // right
	}

	// print mem with visited
	//for i, i2 := range mem {
	//	for i3, s := range i2 {
	//		if _, f := visitedSet[getSetKey(loc{i3, i})]; f {
	//			if mem[i][i3] == "." {
	//				fmt.Print("O")
	//			} else {
	//				fmt.Print(s)
	//			}
	//		} else {
	//			fmt.Print(s)
	//		}
	//	}
	//	fmt.Print("\n")
	//}

	return -1
}

func part1(mem [][]string, corruptedBytes []loc, start, end loc, noOfBytes int) int {
	// apply first n corrupted bytes
	for i := 0; i < noOfBytes; i++ {
		mem[corruptedBytes[i].y][corruptedBytes[i].x] = "#"
	}

	return search(mem, start, end)
}

func part2(corruptedBytes []loc, start, end loc, memSize int) loc {
	for i := 0; i < len(corruptedBytes); i++ {
		mem := genMem(memSize)
		dist := part1(mem, corruptedBytes, start, end, i)

		if dist == -1 {
			// unreachable
			return loc{corruptedBytes[i-1].x, corruptedBytes[i-1].y}
		}
	}

	return loc{-1, -1}
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*op

type op struct {
	cur      loc
	distance int
	i        int // internal
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
