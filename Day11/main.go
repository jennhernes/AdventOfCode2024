package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day11/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		stones := []map[int]int{}
		stonesStr := strings.Split(line, " ")
		stonesInt := map[int]int{}
		for i := 0; i < len(stonesStr); i++ {
			if stonesStr[i] == "" || stonesStr[i] == "\n" {
				continue
			}
			s, err := strconv.Atoi(stonesStr[i])
			if err != nil {
				fmt.Printf("Could not parse string '%v' as int: %v", stonesStr[i], err)
				return
			}
			if _, exists := stonesInt[s]; !exists {
				stonesInt[s] = 0
			}
			stonesInt[s]++
		}
		stones = append(stones, stonesInt)
		stones = append(stones, stonesInt)
		for i := 0; i < 75; i++ {
			oldIndex := i % 2
			newIndex := (i + 1) % 2
			stones[newIndex] = map[int]int{}
			for k, v := range stones[oldIndex] {
				if k == 0 {
					if _, exists := stones[newIndex][1]; !exists {
						stones[newIndex][1] = 0
					}
					stones[newIndex][1] += v
				} else if s := strconv.Itoa(k); len(s)%2 == 0 {
					left, err := strconv.Atoi(s[:len(s)/2])
					if err != nil {
						fmt.Printf("Could not parse string '%v' as int: %v", s, err)
						return
					}
					if _, exists := stones[newIndex][left]; !exists {
						stones[newIndex][left] = 0
					}
					stones[newIndex][left] += v
					right, err := strconv.Atoi(s[len(s)/2:])
					if err != nil {
						fmt.Printf("Could not parse string '%v' as int: %v", s, err)
						return
					}
					if _, exists := stones[newIndex][right]; !exists {
						stones[newIndex][right] = 0
					}
					stones[newIndex][right] += v
				} else {
					n := k * 2024
					if _, exists := stones[newIndex][n]; !exists {
						stones[newIndex][n] = 0
					}
					stones[newIndex][n] += v
				}
			}
			if i == 24 {
				total = 0
				for _, v := range stones[1] {
					total += v
				}
				fmt.Printf("Part 1: %v\n", total)
			}
		}
		total = 0
		for _, v := range stones[1] {
			total += v
		}
		fmt.Printf("Part 2: %v\n", total)
	}
}
