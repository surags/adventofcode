package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type graphNode struct {
	name  string
	left  *graphNode
	right *graphNode
}

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	//file, err := os.Open("sample3.in")
	file, err := os.Open("dec8.in")

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
	directions, nodeMap := generateGraph(input)

	part1(directions, nodeMap["AAA"])
	part2(directions, nodeMap)
}

func generateGraph(input []string) (string, map[string]*graphNode) {
	directions := strings.TrimSpace(input[0])
	nodeMap := map[string]*graphNode{}

	for i := 2; i < len(input); i++ {
		line := input[i]

		splt := strings.Split(line, "=")
		nodeName := strings.TrimSpace(splt[0])

		node, exists := nodeMap[nodeName]
		if !exists {
			nodeMap[nodeName] = &graphNode{name: nodeName}
			node = nodeMap[nodeName]
		}

		lrSplt := strings.Split(strings.TrimSpace(splt[1]), ",")
		l := strings.TrimSpace(lrSplt[0][1:])
		r := strings.TrimSpace(lrSplt[1][:len(lrSplt[1])-1])

		lNode, exists := nodeMap[l]
		if !exists {
			nodeMap[l] = &graphNode{name: l}
			lNode = nodeMap[l]
		}

		rNode, exists := nodeMap[r]
		if !exists {
			nodeMap[r] = &graphNode{name: r}
			rNode = nodeMap[r]
		}

		node.left = lNode
		node.right = rNode
	}

	return directions, nodeMap
}

func part1(directions string, startNode *graphNode) {
	if startNode == nil {
		return
	}

	noOfSteps := 0
	curNode := startNode
	found := false

	for i := 0; i < 1; i-- {
		for _, v := range directions {
			direction := string(v)
			noOfSteps++
			if direction == "L" {
				curNode = curNode.left
			} else {
				curNode = curNode.right
			}

			if curNode.name == "ZZZ" {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Printf("Part 1: No of Steps: %d\n", noOfSteps)
}

func isZNode(node *graphNode) bool {
	lastIndex := len(node.name) - 1

	if string(node.name[lastIndex]) == "Z" {
		return true
	}

	return false
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(directions string, nodeMap map[string]*graphNode) {
	noOfSteps := 0
	startNodes := []*graphNode{}
	curNodes := []*graphNode{}
	reachZFreq := []int{}
	reachZ := map[string]interface{}{}

	for nodeName, node := range nodeMap {
		lastIndex := len(nodeName) - 1

		if string(nodeName[lastIndex]) == "A" {
			startNodes = append(curNodes, node)
			curNodes = append(curNodes, node)
			//reachZ[nodeName] = false
		}
	}

	noOfToProcessNodes := len(startNodes)
	foundNodes := 0
	found := false

	for i := 0; i < 1; i-- {
		for _, v := range directions {
			direction := string(v)
			noOfSteps++
			for i, _ := range curNodes {
				if direction == "L" {
					curNodes[i] = curNodes[i].left
				} else {
					curNodes[i] = curNodes[i].right
				}

				if isZNode(curNodes[i]) {
					//nodeCountMap[startNodes[i].name] = append(nodeCountMap[startNodes[i].name], noOfSteps)
					_, reached := reachZ[startNodes[i].name]
					if !reached {
						reachZFreq = append(reachZFreq, noOfSteps)
						foundNodes++
					}

					if foundNodes == noOfToProcessNodes {
						found = true
						break
					}
					//fmt.Printf("%v \n", nodeCountMap)
				}
			}
		}

		if found {
			break
		}
	}

	noOfTrueSteps := LCM(reachZFreq[0], reachZFreq[1], reachZFreq[2:]...)

	fmt.Printf("Part 2: No of Steps: %d\n", noOfTrueSteps)
}
