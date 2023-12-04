package main

import (
	"fmt"
	"regexp"
	"strings"

	"example.com/helper"
)

func main() {
	linesPtr := helper.ReadFileToArray("input.txt")
	fmt.Println(PartOne(linesPtr))
	fmt.Println(PartTwo(linesPtr))
}

func PartOne(linesPtr *[]string) int {
	res := 0
	for _, l := range *linesPtr {
		if l == "" {
			continue
		}
		pts, _ := getWinningPoints(l)
		res += pts
	}
	return res
}

func PartTwo(linesPtr *[]string) int {
	var noOfcards []int

	for i := 0; i < len(*linesPtr); i++ {
		noOfcards = append(noOfcards, 1)
	}

	for k, l := range *linesPtr {
		if _, c := getWinningPoints(l); c > 0 {
			for i := k + 1; i < k+1+c; i++ {
				noOfcards[i] += noOfcards[k]
			}
		}
	}
	res := 0
	for _, v := range noOfcards {
		res += v
	}
	return res
}

func getWinningPoints(line string) (int, int) {
	res := 0
	cards := []int{}

	win := make(map[int]bool)
	a := regexp.MustCompile(`(Card )|(: )|( \| )`).Split(line, -1)
	w := strings.Fields(a[2])
	c := strings.Fields(a[3])
	for _, v := range w {
		win[helper.ToInt(v)] = true
	}
	for _, v := range c {
		if win[helper.ToInt(v)] {
			if res == 0 {
				res = 1
			} else {
				res *= 2
			}
			cards = append(cards, helper.ToInt(v))
		}
	}

	return res, len(cards)
}
