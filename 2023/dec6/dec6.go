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
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec6.in")

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

	times, distances := parseTimesAndDistance(input)
	part1(times, distances)
	part2(times, distances)
}

func parseTimesAndDistance(input []string) ([]int, []int) {
	timesStr := strings.Split(strings.TrimSpace(strings.Split(input[0], ":")[1]), " ")
	var times []int
	for _, timeStr := range timesStr {
		if timeStr == "" {
			continue
		}
		time, _ := strconv.Atoi(timeStr)
		times = append(times, time)
	}

	distsStr := strings.Split(strings.TrimSpace(strings.Split(input[1], ":")[1]), " ")
	var distances []int
	for _, distanceStr := range distsStr {
		if distanceStr == "" {
			continue
		}
		dist, _ := strconv.Atoi(distanceStr)
		distances = append(distances, dist)
	}

	return times, distances
}

func part1(times, distances []int) {
	noOfWays := 1
	noOfRaces := len(times)
	for i := 0; i < noOfRaces; i++ {
		time := times[i]
		distance := distances[i]

		successCount := 0
		lastDistance := 0
		for holdTime := 1; holdTime <= time; holdTime++ {
			speed := holdTime
			distanceTravelled := speed * (time - holdTime)
			if distanceTravelled > distance {
				successCount++
			}

			if distanceTravelled < lastDistance {
				break
			}
		}

		noOfWays *= successCount
	}

	fmt.Printf("Part 1: No Of Ways: %d\n", noOfWays)
}

func part2(times, distances []int) {
	newRaceTimeStr := ""
	newRaceDistStr := ""

	for i := 0; i < len(times); i++ {
		newRaceTimeStr = newRaceTimeStr + fmt.Sprintf("%d", times[i])
		newRaceDistStr = newRaceDistStr + fmt.Sprintf("%d", distances[i])
	}

	newRaceTime, _ := strconv.Atoi(newRaceTimeStr)
	newRaceDist, _ := strconv.Atoi(newRaceDistStr)

	successCount := 0
	lastDistance := 0
	for holdTime := 1; holdTime <= newRaceTime; holdTime++ {
		speed := holdTime
		distanceTravelled := speed * (newRaceTime - holdTime)
		if distanceTravelled > newRaceDist {
			successCount++
		}

		if distanceTravelled < lastDistance {
			break
		}
	}

	fmt.Printf("Part 2: No Of Ways: %d\n", successCount)
}
