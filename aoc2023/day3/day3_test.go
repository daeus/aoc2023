package main

import (
	"testing"

	"example.com/helper"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361
	actual := PartOne(&input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	/* 	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	} */
	input := helper.ReadFileToArray("input.txt")

	expected := 467835
	actual := PartTwo(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
