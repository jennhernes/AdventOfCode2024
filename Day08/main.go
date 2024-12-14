package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day08/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v\n", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	grid := [][]rune{}
	antennae := map[rune][]Pos{}
	for i, line := range lines {
		if line == "" {
			continue
		}
		runes := []rune(line)
		grid = append(grid, runes)
		for j, r := range runes {
			if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '0' || r > '9') {
				continue
			}
			if _, exists := antennae[r]; !exists {
				antennae[r] = []Pos{}
			}
			antennae[r] = append(antennae[r], Pos{j, i})
		}
	}

	antinodes := []Pos{}
	for _, posList := range antennae {
		if len(posList) < 2 {
			continue
		}
		for i := 0; i < len(posList)-1; i++ {
			for j := i + 1; j < len(posList); j++ {
				node1 := Pos{}
				node2 := Pos{}
				node1.x = posList[i].x + (posList[i].x - posList[j].x)
				node1.y = posList[i].y + (posList[i].y - posList[j].y)
				node2.x = posList[j].x + (posList[j].x - posList[i].x)
				node2.y = posList[j].y + (posList[j].y - posList[i].y)
				//fmt.Printf("(%v, %v) -> (%v, %v): y = %vx + %v\n", posList[i].x, posList[i].y, posList[j].x, posList[j].y, m, b)
				if node1.x >= 0 && node1.y >= 0 && node1.x < len(grid[0]) && node1.y < len(grid) {
					//fmt.Printf("Node1 (%v, %v)\n", node1.x, node1.y)
					antinodes = append(antinodes, node1)
				}
				if node2.x >= 0 && node2.y >= 0 && node2.x < len(grid[0]) && node2.y < len(grid) {
					//fmt.Printf("Node2 (%v, %v)\n", node2.x, node2.y)
					antinodes = append(antinodes, node2)
				}
			}
		}
	}

	nodeSet := map[string]int{}
	for _, p := range antinodes {
		//fmt.Printf("Node rune: '%c'\n", r)
		//fmt.Printf("(%v, %v)T\n", p.y, p.x)
		nodeSet[fmt.Sprintf("%v,%v", p.x, p.y)] = 1
	}

	fmt.Printf("Part 1: %v\n", len(nodeSet))

	antinodes = []Pos{}
	for _, posList := range antennae {
		if len(posList) < 2 {
			continue
		}
		antinodes = append(antinodes, posList[len(posList)-1])
		for i := 0; i < len(posList)-1; i++ {
			antinodes = append(antinodes, posList[i])
			for j := i + 1; j < len(posList); j++ {
				for k := 1; ; k++ {
					node1 := Pos{}
					node2 := Pos{}
					node1.x = posList[i].x + k*(posList[i].x-posList[j].x)
					node1.y = posList[i].y + k*(posList[i].y-posList[j].y)
					node2.x = posList[j].x + k*(posList[j].x-posList[i].x)
					node2.y = posList[j].y + k*(posList[j].y-posList[i].y)
					addedNode := false
					//fmt.Printf("(%v, %v) -> (%v, %v): y = %vx + %v\n", posList[i].x, posList[i].y, posList[j].x, posList[j].y, m, b)
					if node1.x >= 0 && node1.y >= 0 && node1.x < len(grid[0]) && node1.y < len(grid) {
						//fmt.Printf("Node1 (%v, %v)\n", node1.x, node1.y)
						antinodes = append(antinodes, node1)
						addedNode = true
					}
					if node2.x >= 0 && node2.y >= 0 && node2.x < len(grid[0]) && node2.y < len(grid) {
						//fmt.Printf("Node2 (%v, %v)\n", node2.x, node2.y)
						antinodes = append(antinodes, node2)
						addedNode = true
					}
					if !addedNode {
						break
					}
				}
			}
		}
	}

	nodeSet = map[string]int{}
	for _, p := range antinodes {
		//fmt.Printf("Node rune: '%c'\n", r)
		//fmt.Printf("(%v, %v)T\n", p.y, p.x)
		nodeSet[fmt.Sprintf("%v,%v", p.x, p.y)] = 1
	}

	fmt.Printf("Part 2: %v\n", len(nodeSet))
}

type Pos struct {
	x int
	y int
}

func dist(a, b Pos) int {
	dx := a.x - b.x
	if dx < 0 {
		dx *= -1
	}
	dy := a.y - b.y
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}
