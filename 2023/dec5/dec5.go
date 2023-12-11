package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type rangeMap map[rng]int

type rng struct {
	start int
	end   int
}

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec5.in")

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

	seeds, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap := generateMaps(input)

	printLowestLocNo(seeds, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap)
	printLowestLocNoPart2(seeds, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap)
}

func getMapValue(input int, curMap rangeMap) int {
	for inputRange, diff := range curMap {
		if input >= inputRange.start && input < inputRange.end {
			return input + diff
		}
	}

	return input
}

func updateMap(line string, curMap rangeMap) {
	values := strings.Split(line, " ")
	dest, _ := strconv.Atoi(values[0])
	src, _ := strconv.Atoi(values[1])
	inpRange, _ := strconv.Atoi(values[2])

	diff := dest - src
	key := rng{start: src, end: src + inpRange}
	curMap[key] = diff
}

func generateMaps(inputs []string) ([]int, rangeMap, rangeMap, rangeMap, rangeMap, rangeMap, rangeMap, rangeMap) {
	var seeds []int
	var ssMap = rangeMap{rng{}: 0}
	var sfMap = rangeMap{rng{}: 0}
	var fwMap = rangeMap{rng{}: 0}
	var wlMap = rangeMap{rng{}: 0}
	var ltMap = rangeMap{rng{}: 0}
	var thMap = rangeMap{rng{}: 0}
	var hlMap = rangeMap{rng{}: 0}

	textToMap := map[string]rangeMap{
		"seed-to-soil map:":            ssMap,
		"soil-to-fertilizer map:":      sfMap,
		"fertilizer-to-water map:":     fwMap,
		"water-to-light map:":          wlMap,
		"light-to-temperature map:":    ltMap,
		"temperature-to-humidity map:": thMap,
		"humidity-to-location map:":    hlMap,
	}

	seedsLine := inputs[0]
	seedsNos := strings.Split(strings.TrimSpace(strings.Split(seedsLine, ":")[1]), " ")
	for _, no := range seedsNos {
		seedNo, err := strconv.Atoi(no)
		if err == nil {
			seeds = append(seeds, seedNo)
		}
	}

	//baseSeed := 0
	//for i, no := range seedsNos {
	//	parsedNo, _ := strconv.Atoi(no)
	//	if i%2 == 0 {
	//		baseSeed = parsedNo
	//	} else {
	//		for i := 0; i < parsedNo; i++ {
	//			seedsp2 = append(seedsp2, baseSeed+i)
	//		}
	//	}
	//}

	var currentMap rangeMap = nil
	for _, line := range inputs[1:] {
		if line == "" {
			// reset
			currentMap = nil
			continue
		}

		mapName, isHeader := textToMap[line]
		if isHeader {
			currentMap = mapName
			continue
		}

		updateMap(line, currentMap)
	}

	return seeds, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap
}

func getLocation(seed int, ssMap rangeMap, sfMap rangeMap, fwMap rangeMap, wlMap rangeMap, ltMap rangeMap, thMap rangeMap, hlMap rangeMap) int {
	soil := getMapValue(seed, ssMap)
	fertilizer := getMapValue(soil, sfMap)
	water := getMapValue(fertilizer, fwMap)
	light := getMapValue(water, wlMap)
	temp := getMapValue(light, ltMap)
	humidity := getMapValue(temp, thMap)
	location := getMapValue(humidity, hlMap)

	return location
}

func printLowestLocNo(seeds []int, ssMap rangeMap, sfMap rangeMap, fwMap rangeMap, wlMap rangeMap, ltMap rangeMap, thMap rangeMap, hlMap rangeMap) {
	//fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", seeds, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap)
	minLocationNum := math.MaxInt
	for _, seed := range seeds {
		loc := getLocation(seed, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap)
		if loc < minLocationNum {
			minLocationNum = loc
		}
	}

	fmt.Printf("Part 1: Lowest location number: %d\n", minLocationNum)
}

func printLowestLocNoPart2(seeds []int, ssMap rangeMap, sfMap rangeMap, fwMap rangeMap, wlMap rangeMap, ltMap rangeMap, thMap rangeMap, hlMap rangeMap) {
	baseSeed := 0
	minLocationNum := math.MaxInt
	for i, no := range seeds {
		if i%2 == 0 {
			baseSeed = no
		} else {
			for i := 0; i < no; i++ {
				//seedsp2 = append(seedsp2, baseSeed+i)
				loc := getLocation(baseSeed+i, ssMap, sfMap, fwMap, wlMap, ltMap, thMap, hlMap)
				if loc < minLocationNum {
					minLocationNum = loc
				}
			}
		}
	}

	fmt.Printf("Part 2: Lowest location number: %d\n", minLocationNum)
}
