package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	//filePath := "./test.txt"
	filePath := "../../AdventOfCodeInputs/2024/Day02/input.txt"
	contents, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Couldn't open file %v: %v", filePath, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	reports := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := []int{}
		for _, s := range strings.Split(line, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Couldn't parse %v from line %v: %v", s, line, err)
				return
			}

			nums = append(nums, n)
		}
		reports = append(reports, nums)
	}

	safeCount := 0
	for i := 0; i < len(reports); i++ {
		safe := true
		previousDiff := reports[i][1] - reports[i][0]
		for j := 1; j < len(reports[i]); j++ {
			diff := reports[i][j] - reports[i][j-1]
			absDiff := abs(diff)
			if absDiff < 1 || absDiff > 3 || ((previousDiff > 0 && diff < 0) || (previousDiff < 0 && diff > 0)) {
				safe = false
				break
			}
			previousDiff = diff
		}

		if safe {
			safeCount++
		}
	}

	fmt.Printf("Part 1: %v\n", safeCount)

	safeCount = 0
	for i := 0; i < len(reports); i++ {
		safe := true
		for k := 0; k < len(reports); k++ {
			safe = true
			previousDiff := reports[i][1] - reports[i][0]
			if k == 0 {
				previousDiff = reports[i][2] - reports[i][1]
			} else if k == 1 {
				previousDiff = reports[i][2] - reports[i][0]
			}
			start := 1
			if k == 0 {
				start = 2
			}
			for j := start; j < len(reports[i]); j++ {
				if j == k {
					continue
				}
				diff := reports[i][j] - reports[i][j-1]
				if j-1 == k {
					diff = reports[i][j] - reports[i][j-2]
				}
				absDiff := abs(diff)
				if absDiff < 1 || absDiff > 3 || ((previousDiff > 0 && diff < 0) || (previousDiff < 0 && diff > 0)) {
					safe = false
					break
				}
				previousDiff = diff
			}
			if safe {
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Printf("Part 2: %v", safeCount)
}
