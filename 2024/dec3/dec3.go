package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	//file, err := os.Open("dec3.in")
	file, err := os.Open("dec3_2.in") // combine to a single line

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

	sumMultiples := part1(input)
	fmt.Println("Part 1: Sum multiples: " + strconv.Itoa(sumMultiples))

	sumMultiplesWithInstructions := part2(input)
	fmt.Println("Part 2: Sum multiples with instructions: " + strconv.Itoa(sumMultiplesWithInstructions))
}

func part1(input []string) int {
	sum := 0
	for _, s := range input {
		r := regexp.MustCompile("mul\\((?P<Int1>\\d*),(?P<Int2>\\d*)\\).*?")

		matches := r.FindAllStringSubmatch(s, -1)
		for _, match := range matches {
			// match [0] = mul(x, y) match[1] = x match[2] = y
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		}
	}

	return sum
}

func part2(input []string) int {
	sum := 0
	for _, s := range input {
		// maybe someday I'll get this single regex to work
		//r := regexp.MustCompile("((?P<do>do\\(\\)).*?)*(?P<dont>don't\\(\\).*?)*(?P<mul>mul\\((?P<Int1>\\d*),(?P<Int2>\\d*)\\).*?)")

		// Split string into dos and don'ts
		doList := []string{}
		dontList := []string{}

		a := strings.Split(s, "do()")
		for _, s2 := range a {
			// remove don;ts
			b := strings.Split(s2, "don't()")
			doList = append(doList, b[0])
			dontList = append(dontList, b[1:]...)
			fmt.Printf("Len b: %d\n", len(b))
		}

		for _, s3 := range doList {
			r := regexp.MustCompile("mul\\((?P<Int1>\\d*),(?P<Int2>\\d*)\\).*?")
			matches := r.FindAllStringSubmatch(s3, -1)
			for _, match := range matches {
				// match [0] = mul(x, y) match[1] = x match[2] = y
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}
	}

	return sum
}
