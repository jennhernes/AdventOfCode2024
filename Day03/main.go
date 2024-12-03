package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day03/input.txt"
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open file %v: %v\n", filename, err)
		return
	}

	fullRegex, err := regexp.Compile("(mul\\((\\d{1,3}),(\\d{1,3})\\))")
	if err != nil {
		fmt.Printf("Could not compile regex full: %v\n", err)
		return
	}

	total := 0
	for _, match := range fullRegex.FindAllStringSubmatch(string(contents), -1) {
		//fmt.Printf("%v\n", match)
		x, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("Could not convert %v: %v\n", match[2], err)
			return
		}
		y, err := strconv.Atoi(match[3])
		if err != nil {
			fmt.Printf("Could not convert %v: %v\n", match[3], err)
			return
		}
		total += x * y
	}

	fmt.Printf("Part 1: %v\n", total)

	fullRegex, err = regexp.Compile("(mul\\((\\d{1,3}),(\\d{1,3})\\))|(do\\(\\))|(don't\\(\\))")
	if err != nil {
		fmt.Printf("Could not compile regex full: %v\n", err)
		return
	}
	enabled := true
	total = 0
	for _, match := range fullRegex.FindAllStringSubmatch(string(contents), -1) {
		//fmt.Printf("%v\n", match)
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		} else if enabled {
			x, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Printf("Could not convert %v: %v\n", match[2], err)
				return
			}
			y, err := strconv.Atoi(match[3])
			if err != nil {
				fmt.Printf("Could not convert %v: %v\n", match[3], err)
				return
			}
			total += x * y
		}
	}

	fmt.Printf("Part 2: %v\n", total)
}
