package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Hello() {
	fmt.Println("Hello World")
}

func ReadFileToArray(f string) *[]string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return &lines
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
