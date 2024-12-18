package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day13/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	total := 0
	for i := 0; i < len(lines); i++ {
		a, err := processButton(lines[i])
		if err != nil {
			return
		}
		i++
		b, err := processButton(lines[i])
		if err != nil {
			return
		}
		i++
		p, err := processPrize(lines[i])
		if err != nil {
			return
		}
		i++

		m := (p.x*b.y - p.y*b.x) / (a.x*b.y - b.x*a.y)
		n := (p.y*a.x - p.x*a.y) / (a.x*b.y - b.x*a.y)
		fmt.Printf("a.x = %v, a.y = %v; b.x = %v, b.y = %v; p.x = %v, p.y = %v\n", a.x, a.y, b.x, b.y, p.x, p.y)
		fmt.Printf("A = %v, B = %v\n", m, n)
		if a.x*m+b.x*n != p.x || a.y*m+b.y*n != p.y {
			continue
		}
		total += 3*m + n
		fmt.Printf("i = %v\n", i)
	}

	fmt.Printf("Part 1: %v\n", total)

	total = 0
	for i := 0; i < len(lines); i++ {
		a, err := processButton(lines[i])
		if err != nil {
			return
		}
		i++
		b, err := processButton(lines[i])
		if err != nil {
			return
		}
		i++
		p, err := processPrize(lines[i])
		if err != nil {
			return
		}
		p.x += 10000000000000
		p.y += 10000000000000
		i++

		m := (p.x*b.y - p.y*b.x) / (a.x*b.y - b.x*a.y)
		n := (p.y*a.x - p.x*a.y) / (a.x*b.y - b.x*a.y)
		fmt.Printf("a.x = %v, a.y = %v; b.x = %v, b.y = %v; p.x = %v, p.y = %v\n", a.x, a.y, b.x, b.y, p.x, p.y)
		fmt.Printf("A = %v, B = %v\n", m, n)
		if a.x*m+b.x*n != p.x || a.y*m+b.y*n != p.y {
			continue
		}
		total += 3*m + n
		fmt.Printf("i = %v\n", i)
	}

	fmt.Printf("Part 2: %v\n", total)
}

func processPrize(line string) (*Vector, error) {
	tokens := strings.Split(line, " ")
	x, err := strconv.Atoi(tokens[1][2 : len(tokens[1])-1])
	if err != nil {
		fmt.Printf("Could not convert %v to int: %v", tokens[1], err)
		return nil, err
	}
	y, err := strconv.Atoi(tokens[2][2:])
	if err != nil {
		fmt.Printf("Could not convert %v to int: %v", tokens[2], err)
		return nil, err
	}
	result := &Vector{
		x: x,
		y: y,
	}
	return result, nil
}

func processButton(line string) (*Vector, error) {
	tokens := strings.Split(line, " ")
	for _, s := range tokens {
		fmt.Printf("'%v' ", s)
	}
	fmt.Println()
	x, err := strconv.Atoi(tokens[2][2 : len(tokens[2])-1])
	if err != nil {
		fmt.Printf("Could not convert %v to int: %v", tokens[2], err)
		return nil, err
	}
	y, err := strconv.Atoi(tokens[3][2:])
	if err != nil {
		fmt.Printf("Could not convert %v to int: %v", tokens[3], err)
		return nil, err
	}
	result := &Vector{
		x: x,
		y: y,
	}
	return result, nil
}

type Vector struct {
	x int
	y int
}
