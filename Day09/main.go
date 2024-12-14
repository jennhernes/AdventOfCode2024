package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day09/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v", filename, err)
		return
	}
	lines := strings.Split(string(contents), "\n'")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		disk := []int{}
		runes := []rune(line)
		id := 0
		free := false
		for _, r := range runes {
			if free {
				size := int(r - '0')
				for i := 0; i < size; i++ {
					disk = append(disk, -1)
				}
			} else {
				size := int(r - '0')
				for i := 0; i < size; i++ {
					disk = append(disk, id)
				}
				id++
			}
			free = !free
		}
		j := len(disk) - 1
		for {
			if disk[j] == -1 {
				j--
			} else {
				break
			}
		}
		for i := 0; i < len(disk); i++ {
			if i >= j {
				break
			}
			if disk[i] == -1 {
				disk[i] = disk[j]
				disk[j] = -1
				for {
					j--
					if disk[j] != -1 {
						break
					}
				}
			}
		}
		for i := 0; i < len(disk); i++ {
			if disk[i] == -1 {
				break
			}
			//fmt.Printf("%v * %v\n", i, disk[i])
			total += i * disk[i]
		}
	}

	fmt.Printf("Part 1: %v\n", total)

	total = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		disk := []int{}
		runes := []rune(line)
		id := 0
		free := false
		for _, r := range runes {
			if free {
				size := int(r - '0')
				for i := 0; i < size; i++ {
					disk = append(disk, -1)
				}
			} else {
				size := int(r - '0')
				for i := 0; i < size; i++ {
					disk = append(disk, id)
				}
				id++
			}
			free = !free
		}

		for i := len(disk) - 1; i >= 0; i-- {
			if disk[i] == -1 || disk[i] == 0 {
				continue
			}
			fileSize := 0
			for j := 0; j < i; j++ {
				if disk[i-j] != disk[i] {
					break
				}
				fileSize++
			}
			fmt.Printf("File size for %v: %v\n", disk[i], fileSize)

			sizingFreeSpace := false
			freeSize := 0
			for j := 0; j < i; j++ {
				if sizingFreeSpace {
					if disk[j] != -1 {
						if freeSize >= fileSize {
							for k := 0; k < fileSize; k++ {
								disk[j-freeSize+k] = disk[i]
							}
							for k := 0; k < fileSize; k++ {
								disk[i-k] = -1
							}
							break
						}
						sizingFreeSpace = false

					} else {
						freeSize++
					}
				} else {
					freeSize = 0
					if disk[j] == -1 {
						freeSize++
						sizingFreeSpace = true
					}
				}
			}
			if fileSize > 0 {
				i -= fileSize
				i++
			}
		}

		for i := 0; i < len(disk); i++ {
			if disk[i] == -1 {
				continue
			}
			//fmt.Printf("%v * %v\n", i, disk[i])
			total += i * disk[i]
		}
	}

	fmt.Printf("Part 2: %v\n", total)
}
