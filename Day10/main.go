package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day10/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")

	grid := [][]rune{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	total := 0
	repeats := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				result := findTrails(i, j, grid)
				slices.SortFunc(result, func(a, b Pos) int {
					if a.x == b.x {
						return cmp.Compare(a.y, b.y)
					}
					return cmp.Compare(a.x, b.x)
				})
				total += len(result)
				for k := 0; k < len(result)-1; k++ {
					if result[k].x == result[k+1].x && result[k].y == result[k+1].y {
						total--
						repeats++
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %v\n", total)
	fmt.Printf("Part 2: %v\n", total+repeats)
}

func findTrails(i int, j int, grid [][]rune) []Pos {
	if grid[i][j] == '9' {
		return []Pos{{j, i}}
	}
	nextHeight := grid[i][j] + 1
	result := []Pos{}
	if i > 0 && grid[i-1][j] == nextHeight {
		result = append(result, findTrails(i-1, j, grid)...)
	}
	if j > 0 && grid[i][j-1] == nextHeight {
		result = append(result, findTrails(i, j-1, grid)...)
	}
	if i < len(grid)-1 && grid[i+1][j] == nextHeight {
		result = append(result, findTrails(i+1, j, grid)...)
	}
	if j < len(grid[i])-1 && grid[i][j+1] == nextHeight {
		result = append(result, findTrails(i, j+1, grid)...)
	}
	return result
}

type Pos struct {
	x int
	y int
}
