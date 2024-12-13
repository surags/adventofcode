package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type loc struct {
	x, y int
}

type machine struct {
	a, b, prize loc
}

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec13.in")

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

	machines := parseInput(input)

	noOfTokens := part1(machines)
	fmt.Println("Part 1: Number of Tokens to spend: " + strconv.Itoa(noOfTokens))

	noOfTokensP2 := part2(machines)
	fmt.Println("Part 2: Number of Corrected Tokens to spend: " + strconv.Itoa(noOfTokensP2))
}

func parseInput(input []string) []machine {
	machines := []machine{}
	m := machine{}
	for _, s := range input {
		ts := strings.Split(s, " ")
		if strings.HasPrefix(s, "Button A") {
			// Button A: X+94, Y+34
			x_s := strings.Trim(ts[2], "X+,")
			x, _ := strconv.Atoi(x_s)
			y_s := strings.Trim(ts[3], "Y+,")
			y, _ := strconv.Atoi(y_s)
			m.a.x = x
			m.a.y = y
		}

		if strings.HasPrefix(s, "Button B") {
			// Button B: X+94, Y+34
			x_s := strings.Trim(ts[2], "X+,")
			x, _ := strconv.Atoi(x_s)
			y_s := strings.Trim(ts[3], "Y+,")
			y, _ := strconv.Atoi(y_s)
			m.b.x = x
			m.b.y = y
		}

		if strings.HasPrefix(s, "Prize:") {
			// Prize: X=7870, Y=6450
			x_s := strings.Trim(ts[1], "X=,")
			x, _ := strconv.Atoi(x_s)
			y_s := strings.Trim(ts[2], "Y=")
			y, _ := strconv.Atoi(y_s)
			m.prize.x = x
			m.prize.y = y
		}

		if s == "" {
			machines = append(machines, m)
			m = machine{}
		}
	}

	machines = append(machines, m)

	return machines
}

func part1(machines []machine) int {
	tokens := 0
	for _, mach := range machines {
		m := float32(-(mach.prize.x*mach.b.y - mach.b.x*mach.prize.y)) / float32(mach.a.y*mach.b.x-mach.a.x*mach.b.y)
		n := float32(mach.a.y*mach.prize.x-mach.a.x*mach.prize.y) / float32(mach.a.y*mach.b.x-mach.a.x*mach.b.y)

		if float32(int32(m)) != m || float32(int32(n)) != n {
			// decimal
			continue
		}

		tokens += 3*int(m) + int(n)
	}
	return tokens
}

func part2(machines []machine) int {
	tokens := 0
	for _, mach := range machines {
		mach.prize.x += 10000000000000
		mach.prize.y += 10000000000000

		m := float64(-(mach.prize.x*mach.b.y - mach.b.x*mach.prize.y)) / float64(mach.a.y*mach.b.x-mach.a.x*mach.b.y)
		n := float64(mach.a.y*mach.prize.x-mach.a.x*mach.prize.y) / float64(mach.a.y*mach.b.x-mach.a.x*mach.b.y)

		if float64(int64(m)) != m || float64(int64(n)) != n {
			// decimal
			continue
		}

		tokens += 3*int(m) + int(n)
	}
	return tokens
}
