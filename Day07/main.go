package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day07/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v", filename, err)
		return
	}

	total := 0
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, ": ")
		if len(tokens) != 2 {
			fmt.Printf("Not 2 tokens: %v", tokens)
			return
		}
		goal, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Could not process '%v': %v", tokens[0], err)
			return
		}
		pieces := []int{}
		for _, s := range strings.Split(tokens[1], " ") {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Could not process '%v': %v", s, err)
				return
			}
			pieces = append(pieces, i)
		}
		results := [][]int{}
		layer := []int{pieces[0], pieces[0]}
		results = append(results, layer)
		for i := 1; i < len(pieces); i++ {
			layer = []int{}
			for _, x := range results[i-1] {
				layer = append(layer, x+pieces[i], x*pieces[i])
			}
			results = append(results, layer)
		}
		for _, x := range results[len(results)-1] {
			if x == goal {
				total += goal
				break
			}
		}
	}

	fmt.Printf("Part 1: %v\n", total)

	total = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, ": ")
		if len(tokens) != 2 {
			fmt.Printf("Not 2 tokens: %v", tokens)
			return
		}
		goal, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Could not process '%v': %v", tokens[0], err)
			return
		}
		pieces := []int{}
		for _, s := range strings.Split(tokens[1], " ") {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Could not process '%v': %v", s, err)
				return
			}
			pieces = append(pieces, i)
		}
		results := [][]int{}
		layer := []int{pieces[0]}
		results = append(results, layer)
		for i := 1; i < len(pieces); i++ {
			layer = []int{}
			for _, x := range results[i-1] {
				concat, err := strconv.Atoi(fmt.Sprintf("%v%v", x, pieces[i]))
				if err != nil {
					fmt.Printf("Could not concat %v with %v: %v", x, pieces[i], err)
					return
				}
				layer = append(layer, x+pieces[i], x*pieces[i], concat)
			}
			results = append(results, layer)
		}
		for _, x := range results[len(results)-1] {
			if x == goal {
				total += goal
				break
			}
		}
	}

	fmt.Printf("Part 2: %v", total)
}
