package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec3.in")

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

	schematic := buildSchematic(input)
	//fmt.Printf("Schematic: \n %v \n", schematic)

	part1(schematic)
	part2(schematic)
}

func buildSchematic(inputs []string) []string {
	schematic := make([]string, len(inputs))

	for index, input := range inputs {
		schematic[index] = input
	}

	return schematic
}

var digits = map[string]interface{}{"0": nil, "1": nil, "2": nil, "3": nil, "4": nil, "5": nil, "6": nil, "7": nil, "8": nil, "9": nil}

var partLocMap = map[string]*int{}

func isPartNumber(i, j, leng int, buildNum string, schematic []string) (int, bool) {
	conv, err := strconv.Atoi(buildNum)
	if err != nil {
		return -1, false
	}

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+leng; y++ {
			if x < 0 || y < 0 {
				continue
			}

			if x >= len(schematic) || y >= len(schematic[0]) {
				continue
			}

			ch := string(schematic[x][y])
			_, isDigit := digits[ch]
			if !isDigit && ch != "." {
				return conv, true
			}
		}
	}

	return conv, false
}

func part1(schematic []string) {
	partNums := []int{}
	partSum := 0

	for i, line := range schematic {
		buildNumber := ""
		leng := 0
		j := 0
		for ind, ch := range line {
			strCh := string(ch)
			_, ok := digits[strCh]
			if ok {
				buildNumber = buildNumber + strCh
				leng++
			} else {
				num, res := isPartNumber(i, j, leng, buildNumber, schematic)
				if res {
					partNums = append(partNums, num)
					partSum += num
					for n := j; n < j+leng; n++ {
						partLocMap[fmt.Sprintf("%d_%d", i, n)] = &num
					}
				}

				buildNumber = ""
				j = ind + 1
				leng = 0
			}
		}

		if buildNumber != "" {
			num, res := isPartNumber(i, j, leng, buildNumber, schematic)
			if res {
				partNums = append(partNums, num)
				partSum += num
				for n := j; n < j+leng; n++ {
					partLocMap[fmt.Sprintf("%d_%d", i, n)] = &num
				}
			}
		}
	}

	//fmt.Printf("Part nums: %v \n", partNums)
	fmt.Printf("Part 2: Part sum: %v \n", partSum)
}

func isGear(i, j int, schematic []string) (int, bool) {
	parts := map[*int]interface{}{}
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x < 0 || y < 0 {
				continue
			}

			if x >= len(schematic) || y >= len(schematic[0]) {
				continue
			}

			potentialPart, isPart := partLocMap[fmt.Sprintf("%d_%d", x, y)]
			if isPart {
				parts[potentialPart] = nil
			}
		}
	}

	if len(parts) == 2 {
		gearRatio := 1
		for k, _ := range parts {
			gearRatio *= *k
		}
		return gearRatio, true
	}

	return -1, false
}

func part2(schematic []string) {
	gearSum := 0

	for i, line := range schematic {
		for j, ch := range line {
			strCh := string(ch)
			if strCh == "*" {
				gearRatio, res := isGear(i, j, schematic)
				if res {
					gearSum += gearRatio
				}
			}
		}
	}

	fmt.Printf("Part2: Gear Ratio sum: %v \n", gearSum)
}
