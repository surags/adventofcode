package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec1.in")

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

	totalDistance := part1(input)
	fmt.Println("Part 1: Total distance: " + strconv.Itoa(totalDistance))

	similarityScore := part2(input)
	fmt.Println("Part 2: Similarity score: " + strconv.Itoa(similarityScore))
}

// calculate distance
func part1(input []string) int {
	var list1, list2 []int
	for _, v := range input {
		line := strings.Split(v, "   ")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	fmt.Println(list1)
	fmt.Println(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return distance
}

// calculate similarity score
func part2(input []string) int {
	var list1, list2 []int
	for _, v := range input {
		line := strings.Split(v, "   ")
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	list2Count := map[int]int{}
	for _, v := range list2 {
		list2Count[v]++
	}

	similarityScore := 0

	for i := 0; i < len(list1); i++ {
		similarityScore += list1[i] * list2Count[list1[i]]
	}

	return similarityScore
}
