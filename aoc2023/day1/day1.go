package main

import (
	"fmt"
	"regexp"

	"example.com/helper"
)

func main() {
	helper.Hello()
	linesPtr := helper.ReadFileToArray("input.txt")
	fmt.Println(PartOne(linesPtr))
	fmt.Println(PartTwo(linesPtr))
}

func PartOne(lines *[]string) int {
	total := 0
	for _, line := range *lines {
		var res []string
		for _, c := range line {
			if c >= '0' && c <= '9' {
				res = append(res, string(c))
			}
		}
		total += helper.ToInt(res[0] + res[len(res)-1])
	}
	return total
}

func PartTwo(lines *[]string) int {
	total := 0

	for _, line := range *lines {
		var res []string
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				res = append(res, string(c))
			} else {
				match, num, _ := isNum(line[i:])
				if match {
					res = append(res, fmt.Sprint(num))

					// cant skip as there is edge case eg. twone lol
					//i += skip
				}
			}
		}
		total += helper.ToInt(res[0] + res[len(res)-1])
	}
	return total
}

func isNum(str string) (bool, int, int) {
	ref := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for key, val := range ref {
		match, _ := regexp.MatchString("^"+key, str)
		if match {
			return true, val, len(key) - 1
		}
	}
	return false, 0, 0

}
