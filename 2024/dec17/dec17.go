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

type registers struct {
	A, B, C int
}

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec17.in")

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

	r, p := parseInput(input)
	//fmt.Println(p)
	//fmt.Println(r)

	out := part1(r, p)
	fmt.Println("Part 1: Program Output: " + out)

	regA := part2(r, p)
	fmt.Println("Part 2: Value of Register A: " + strconv.Itoa(regA))
}

func parseInput(input []string) (registers, []int) {
	r := registers{}
	p := []int{}

	for _, s := range input {
		if strings.Contains(s, "Register A") {
			v := strings.Split(s, ":")
			r.A, _ = strconv.Atoi(strings.Trim(v[1], " "))
		}
		if strings.Contains(s, "Register B") {
			v := strings.Split(s, ":")
			r.B, _ = strconv.Atoi(strings.Trim(v[1], " "))
		}
		if strings.Contains(s, "Register C") {
			v := strings.Split(s, ":")
			r.C, _ = strconv.Atoi(strings.Trim(v[1], " "))
		}
		if strings.Contains(s, "Program") {
			v := strings.Split(s, ":")
			for _, s2 := range strings.Split(strings.Trim(v[1], " "), ",") {
				instr, _ := strconv.Atoi(s2)
				p = append(p, instr)
			}
		}
	}

	return r, p
}

func getComboOperandValue(operand int, r *registers) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	default:
		fmt.Println("Invalid operand")
	}

	return -1
}

func runIntruction(opcode, operand int, r *registers) (int, int) {
	switch opcode {
	case 0:
		// adv i.e. div
		r.A = r.A / int(math.Pow(2.0, float64(getComboOperandValue(operand, r))))
	case 1:
		// bxl i.e. bitwise XOR
		r.B = r.B ^ operand
	case 2:
		// bst
		r.B = getComboOperandValue(operand, r) % 8
	case 3:
		// jnz
		if r.A != 0 {
			instrPointer := operand
			return -1, instrPointer
		}
	case 4:
		// bxc
		r.B = r.B ^ r.C
	case 5:
		// out
		return getComboOperandValue(operand, r) % 8, -1
	case 6:
		// bdv i.e. div
		r.B = r.A / int(math.Pow(2.0, float64(getComboOperandValue(operand, r))))
	case 7:
		// bdv i.e. div
		r.C = r.A / int(math.Pow(2.0, float64(getComboOperandValue(operand, r))))
	}

	return -1, -1
}

func part1(ir registers, p []int) string {
	r := ir
	out := []string{}
	instructionPtr := 0

	for {
		if instructionPtr >= len(p) {
			break
		}

		output, updatePointer := runIntruction(p[instructionPtr], p[instructionPtr+1], &r)
		if output != -1 {
			out = append(out, strconv.Itoa(output))
		}

		if updatePointer == -1 {
			instructionPtr += 2
		} else {
			instructionPtr = updatePointer
		}
	}

	outStr := strings.Join(out, ",")
	return outStr
}

func part2(ir registers, p []int) int {
	rA := -1
	pS := []string{}
	for _, i := range p {
		pS = append(pS, strconv.Itoa(i))
	}
	target := strings.Join(pS, ",")
	// Printing outputs indicates len(output) = n increases every 8^n-1.
	// For a program output of length 16 r.A >= 8^15 r.A < 8^16
	// Last digit follows pattern 5, 7, 1, 0, 3, 2 and repeats 8^n-1 times
	// For a program output with last digit 0 r.A >= 8 ^ 15 + (3 * 8^15) & r.A < 8^16
	// 2nd Last digit follow pattern 1, 5, 6, 4, 1, 0, 3, 2, 5, 2, 1, 5,
	for i := math.Pow(8, 15) + 3*math.Pow(8, 15); i < math.Pow(8, 16); i++ {
		r := ir
		r.A = int(i)
		out := []string{}
		instructionPtr := 0

		for {
			if instructionPtr >= len(p) {
				break
			}

			output, updatePointer := runIntruction(p[instructionPtr], p[instructionPtr+1], &r)
			if output != -1 {
				out = append(out, strconv.Itoa(output))
			}

			if updatePointer == -1 {
				instructionPtr += 2
			} else {
				instructionPtr = updatePointer
			}
		}

		outStr := strings.Join(out, ",")
		//fmt.Println(outStr)
		if outStr == target {
			rA = int(i)
			break
		}
	}

	return rA
}
