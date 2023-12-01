package main

import (
	"testing"

	"example.com/helper"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142
	actual := PartOne(&input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	// input := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// }

	linesPtr := helper.ReadFileToArray("input.txt")

	expected := 281
	actual := PartTwo(linesPtr)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
