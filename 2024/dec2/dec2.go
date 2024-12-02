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

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec2.in")

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

	safeReports := part1(input)
	fmt.Println("Part 1: No of safe reports: " + strconv.Itoa(safeReports))

	safeReportsWithForgiveness := part2(input)
	fmt.Println("Part 2: No of safe reports with forgiveness: " + strconv.Itoa(safeReportsWithForgiveness))

	safeReportsWithForgiveness = part2BruteForce(input)
	fmt.Println("Part 2: No of safe reports with forgiveness (brute force): " + strconv.Itoa(safeReportsWithForgiveness))
}

func isReportSafe(report []int) bool {
	direction := 0
	for i := 1; i < len(report); i++ {
		delta := report[i] - report[i-1]
		if math.Abs(float64(delta)) < 1.0 || math.Abs(float64(delta)) > 3.0 {
			return false
		}

		if i == 1 {
			if math.Abs(float64(delta)) < 1.0 || math.Abs(float64(delta)) > 3.0 {
				return false
			}
			direction = (delta) / int(math.Abs(float64(delta)))
			continue
		}

		if report[i] > report[i-1] && direction != 1 || report[i] < report[i-1] && direction != -1 {
			return false
		}
	}

	return true
}

func isReportSafeWithForgiveness(report []int) bool {
	direction := 0
	forgivenessAllowed := true
	forgivenessIndex := -1
	for i := 1; i < len(report); i++ {
		delta := report[i] - report[i-1]
		if forgivenessIndex != -1 && i-1 == forgivenessIndex {
			// skip number
			delta = report[i] - report[i-2]
		}
		if i == 1 {
			if math.Abs(float64(delta)) < 1.0 || math.Abs(float64(delta)) > 3.0 {
				if forgivenessAllowed {
					forgivenessAllowed = false
					forgivenessIndex = i
					continue
				} else {
					return false
				}
			}
			direction = (delta) / int(math.Abs(float64(delta)))
			continue
		}

		if math.Abs(float64(delta)) < 1.0 || math.Abs(float64(delta)) > 3.0 {
			if forgivenessAllowed {
				forgivenessAllowed = false
				forgivenessIndex = i
				continue
			} else {
				return false
			}
		}

		if i == 2 && forgivenessAllowed == false {
			if delta == 0 {
				return false
			}
			// skipped index 1 so recalculate direction
			direction = (delta) / int(math.Abs(float64(delta)))
		}

		if report[i] > report[i-1] && direction != 1 || report[i] < report[i-1] && direction != -1 {
			if forgivenessAllowed {
				forgivenessAllowed = false
				forgivenessIndex = i
				continue
			} else {
				return false
			}
		}
	}

	return true
}

func part1(input []string) int {
	sumSafeReports := 0
	for _, s := range input {
		report := []int{}
		reportS := strings.Split(s, " ")
		for _, v := range reportS {
			vInt, _ := strconv.Atoi(v)
			report = append(report, vInt)
		}
		if isReportSafe(report) {
			sumSafeReports++
		}
	}
	return sumSafeReports
}

func part2(input []string) int {
	sumSafeReports := 0
	for _, s := range input {
		report := []int{}
		reportS := strings.Split(s, " ")
		for _, v := range reportS {
			vInt, _ := strconv.Atoi(v)
			report = append(report, vInt)
		}
		if isReportSafeWithForgiveness(report) {
			sumSafeReports++
		} else {
			// Doesn't account for case 0
			if isReportSafe(report[1:]) {
				sumSafeReports++
			}
		}
	}
	return sumSafeReports
}

func part2BruteForce(input []string) int {
	sumSafeReports := 0
	for _, s := range input {
		report := []int{}
		reportS := strings.Split(s, " ")
		for _, v := range reportS {
			vInt, _ := strconv.Atoi(v)
			report = append(report, vInt)
		}
		if isReportSafe(report) {
			sumSafeReports++
		} else {
			// remove each level and check
			for i := 0; i < len(report); i++ {
				shortReport := []int{}
				shortReport = append(shortReport, report[:i]...)
				shortReport = append(shortReport, report[i+1:]...)
				if isReportSafe(shortReport) {
					sumSafeReports++
					break
				}
			}
		}
	}
	return sumSafeReports
}
