package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type loc struct {
	x, y int
}

type set map[loc]interface{}

func (s set) contains(loc loc) bool {
	_, ok := s[loc]
	return ok
}

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	//file, err := os.Open("sample3.in")
	file, err := os.Open("dec12.in")

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

	arrangement := parseInput(input)

	totalPriceOfFencing := part1(arrangement)
	fmt.Println("Part 1: Total Price of Fencing: " + strconv.Itoa(totalPriceOfFencing))

	totalBulkPricingOfFencing := part2(arrangement)
	fmt.Println("Part 2: Total Bulk Pricing of Fencing: " + strconv.Itoa(totalBulkPricingOfFencing))
}

func parseInput(input []string) [][]string {
	arrangement := [][]string{}
	for _, line := range input {
		arrangementLine := []string{}
		for _, plots := range line {
			arrangementLine = append(arrangementLine, string(plots))
		}
		arrangement = append(arrangement, arrangementLine)
	}

	return arrangement
}

func processRegion(arrangement [][]string, region set, regionStart map[loc]loc, plantType string, curL, prevL, startL loc) {
	if curL.x < 0 || curL.x >= len(arrangement) || curL.y < 0 || curL.y >= len(arrangement[0]) {
		// out of bounds
		return
	}

	if region.contains(curL) {
		// already found
		return
	}

	if curL == prevL {
		return
	}

	if plantType != arrangement[curL.x][curL.y] {
		// don't proceed
		return
	}

	// add to region
	region[curL] = nil
	regionStart[curL] = startL

	processRegion(arrangement, region, regionStart, plantType, loc{x: curL.x - 1, y: curL.y}, curL, startL) // up
	processRegion(arrangement, region, regionStart, plantType, loc{x: curL.x + 1, y: curL.y}, curL, startL) // down
	processRegion(arrangement, region, regionStart, plantType, loc{x: curL.x, y: curL.y - 1}, curL, startL) // left
	processRegion(arrangement, region, regionStart, plantType, loc{x: curL.x, y: curL.y + 1}, curL, startL) // right
}

func buildRegions(arrangement [][]string, regionMap map[loc]set, regionStart map[loc]loc) {
	for i, plotLines := range arrangement {
		for j, plot := range plotLines {
			if _, found := regionStart[loc{x: i, y: j}]; found {
				// work already done
				continue
			}

			l := loc{x: i, y: j}
			regionStart[l] = l
			region := set{l: nil}
			regionMap[l] = region
			processRegion(arrangement, region, regionStart, plot, loc{x: l.x - 1, y: l.y}, l, l) // up
			processRegion(arrangement, region, regionStart, plot, loc{x: l.x + 1, y: l.y}, l, l) // down
			processRegion(arrangement, region, regionStart, plot, loc{x: l.x, y: l.y - 1}, l, l) // left
			processRegion(arrangement, region, regionStart, plot, loc{x: l.x, y: l.y + 1}, l, l) // right
		}
	}
}

func perimeter(arrangement [][]string, region set, l loc, visited set) int {
	if !region.contains(l) {
		// it's a wall here
		return 1
	}

	if l.x < 0 || l.x >= len(arrangement) || l.y < 0 || l.y >= len(arrangement[0]) {
		// out of bounds wall
		return 1
	}

	if visited.contains(l) {
		return 0
	}

	visited[l] = struct{}{}

	return perimeter(arrangement, region, loc{x: l.x - 1, y: l.y}, visited) + // up
		perimeter(arrangement, region, loc{x: l.x + 1, y: l.y}, visited) + // down
		perimeter(arrangement, region, loc{x: l.x, y: l.y - 1}, visited) + // left
		perimeter(arrangement, region, loc{x: l.x, y: l.y + 1}, visited) // right
}

func calculatePriceOfFencing(arrangement [][]string, regionMap map[loc]set) int {
	priceOfFencing := 0
	for l, region := range regionMap {
		area := len(region)
		v := set{l: struct{}{}}
		p := perimeter(arrangement, region, loc{x: l.x - 1, y: l.y}, v) + // up
			perimeter(arrangement, region, loc{x: l.x + 1, y: l.y}, v) + // down
			perimeter(arrangement, region, loc{x: l.x, y: l.y - 1}, v) + // left
			perimeter(arrangement, region, loc{x: l.x, y: l.y + 1}, v) // right

		priceOfFencing += p * area
	}
	return priceOfFencing
}

func part1(arrangement [][]string) int {
	regionMap := map[loc]set{}
	regionStart := map[loc]loc{}

	buildRegions(arrangement, regionMap, regionStart)
	return calculatePriceOfFencing(arrangement, regionMap)
}

type wall struct {
	a, b loc
}

type wallSet map[wall]struct{}

func buildWalls(arrangement [][]string, region set, plantType string, l loc, prevL loc, visited set, walls wallSet) int {
	if !region.contains(l) {
		// it's a wall here
		wall := wall{a: prevL, b: l}
		walls[wall] = struct{}{}
		return 1
	}

	if l.x < 0 || l.x >= len(arrangement) || l.y < 0 || l.y >= len(arrangement[0]) {
		// out of bounds wall
		wall := wall{a: prevL, b: l}
		walls[wall] = struct{}{}
		return 1
	}

	if visited.contains(l) {
		return 0
	}

	visited[l] = struct{}{}

	if arrangement[l.x][l.y] != plantType {
		// wall
		wall := wall{a: prevL, b: l}
		walls[wall] = struct{}{}
		return 1
	}

	return buildWalls(arrangement, region, plantType, loc{x: l.x - 1, y: l.y}, l, visited, walls) + // up
		buildWalls(arrangement, region, plantType, loc{x: l.x + 1, y: l.y}, l, visited, walls) + // down
		buildWalls(arrangement, region, plantType, loc{x: l.x, y: l.y - 1}, l, visited, walls) + // left
		buildWalls(arrangement, region, plantType, loc{x: l.x, y: l.y + 1}, l, visited, walls) // right
}

func noOfCollapsibleWalls(walls wallSet) int {
	removeWalls := 0
	for w, _ := range walls {
		if w.b.x-w.a.x == 1 {
			// up wall
			// check right and collapse
			possibleConnectedWall := wall{a: loc{w.a.x, w.a.y + 1}, b: loc{w.b.x, w.b.y + 1}}
			if _, found := walls[possibleConnectedWall]; found {
				removeWalls++
			}
		}

		if w.a.x-w.b.x == 1 {
			// down wall
			// check right and collapse
			possibleConnectedWall := wall{a: loc{w.a.x, w.a.y + 1}, b: loc{w.b.x, w.b.y + 1}}
			if _, found := walls[possibleConnectedWall]; found {
				removeWalls++
			}
		}

		if w.b.y-w.a.y == 1 {
			// right wall
			// check down and collapse
			possibleConnectedWall := wall{a: loc{w.a.x + 1, w.a.y}, b: loc{w.b.x + 1, w.b.y}}
			if _, found := walls[possibleConnectedWall]; found {
				removeWalls++
			}
		}

		if w.a.y-w.b.y == 1 {
			// left wall
			// check down and collapse
			possibleConnectedWall := wall{a: loc{w.a.x + 1, w.a.y}, b: loc{w.b.x + 1, w.b.y}}
			if _, found := walls[possibleConnectedWall]; found {
				removeWalls++
			}
		}
	}
	return removeWalls
}

func calculateBulkPricingOfFencing(arrangement [][]string, regionMap map[loc]set) int {
	priceOfFencing := 0
	for l, region := range regionMap {
		plantType := arrangement[l.x][l.y]
		area := len(region)
		v := set{l: struct{}{}}
		walls := wallSet{}
		p := buildWalls(arrangement, region, plantType, loc{x: l.x - 1, y: l.y}, l, v, walls) + // up
			buildWalls(arrangement, region, plantType, loc{x: l.x + 1, y: l.y}, l, v, walls) + // down
			buildWalls(arrangement, region, plantType, loc{x: l.x, y: l.y - 1}, l, v, walls) + // left
			buildWalls(arrangement, region, plantType, loc{x: l.x, y: l.y + 1}, l, v, walls) // right

		subWalls := noOfCollapsibleWalls(walls)

		//fmt.Printf("P: %d Walls: %d collapsible: %d \n", p, len(walls), subWalls)

		priceOfFencing += (p - subWalls) * area
	}
	return priceOfFencing
}

func part2(arrangement [][]string) int {
	regionMap := map[loc]set{}
	regionStart := map[loc]loc{}

	buildRegions(arrangement, regionMap, regionStart)
	return calculateBulkPricingOfFencing(arrangement, regionMap)
}
