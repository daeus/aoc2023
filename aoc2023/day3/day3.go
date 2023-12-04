package main

import (
	"fmt"

	"example.com/helper"
)

func main() {
	linesPtr := helper.ReadFileToArray("input.txt")
	fmt.Println(PartOne(linesPtr))
	fmt.Println(PartTwo(linesPtr))
}

func PartOne(linesPtr *[]string) int {
	lines := *linesPtr

	res := 0

	for i := 0; i < len(lines); i++ {
		currentLinePtr := &lines[i]
		isNumber := false
		start := -1
		end := -1
		currentValue := ""

		for j := 0; j < len(*currentLinePtr); j++ {
			c := (*currentLinePtr)[j]
			if c >= '0' && c <= '9' {
				if !isNumber {
					start = j
					isNumber = true
				}
				currentValue += string(c)

				// if last char
				if j == len(*currentLinePtr)-1 {
					end = j
					if isAdjSymbol(lines, i, start, end) {
						res += helper.ToInt(currentValue)
					}
					currentValue = ""
					start = -1
					end = -1
					isNumber = false
				}
			} else {
				if isNumber {
					end = j - 1
					if isAdjSymbol(lines, i, start, end) {
						res += helper.ToInt(currentValue)
					}
					currentValue = ""
					start = -1
					end = -1
					isNumber = false
				}
			}
		}
	}
	return res
}

func PartTwo(linesPtr *[]string) int {
	lines := *linesPtr

	res := 0

	gearMap := make(map[string][]int)

	for i := 0; i < len(lines); i++ {
		currentLinePtr := &lines[i]
		isNumber := false
		start := -1
		end := -1
		currentValue := ""

		for j := 0; j < len(*currentLinePtr); j++ {
			c := (*currentLinePtr)[j]
			if c >= '0' && c <= '9' {
				if !isNumber {
					start = j
					isNumber = true
				}
				currentValue += string(c)

				// if last char
				if j == len(*currentLinePtr)-1 {
					end = j
					hasGear, gears := isAdjGear(lines, i, start, end)
					if hasGear {
						for _, gear := range gears {
							gearMap[gear] = append(gearMap[gear], helper.ToInt(currentValue))
						}
					}
					currentValue = ""
					start = -1
					end = -1
					isNumber = false
				}
			} else {
				if isNumber {
					end = j - 1
					hasGear, gears := isAdjGear(lines, i, start, end)
					if hasGear {
						for _, gear := range gears {
							gearMap[gear] = append(gearMap[gear], helper.ToInt(currentValue))
						}
					}
					currentValue = ""
					start = -1
					end = -1
					isNumber = false
				}
			}
		}
	}

	for _, g := range gearMap {
		if len(g) == 2 {
			res += g[0] * g[1]
		}
	}
	return res
}

func isAdjGear(lines []string, i int, start int, end int) (bool, []string) {
	searchStart := start
	searchEnd := end
	hasGear := false
	gears := []string{}
	if start > 0 {
		searchStart = start - 1
	}

	if end < len(lines[i])-1 {
		searchEnd = end + 1
	}

	if isSymbol(string(lines[i][searchStart])) {
		hasGear = true
		gears = append(gears, fmt.Sprint(i)+"_"+fmt.Sprint(searchStart))
	}

	if isSymbol(string(lines[i][searchEnd])) {
		hasGear = true
		gears = append(gears, fmt.Sprint(i)+"_"+fmt.Sprint(searchEnd))

	}

	// search upper line
	if i > 0 {
		searchLine := lines[i-1]
		for k, c := range searchLine[searchStart : searchEnd+1] {
			if isGear(string(c)) {
				hasGear = true
				gears = append(gears, fmt.Sprint(i-1)+"_"+fmt.Sprint(k+searchStart))
			}
		}
	}

	// search lower line
	if i < len(lines)-1 {
		searchLine := lines[i+1]
		for k, c := range searchLine[searchStart : searchEnd+1] {
			if isGear(string(c)) {
				hasGear = true
				gears = append(gears, fmt.Sprint(i+1)+"_"+fmt.Sprint(k+searchStart))
			}
		}
	}
	return hasGear, gears
}

func isGear(c string) bool {
	return c == "*"
}

func isAdjSymbol(lines []string, i int, start int, end int) bool {
	searchStart := start
	searchEnd := end
	if start > 0 {
		searchStart = start - 1
	}

	if end < len(lines[i])-1 {
		searchEnd = end + 1
	}

	if isSymbol(string(lines[i][searchStart])) || isSymbol(string(lines[i][searchEnd])) {
		return true
	}

	// search upper line
	if i > 0 {
		for _, c := range lines[i-1][searchStart : searchEnd+1] {
			if isSymbol(string(c)) {
				return true
			}
		}
	}

	// search lower line
	if i < len(lines)-1 {
		for _, c := range lines[i+1][searchStart : searchEnd+1] {
			if isSymbol(string(c)) {
				return true
			}
		}
	}
	return false
}

func isSymbol(c string) bool {
	symbols := []string{"@", "#", "$", "%", "-", "/", "*", "=", "*", "+", "&"}
	for _, s := range symbols {
		if s == c {
			return true
		}
	}
	return false
}
