package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day4part2() {
	relativePath := "Day04/input.txt"

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

	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		twoRanges := strings.Split(line, ",")
		if len(twoRanges) != 2 {
			log.Fatal("bad entry")
		}
		r1, err := parseRange(twoRanges[0])
		if err != nil {
			log.Fatal(err)
		}
		r2, err := parseRange(twoRanges[1])
		if err != nil {
			log.Fatal(err)
		}
		if doOverlap(r1, r2) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func doOverlap(r1, r2 Range) bool {
	if r1.stop < r2.start || r2.stop < r1.start {
		return false
	}

	return true
}
