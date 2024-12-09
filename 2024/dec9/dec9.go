package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	//file, err := os.Open("sample.in")
	file, err := os.Open("dec9.in")

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

	disk, diskAltRep := parseInput(input)
	//fmt.Println(disk)

	checksum := part1(disk)
	fmt.Println("Part 1: Compacted File System Checksum (fragmented): " + strconv.Itoa(checksum))

	checksum2 := part2(disk, diskAltRep)
	fmt.Println("Part 2: Compacted File System Checksum (defragmented): " + strconv.Itoa(checksum2))
}

type diskBlock struct {
	blockType int // 0 file 1 space
	id        int
	length    int
}

func parseInput(input []string) ([]string, []diskBlock) {
	id := 0
	isFile := true
	disk := []string{}
	diskAltRep := []diskBlock{}
	for _, blockI := range input[0] {
		db := diskBlock{}
		block, _ := strconv.Atoi(string(blockI))
		db.length = block
		for i := 0; i < block; i++ {
			if isFile {
				disk = append(disk, strconv.Itoa(id))
			} else {
				disk = append(disk, ".")
			}
		}
		if !isFile {
			db.blockType = 1
			id++
		} else {
			db.blockType = 0
			db.id = id
		}
		isFile = !isFile
		diskAltRep = append(diskAltRep, db)
	}
	return disk, diskAltRep
}

func compactDisk(disk []string) []string {
	compactedDisk := make([]string, len(disk))
	for i := 0; i < len(disk); i++ {
		compactedDisk[i] = disk[i]
	}

	i := 0
	j := len(disk) - 1

	for {
		if i >= j {
			break
		}

		if compactedDisk[j] == "." {
			j--
			continue
		}

		if compactedDisk[i] == "." {
			// free position to swap
			compactedDisk[i] = compactedDisk[j]
			compactedDisk[j] = "."
			j--
		}
		i++
	}

	return compactedDisk
}

func part1(disk []string) int {
	compactedDisk := compactDisk(disk)
	checksum := 0
	for i := 0; i < len(disk); i++ {
		if compactedDisk[i] == "." {
			continue
		}
		block, _ := strconv.Atoi(compactedDisk[i])
		checksum += i * block
	}

	return checksum
}

func compactDiskWithoutFragmentation(disk []string, diskAltRep []diskBlock) []string {
	compactedDisk := make([]string, len(disk))

	compactedDiskAltRep := make([]diskBlock, len(diskAltRep))
	for i := 0; i < len(diskAltRep); i++ {
		compactedDiskAltRep[i] = diskAltRep[i]
	}

	for k := len(diskAltRep) - 1; k >= 0; k-- {
		db := diskAltRep[k]
		// find a free spot to put this
		for i, b := range compactedDiskAltRep {
			if b.blockType != 1 {
				continue
			}
			if db.length > b.length {
				continue
			}
			if db.id <= b.id {
				break
			}

			// split current space into block + remaining space
			compactedDiskAltRep[i].length = b.length - db.length
			compactedDiskAltRep = slices.Insert(compactedDiskAltRep, i, db)
			break
		}
	}

	index := 0
	visitedSet := map[int]bool{}
	for _, block := range compactedDiskAltRep {
		_, visited := visitedSet[block.id]
		for i := 0; i < block.length; i++ {
			if block.blockType == 1 || visited {
				compactedDisk[index] = "."
			} else {
				compactedDisk[index] = strconv.Itoa(block.id)
			}
			index++
			visitedSet[block.id] = true
		}

	}

	return compactedDisk
}

func part2(disk []string, diskAltRep []diskBlock) int {
	compactedDisk := compactDiskWithoutFragmentation(disk, diskAltRep)
	checksum := 0
	for i := 0; i < len(disk); i++ {
		block, _ := strconv.Atoi(compactedDisk[i])
		checksum += i * block
	}

	return checksum
}
