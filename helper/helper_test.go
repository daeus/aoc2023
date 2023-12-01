package helper

import (
	"testing"
)

func TestReadFileToArray(t *testing.T) {
	actual := ReadFileToArray("test.txt")
	expected := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
