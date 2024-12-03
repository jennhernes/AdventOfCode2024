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

	mulRegex, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	if err != nil {
		fmt.Printf("Could not compile regex main: %v\n", err)
		return
	}
	numRegex, err := regexp.Compile("\\d{1,3}")
	if err != nil {
		fmt.Printf("Could not compile regex nums: %v\n", err)
		return
	}

	matches := mulRegex.FindAllString(string(contents), -1)

	total := 0
	for _, match := range matches {
		prod := 1
		//fmt.Println(match)
		sa := numRegex.FindAllString(match, 2)
		for _, s := range sa {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Tried to convert %v to int: %v\n", s, err)
				return
			}
			prod *= n
		}
		total += prod
	}

	fmt.Printf("Part 1: %v\n", total)

	doRegex, err := regexp.Compile("do\\(\\)")
	if err != nil {
		fmt.Printf("Could not compile regex do(): %v\n", err)
		return
	}

	dontRegex, err := regexp.Compile("don't\\(\\)")
	if err != nil {
		fmt.Printf("Could not compile regex do(): %v\n", err)
		return
	}

	dos := doRegex.FindAllStringIndex(string(contents), -1)
	donts := dontRegex.FindAllStringIndex(string(contents), -1)
	matchIndices := mulRegex.FindAllStringIndex(string(contents), -1)

	total = 0
	for i, match := range matches {
		doIndex := 0
		dontIndex := -1
		for _, a := range dos {
			//fmt.Printf("Index for %v is %v\n", match, matchIndices[i])
			if a[0] > matchIndices[i][0] {
				break
			}
			doIndex = a[0]
		}
		for _, a := range donts {
			if a[0] > matchIndices[i][0] {
				break
			}
			dontIndex = a[0]
		}
		if dontIndex > doIndex {
			continue
		}
		prod := 1
		//fmt.Println(match)
		sa := numRegex.FindAllString(match, 2)
		for _, s := range sa {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Tried to convert %v to int: %v\n", s, err)
				return
			}
			prod *= n
		}
		total += prod
	}

	fmt.Printf("Part 2: %v\n", total)
}
