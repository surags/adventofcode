package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("sample.in")
	// file, err := os.Open("sample2.in")
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

	part1(input)
	part2(input)
}

func part1(input []string) {
	cvSum := 0
	for _, str := range input {
		fDig := 0
		lDig := 0
		str_len := len(str)
		for i := 0; i < str_len; i++ {
			dig, err := strconv.Atoi(string(str[i]))
			if err == nil {
				fDig = dig
				break
			}
		}

		for i := str_len - 1; i >= 0; i-- {
			dig, err := strconv.Atoi(string(str[i]))
			if err == nil {
				lDig = dig
				break
			}
		}

		cal := fDig*10 + lDig
		cvSum += cal
	}

	fmt.Printf("Part 1: %d\n", cvSum)
}

var acceptableDigits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digitMap = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func getDigits(digits []string, str string) []string {
	smallIndex := 10000000000
	firstDigit := ""
	found := false
	for _, digit := range acceptableDigits {
		i := strings.Index(str, digit)
		if i != -1 {
			if i < smallIndex {
				found = true
				firstDigit = digit
				smallIndex = i
			}
		}
	}

	if found == true {
		digits = append(digits, firstDigit)
		startIndex := smallIndex + max(len(firstDigit)-1, 1)
		digits = getDigits(digits, str[startIndex:])
	}

	return digits
}

func part2(input []string) {
	cvSum := 0
	for _, str := range input {
		presentDigits := []string{}
		presentDigits = getDigits(presentDigits, str)

		cal := digitMap[presentDigits[0]]*10 + digitMap[presentDigits[len(presentDigits)-1]]
		// fmt.Printf("%s \t %v \t %d \n", str, presentDigits, cal)
		cvSum += cal
	}

	fmt.Printf("Part 2: %d\n", cvSum)
}
