package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec19.in")

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

	designs, availablePatterns := parseInput(input)
	sort.Strings(availablePatterns)

	numPossible, possibleDesigns := part1(designs, availablePatterns)
	fmt.Println("Part 1: No Of Possible Designs: " + strconv.Itoa(numPossible))

	noOfWaysToMakeDesigns := part2(possibleDesigns, availablePatterns)
	fmt.Printf("Part 2: No Of Ways To Make Possible Designs: " + strconv.Itoa(noOfWaysToMakeDesigns))
}

func parseInput(input []string) ([]string, []string) {
	availablePatterns := []string{}
	// i = 0 available patterns
	s := strings.Split(input[0], ",")
	for _, s2 := range s {
		availablePatterns = append(availablePatterns, strings.TrimSpace(s2))
	}

	designs := []string{}
	// i = 2+ designs (i=1 => blank line)
	for i := 2; i < len(input); i++ {
		designs = append(designs, strings.TrimSpace(input[i]))
	}

	return designs, availablePatterns
}

func isTargetDesignPossible(targetDesign string, currentDesign string, availablePatterns []string, memMap map[string]bool) bool {
	if len(currentDesign) > len(targetDesign) {
		// cur len exceeded target
		return false
	}

	if currentDesign == targetDesign {
		// target design possible
		return true
	}

	if !strings.HasPrefix(targetDesign, currentDesign) {
		// patterns have diverged, no point continuing
		return false
	}

	// append all available patterns to current design and check
	for _, pattern := range availablePatterns {
		cur := currentDesign + pattern
		// check memMap
		if f, ok := memMap[getMemMapKey(targetDesign, cur)]; ok {
			// found in mem - no need to check
			if f {
				return true
			} else {
				continue
			}
		}
		possible := isTargetDesignPossible(targetDesign, cur, availablePatterns, memMap)
		memMap[getMemMapKey(targetDesign, cur)] = possible
		if possible {
			return true
		}
	}

	return false
}

func getMemMapKey(target string, cur string) string {
	return fmt.Sprintf("%s_%s", target, cur)
}

func part1(designs []string, availablePatterns []string) (int, []string) {
	noPossiblePatterns := 0
	possibleDesigns := []string{}
	memMap := map[string]bool{}
	compactedPatterns := []string{}

	// optimize available patterns
	// if a pattern in available patterns is achievable by a combination of other available patterns, remove it
	// this should reduce the number of possible trie entries
	for i, pattern := range availablePatterns {
		subPatternList := []string{}
		subPatternList = append(subPatternList, availablePatterns[:i]...)
		subPatternList = append(subPatternList, availablePatterns[i+1:]...)
		if !isTargetDesignPossible(pattern, "", subPatternList, memMap) {
			compactedPatterns = append(compactedPatterns, pattern)
		}
	}

	// reset map
	memMap = map[string]bool{}
	for _, design := range designs {
		if isTargetDesignPossible(design, "", compactedPatterns, memMap) {
			possibleDesigns = append(possibleDesigns, design)
			noPossiblePatterns++

		}
	}

	return noPossiblePatterns, possibleDesigns
}

func isTargetDesignPossibleP2(targetDesign string, currentDesign string, availablePatterns []string, memMap map[string]int) int {
	sumNoOfWays := 0
	if len(currentDesign) > len(targetDesign) {
		// cur len exceeded target
		return 0
	}

	if currentDesign == targetDesign {
		// target design possible
		return 1
	}

	if !strings.HasPrefix(targetDesign, currentDesign) {
		// patterns have diverged, no point continuing
		return 0
	}

	// check memory
	if n, ok := memMap[getMemMapKey(targetDesign, currentDesign)]; ok {
		return n
	}

	// append all available patterns to current design and check
	for _, pattern := range availablePatterns {
		cur := currentDesign + pattern
		noOfWays := isTargetDesignPossibleP2(targetDesign, cur, availablePatterns, memMap)
		sumNoOfWays += noOfWays
	}

	memMap[getMemMapKey(targetDesign, currentDesign)] = sumNoOfWays
	return sumNoOfWays
}

func part2(designs []string, availablePatterns []string) int {
	sumNoOfPossibleWays := 0

	// Can't compact available patterns here since it's a possible way

	memMap := map[string]int{}
	for _, design := range designs {
		sumNoOfPossibleWays += isTargetDesignPossibleP2(design, "", availablePatterns, memMap)
	}

	return sumNoOfPossibleWays
}
