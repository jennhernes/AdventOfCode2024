package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day05/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	rules := map[int][]int{}
	processingRules := true
	totalPart1 := 0
	totalPart2 := 0
	for _, line := range lines {
		if line == "" {
			processingRules = false
			continue
		}
		if processingRules {
			pagesStr := strings.Split(line, "|")
			last, err := strconv.Atoi(pagesStr[1])
			if err != nil {
				fmt.Printf("Couldn't parse %v as int: %v", pagesStr[1], err)
				return
			}
			before, exists := rules[last]
			if !exists {
				before = []int{}
			}
			first, err := strconv.Atoi(pagesStr[0])
			if err != nil {
				fmt.Printf("Couldn't parse %v as int: %v", pagesStr[0], err)
				return
			}
			if !slices.Contains(before, first) {
				before = append(before, first)
			}
			rules[last] = before
		} else {
			correct := true
			pagesStr := strings.Split(line, ",")
			pages := []int{}
			for _, ps := range pagesStr {
				p, err := strconv.Atoi(ps)
				if err != nil {
					fmt.Printf("Couldn't parse %v as int: %v", ps, err)
					return
				}
				pages = append(pages, p)
			}
			for i := 0; i < len(pages); i++ {
				requiredAfter, exists := rules[pages[i]]
				if !exists {
					continue
				}
				for j := i + 1; j < len(pages); j++ {
					if slices.Contains(requiredAfter, pages[j]) {
						correct = false
						temp := pages[i]
						pages[i] = pages[j]
						pages[j] = temp
						i = -1
						j = -1
						break
					}
				}
			}
			if correct {
				totalPart1 += pages[(len(pages))/2]
			} else {
				totalPart2 += pages[(len(pages))/2]
			}
		}
	}

	fmt.Printf("Part 1: %v\n", totalPart1)
	fmt.Printf("Part 2: %v\n", totalPart2)
}
