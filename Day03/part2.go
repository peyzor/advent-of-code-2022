package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day3part2() {
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

	group := make([]string, 0, 3)
	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		group = append(group, line)
		if len(group) == 3 {
			r := findCommonCharThree(group[0], group[1], group[2])
			total += priorities[r]
			group = make([]string, 0, 3)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func findCommonCharTwo(str1, str2 string) []rune {
	runes := make([]rune, 0)
	exist := make(map[rune]bool)

	for _, char := range str1 {
		exist[char] = true
	}

	for _, char := range str2 {
		if exist[char] {
			runes = append(runes, char)
		}
	}

	return runes
}

func findCommonCharThree(str1, str2, str3 string) rune {
	var r rune
	exist := make(map[rune]bool)

	commons := findCommonCharTwo(str1, str2)
	for _, char := range commons {
		exist[char] = true
	}

	for _, char := range str3 {
		if exist[char] {
			r = char
			break
		}
	}

	return r
}
