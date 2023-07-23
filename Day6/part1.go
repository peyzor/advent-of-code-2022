package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day6part1() {
	relativePath := "Day6/input.txt"

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

	var traverseCount int
	tracker := make(map[rune]bool)
	scanner := bufio.NewScanner(file)

	var sb strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sb.WriteString(line)
	}

	for i, c := range sb.String() {
		_, ok := tracker[c]
		if ok {
			tracker = make(map[rune]bool)
			continue
		}
		tracker[c] = true
		if len(tracker) == 4 {
			fmt.Printf("idx: %d\n", i)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("traverse count: %d\n", traverseCount)
	fmt.Printf("tracker: %+v\n", tracker)
}
