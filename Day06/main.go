package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day06/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't read file %v: %v", filename, err)
		return
	}

	tokens := strings.Split(string(contents), "\n")
	grid := [][]rune{}
	startPos := Pos{}
	direction := -1
	foundGuard := false
	for i, token := range tokens {
		if token == "" {
			continue
		}
		runeArr := []rune(token)
		grid = append(grid, runeArr)
		if !foundGuard {
			for j, r := range runeArr {
				if r != '.' && r != '#' {
					startPos.x = j
					startPos.y = i
					if r == '^' {
						direction = North
					} else if r == '>' {
						direction = East
					} else if r == 'v' {
						direction = South
					} else if r == '<' {
						direction = West
					}
					foundGuard = true
				}
			}
		}
	}
	if direction == -1 {
		fmt.Printf("Direction wasn't set. Pos: (%v, %v). Grid rune: %c\n", startPos.x, startPos.y, grid[startPos.y][startPos.x])
		return
	} else {
		fmt.Printf("Direction set to %v. Pos: (%v, %v). Grid rune: %c\n", direction, startPos.x, startPos.y, grid[startPos.y][startPos.x])
	}

	pos := Pos{startPos.x, startPos.y}
	total := 1
	grid[pos.y][pos.x] = 'X'
	for {
		nextPos := Pos{pos.x, pos.y}
		if direction == North {
			nextPos.y--
		} else if direction == East {
			nextPos.x++
		} else if direction == South {
			nextPos.y++
		} else if direction == West {
			nextPos.x--
		}
		if (nextPos.y < 0 || nextPos.y >= len(grid)) || (nextPos.x < 0 || nextPos.x >= len(grid[nextPos.y])) {
			break
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		pos.x = nextPos.x
		pos.y = nextPos.y
		if grid[pos.y][pos.x] != 'X' {
			total++
			grid[pos.y][pos.x] = 'X'
		}
	}

	fmt.Printf("Part 1: %v\n", total)

	options := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' || (startPos.y == i && startPos.x == j) || ((i == 0 || grid[i-1][j] != 'X') && (j == 0 || grid[i][j-1] != 'X') && (j == len(grid[i])-1 || grid[i][j+1] != 'X') && (i == len(grid)-1 || grid[i+1][j] != 'X')) {
				continue
			}
			pos.x = startPos.x
			pos.y = startPos.y
			direction = North
			escaped := false
			for k := 0; k < 8000; k++ {
				nextPos := Pos{pos.x, pos.y}
				if direction == North {
					nextPos.y--
				} else if direction == East {
					nextPos.x++
				} else if direction == South {
					nextPos.y++
				} else if direction == West {
					nextPos.x--
				}
				if (nextPos.y < 0 || nextPos.y >= len(grid)) || (nextPos.x < 0 || nextPos.x >= len(grid[nextPos.y])) {
					escaped = true
					break
				}

				if grid[nextPos.y][nextPos.x] == '#' || (nextPos.y == i && nextPos.x == j) {
					direction = (direction + 1) % 4
					continue
				}

				pos.x = nextPos.x
				pos.y = nextPos.y
			}
			if !escaped {
				options++
			}
		}
	}

	fmt.Printf("Part 2: %v\n", options)
}

const ( // Direction
	North = 0
	East  = 1
	South = 2
	West  = 3
)

type Pos struct {
	x int
	y int
}

func (p *Pos) move(dx int, dy int) {
	p.x += dx
	p.y += dy
}
