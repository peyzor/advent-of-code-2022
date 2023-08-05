package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day3Part1() {
	priorities := make(map[rune]int)

	for i := 1; i <= 26; i++ {
		char := rune('a' - 1 + i)
		priorities[char] = i
	}

	for i := 27; i <= 52; i++ {
		char := rune('A' - 27 + i)
		priorities[char] = i
	}

	relativePath := "Day3/input.txt"

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	inputPath := filepath.Join(currentDir, relativePath)

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line)%2 != 0 {
			log.Fatal("bad entry")
		}
		middle := len(line) / 2
		comp1 := line[:middle]
		comp2 := line[middle:]
		r := findCommonChar(comp1, comp2)
		total += priorities[r]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func findCommonChar(str1, str2 string) rune {
	var r rune
	exist := make(map[rune]bool)

	for _, char := range str1 {
		exist[char] = true
	}

	for _, char := range str2 {
		if exist[char] {
			r = char
			break
		}
	}

	return r
}
