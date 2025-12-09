package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("sample.in")
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

	pdtOfSizesOf3LargestCircuits := part1(input)
	fmt.Println("Part 1: Product of sizes of 3 largest circuits: " + strconv.Itoa(pdtOfSizesOf3LargestCircuits))

	pdtOfXCoordsForLast2Boxes := part2(input)
	fmt.Println("Part 2: Product of X coordinates for last 2 boxes: " + strconv.Itoa(pdtOfXCoordsForLast2Boxes))
}

func distance(a, b junctionBox) float64 {
	return math.Sqrt(math.Pow(float64(a.x)-float64(b.x), 2) + math.Pow(float64(a.y)-float64(b.y), 2) + math.Pow(float64(a.z)-float64(b.z), 2))
}

type junctionBox struct {
	x       int
	y       int
	z       int
	id      int
	circuit int
}

type junctionBoxDistance struct {
	a *junctionBox
	b *junctionBox
	d float64
}

func createNConnections(n int, distances []junctionBoxDistance, circuits map[int][]int, junctionBoxMap map[int]*junctionBox) {
	i := 0
	connections := 0
	for i < n {
		v := distances[i]
		circuitA := v.a.circuit
		circuitB := v.b.circuit
		if circuitA == circuitB {
			i++
			continue
		}
		i++
		connections++

		fmt.Printf("Connecting junction boxes: %d, %d from circuits %d and %d\n", v.a.id, v.b.id, circuitA, circuitB)

		// connect boxes
		if len(circuits[circuitA]) > len(circuits[circuitB]) {
			// add circuitB to circuitA
			for _, boxID := range circuits[circuitB] {
				junctionBoxMap[boxID].circuit = circuitA
				circuits[circuitA] = append(circuits[circuitA], boxID)
			}
			delete(circuits, circuitB)
		} else {
			// add circuitA to circuitB
			for _, boxID := range circuits[circuitA] {
				junctionBoxMap[boxID].circuit = circuitB
				circuits[circuitB] = append(circuits[circuitB], boxID)
			}
			delete(circuits, circuitA)
		}
	}
}

func part1(input []string) int {
	junctionBoxes := []junctionBox{}
	junctionBoxMap := make(map[int]*junctionBox)
	for i, v := range input {
		coordsString := strings.Split(v, ",")
		x, _ := strconv.Atoi(coordsString[0])
		y, _ := strconv.Atoi(coordsString[1])
		z, _ := strconv.Atoi(coordsString[2])
		junctionBoxes = append(junctionBoxes, junctionBox{x: x, y: y, z: z, id: i})
		junctionBoxMap[i] = &junctionBoxes[i]
	}

	distances := []junctionBoxDistance{}
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			d := distance(junctionBoxes[i], junctionBoxes[j])
			distances = append(distances, junctionBoxDistance{a: junctionBoxMap[i], b: junctionBoxMap[j], d: d})
		}
	}

	// Sort distances
	slices.SortFunc(distances, func(a, b junctionBoxDistance) int {
		return int(a.d - b.d)
	})

	// Create n circuits
	circuits := map[int][]int{}
	for i, v := range junctionBoxMap {
		circuits[i] = []int{v.id}
		v.circuit = i
	}

	createNConnections(1000, distances, circuits, junctionBoxMap)
	fmt.Println("No of Circuits: ", len(circuits))

	// Sort circuits by size
	var circuitsSizes [][]int
	for i, v := range circuits {
		circuitsSizes = append(circuitsSizes, []int{i, len(v)})
	}
	slices.SortFunc(circuitsSizes, func(a, b []int) int {
		return b[1] - a[1]
	})

	fmt.Printf("Largest Circuits: %d: %d, %d: %d, %d: %d\n", circuitsSizes[0][0], circuitsSizes[0][1], circuitsSizes[1][0], circuitsSizes[1][1], circuitsSizes[2][0], circuitsSizes[2][1])

	// Distance between all junction boxes
	return circuitsSizes[0][1] * circuitsSizes[1][1] * circuitsSizes[2][1]
}

func connectTillSingleCircuit(distances []junctionBoxDistance, circuits map[int][]int, junctionBoxMap map[int]*junctionBox) (junctionBoxDistance, error) {
	i := 0
	connections := 0
	for i < len(distances) {
		v := distances[i]
		circuitA := v.a.circuit
		circuitB := v.b.circuit
		if circuitA == circuitB {
			i++
			continue
		}
		i++
		connections++

		fmt.Printf("Connecting junction boxes: %d, %d from circuits %d and %d\n", v.a.id, v.b.id, circuitA, circuitB)

		// connect boxes
		if len(circuits[circuitA]) > len(circuits[circuitB]) {
			// add circuitB to circuitA
			for _, boxID := range circuits[circuitB] {
				junctionBoxMap[boxID].circuit = circuitA
				circuits[circuitA] = append(circuits[circuitA], boxID)
			}
			delete(circuits, circuitB)
		} else {
			// add circuitA to circuitB
			for _, boxID := range circuits[circuitA] {
				junctionBoxMap[boxID].circuit = circuitB
				circuits[circuitB] = append(circuits[circuitB], boxID)
			}
			delete(circuits, circuitA)
		}
		if len(circuits) == 1 {
			return v, nil
		}
	}
	return junctionBoxDistance{}, fmt.Errorf("Could not connect all junction boxes into a single circuit")
}

func part2(input []string) int {
	junctionBoxes := []junctionBox{}
	junctionBoxMap := make(map[int]*junctionBox)
	for i, v := range input {
		coordsString := strings.Split(v, ",")
		x, _ := strconv.Atoi(coordsString[0])
		y, _ := strconv.Atoi(coordsString[1])
		z, _ := strconv.Atoi(coordsString[2])
		junctionBoxes = append(junctionBoxes, junctionBox{x: x, y: y, z: z, id: i})
		junctionBoxMap[i] = &junctionBoxes[i]
	}

	distances := []junctionBoxDistance{}
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			d := distance(junctionBoxes[i], junctionBoxes[j])
			distances = append(distances, junctionBoxDistance{a: junctionBoxMap[i], b: junctionBoxMap[j], d: d})
		}
	}

	// Sort distances
	slices.SortFunc(distances, func(a, b junctionBoxDistance) int {
		return int(a.d - b.d)
	})

	// Create n circuits
	circuits := map[int][]int{}
	for i, v := range junctionBoxMap {
		circuits[i] = []int{v.id}
		v.circuit = i
	}

	lastConnection, err := connectTillSingleCircuit(distances, circuits, junctionBoxMap)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return lastConnection.a.x * lastConnection.b.x
}
