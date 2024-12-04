package main

import (
	"fmt"
	"os"
)

func findWordFrom(grid [][]rune, r int, c int, word []rune) int {
	if grid[r][c] != word[0] {
		return 0
	}

	total := 0
	found := true
	for i, v := range word {
		if r-i < 0 || grid[r-i][c] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if r-i < 0 || c-i < 0 || grid[r-i][c-i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if r-i < 0 || c+i >= len(grid[r-i]) || grid[r-i][c+i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if r+i >= len(grid) || grid[r+i][c] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if r+i >= len(grid) || c-i < 0 || grid[r+i][c-i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if r+i >= len(grid) || c+i >= len(grid[r+i]) || grid[r+i][c+i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if c-i < 0 || grid[r][c-i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	found = true
	for i, v := range word {
		if c+i >= len(grid[r]) || grid[r][c+i] != v {
			found = false
			break
		}
	}
	if found {
		total++
	}
	return total
}

func findCrossFrom(grid [][]rune, r int, c int, word []rune) int {
	if grid[r][c] != word[1] || r < 1 || c < 1 || r >= len(grid)-1 || c >= len(grid[r])-1 {
		return 0
	}

	if grid[r-1][c-1] == grid[r-1][c+1] && grid[r+1][c-1] == grid[r+1][c+1] {
		if (grid[r-1][c-1] == word[0] && grid[r+1][c+1] == word[2]) || (grid[r-1][c-1] == word[2] && grid[r+1][c+1] == word[0]) {
			return 1
		}
	} else if grid[r-1][c-1] == grid[r+1][c-1] && grid[r-1][c+1] == grid[r+1][c+1] {
		if (grid[r-1][c-1] == word[0] && grid[r+1][c+1] == word[2]) || (grid[r-1][c-1] == word[2] && grid[r+1][c+1] == word[0]) {
			return 1
		}
	}

	return 0
}

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day04/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v", filename, err)
		return
	}

	grid := [][]rune{}
	i := 0
	for _, c := range string(contents) {
		if c == '\n' {
			i++
			continue
		}
		if i == len(grid) {
			grid = append(grid, []rune{})
		}
		grid[i] = append(grid[i], c)
	}

	total := 0
	crossCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			total += findWordFrom(grid, i, j, []rune("XMAS"))
			crossCount += findCrossFrom(grid, i, j, []rune("MAS"))
		}
	}

	fmt.Printf("Part 1: %v\n", total)
	fmt.Printf("Part 2: %v\n", crossCount)
}
