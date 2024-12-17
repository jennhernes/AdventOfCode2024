package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "./test.txt"
	//filename := "../../AdventOfCodeInputs/2024/Day12/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Could not open file %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	grid := [][]rune{}
	checked := [][]bool{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
		b := []bool{}
		for i := 0; i < len(grid[0]); i++ {
			b = append(b, false)
		}
		checked = append(checked, b)
	}

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if checked[i][j] {
				continue
			}
			area, perim := findAreaAndPerimeter(i, j, grid, checked)
			total += area * perim
		}
	}

	fmt.Printf("Part 1: %v\n", total)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			checked[i][j] = false
		}
	}
	total = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if checked[i][j] {
				continue
			}
			area, perim := findAreaAndSides(i, j, grid, checked)
			total += area * perim
			//if grid[i][j] == 'F' {
			//	return
			//}
		}
	}

	fmt.Printf("Part 2: %v\n", total)
}

func findAreaAndSides(i int, j int, grid [][]rune, checked [][]bool) (int, int) {
	area, _ := findAreaAndPerimeter(i, j, grid, checked)
	direction := Right
	prevDirection := Up
	sides := 1
	m := i
	n := j
	//k := 0
	for {
		//if grid[i][j] == 'F' {
		//	if k > 30 {
		//		return 0, 0
		//	}
		//	fmt.Printf("Region %c: [%v][%v] -> %v\n", grid[i][j], i, j, direction)
		//	k++
		//}
		if m == i && n == j && direction == Up {
			break
		}
		switch direction {
		case Right:
			if prevDirection != (direction+1)%TotalDirections && i > 0 && grid[i-1][j] == grid[i][j] {
				prevDirection = direction
				direction = (direction + TotalDirections - 1) % TotalDirections
				sides++
				continue
			} else if j < len(grid[i])-1 && grid[i][j+1] == grid[i][j] {
				prevDirection = direction
				j++
				continue
			} else {
				prevDirection = direction
				direction = (direction + 1) % TotalDirections
				sides++
				continue
			}
		case Down:
			if prevDirection != (direction+1)%TotalDirections && j < len(grid[i])-1 && grid[i][j+1] == grid[i][j] {
				prevDirection = direction
				direction = (direction + TotalDirections - 1) % TotalDirections
				sides++
				continue
			} else if i < len(grid)-1 && grid[i+1][j] == grid[i][j] {
				prevDirection = direction
				i++
				continue
			} else {
				prevDirection = direction
				direction = (direction + 1) % TotalDirections
				sides++
				continue
			}
		case Left:
			if prevDirection != (direction+1)%TotalDirections && i < len(grid[i])-1 && grid[i+1][j] == grid[i][j] {
				prevDirection = direction
				direction = (direction + TotalDirections - 1) % TotalDirections
				sides++
				continue
			} else if j > 0 && grid[i][j-1] == grid[i][j] {
				prevDirection = direction
				j--
				continue
			} else {
				prevDirection = direction
				direction = (direction + 1) % TotalDirections
				sides++
				continue
			}
		case Up:
			if prevDirection != (direction+1)%TotalDirections && j > 0 && grid[i][j-1] == grid[i][j] {
				prevDirection = direction
				direction = (direction + TotalDirections - 1) % TotalDirections
				sides++
				continue
			} else if i > 0 && grid[i-1][j] == grid[i][j] {
				prevDirection = direction
				i--
				continue
			} else {
				prevDirection = direction
				direction = (direction + 1) % TotalDirections
				sides++
				continue
			}
		}
	}

	//fmt.Printf("Region %c: area = %v, sides = %v\n", grid[i][j], area, sides)
	return area, sides
}

const (
	Right           = 0
	Down            = 1
	Left            = 2
	Up              = 3
	TotalDirections = 4
)

func findAreaAndPerimeter(i int, j int, grid [][]rune, checked [][]bool) (int, int) {
	if checked[i][j] {
		return 0, 0
	}
	area := 1
	perim := 0
	checked[i][j] = true
	if i == 0 || grid[i-1][j] != grid[i][j] {
		perim++
	} else if i > 0 {
		a, p := findAreaAndPerimeter(i-1, j, grid, checked)
		area += a
		perim += p
	}

	if j == 0 || grid[i][j-1] != grid[i][j] {
		perim++
	} else if j > 0 {
		a, p := findAreaAndPerimeter(i, j-1, grid, checked)
		area += a
		perim += p
	}
	if i == len(grid)-1 || grid[i+1][j] != grid[i][j] {
		perim++
	} else if i < len(grid)-1 {
		a, p := findAreaAndPerimeter(i+1, j, grid, checked)
		area += a
		perim += p
	}
	if j == len(grid[i])-1 || grid[i][j+1] != grid[i][j] {
		perim++
	} else if j < len(grid[i])-1 {
		a, p := findAreaAndPerimeter(i, j+1, grid, checked)
		area += a
		perim += p
	}

	return area, perim
}
