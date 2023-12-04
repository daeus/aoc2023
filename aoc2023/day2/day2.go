package main

import (
	"fmt"
	"regexp"

	"example.com/helper"
)

type CubeSet struct {
	red   int
	blue  int
	green int
}

func main() {
	linesPtr := helper.ReadFileToArray("input.txt")
	fmt.Println(PartOne(linesPtr))
	fmt.Println(PartTwo(linesPtr))
}

func PartOne(linesPtr *[]string) int {
	lines := *linesPtr
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	res := 0

	for _, line := range lines {
		gameNumber, cubeSets := GetGamesFromLine(line)
		isPossible := true
		for _, cubeSet := range cubeSets {
			if cubeSet.red > redLimit || cubeSet.green > greenLimit || cubeSet.blue > blueLimit {
				isPossible = false
			}
		}
		if isPossible {
			res += gameNumber
		}
	}
	return res
}

func PartTwo(linesPtr *[]string) int {
	res := 0
	lines := *linesPtr
	for _, line := range lines {
		minSet := CubeSet{0, 0, 0}
		_, cubeSets := GetGamesFromLine(line)
		for _, cubeSet := range cubeSets {
			minSet.red = max(minSet.red, cubeSet.red)
			minSet.green = max(minSet.green, cubeSet.green)
			minSet.blue = max(minSet.blue, cubeSet.blue)
		}
		res += minSet.blue * minSet.green * minSet.red
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetGamesFromLine(line string) (int, []CubeSet) {
	r := regexp.MustCompile(`(Game )|(: )|(; *)`)
	a := r.Split(line, -1)
	cubeSets := []CubeSet{}
	for _, g := range a[2:] {
		rg := regexp.MustCompile(`, | `)
		rawSet := rg.Split(g, -1)

		var set CubeSet
		for i := 0; i < len(rawSet)-1; i++ {
			if rawSet[i+1] == "red" {
				set.red = helper.ToInt(rawSet[i])
			}
			if rawSet[i+1] == "blue" {
				set.blue = helper.ToInt(rawSet[i])
			}
			if rawSet[i+1] == "green" {
				set.green = helper.ToInt(rawSet[i])
			}
		}
		cubeSets = append(cubeSets, set)

	}

	return helper.ToInt(a[1]), cubeSets
}
