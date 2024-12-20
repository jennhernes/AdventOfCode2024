package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2024/Day14/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not open %v: %v\n", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	robots := []Robot{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		pvsplit := strings.Split(line, " ")
		pTokens := strings.Split(pvsplit[0], "=")
		pxy := strings.Split(pTokens[1], ",")
		px, err := strconv.Atoi(pxy[0])
		if err != nil {
			fmt.Printf("Could not convert '%v': %v\n", pxy[0], err)
			return
		}
		py, err := strconv.Atoi(pxy[1])
		if err != nil {
			fmt.Printf("Could not convert '%v': %v\n", pxy[1], err)
			return
		}
		vTokens := strings.Split(pvsplit[1], "=")
		vxy := strings.Split(vTokens[1], ",")
		vx, err := strconv.Atoi(vxy[0])
		if err != nil {
			fmt.Printf("Could not convert '%v': %v\n", vxy[0], err)
			return
		}
		vy, err := strconv.Atoi(vxy[1])
		if err != nil {
			fmt.Printf("Could not convert '%v': %v\n", vxy[1], err)
			return
		}
		robots = append(robots, Robot{px, py, vx, vy})
	}

	xMax := 101
	yMax := 103
	quadrants := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	for _, robot := range robots {
		for i := 0; i < 100; i++ {
			robot.px = (robot.px + robot.vx + xMax) % xMax
			robot.py = (robot.py + robot.vy + yMax) % yMax
		}
		if robot.py < yMax/2 {
			if robot.px < xMax/2 {
				quadrants[0]++
			} else if robot.px > xMax/2 {
				quadrants[1]++
			}
		} else if robot.py > yMax/2 {
			if robot.px < xMax/2 {
				quadrants[2]++
			} else if robot.px > xMax/2 {
				quadrants[3]++
			}
		}
	}
	total := 1
	for _, n := range quadrants {
		total *= n
	}

	fmt.Printf("Part 1: %v\n", total)
}

type Robot struct {
	px int
	py int
	vx int
	vy int
}
