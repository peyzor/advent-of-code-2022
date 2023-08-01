package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func day10part1() {
	relativePath := "Day10/input.txt"

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
	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	var cycleCount int
	xregister := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		instruction := strings.Fields(line)
		if instruction[0] == "addx" {
			cycleCount++
			if inSlice(interestingCycles, cycleCount) {
				total += cycleCount * xregister
			}
			amount, err := strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatal(err)
			}
			xregister += amount
			cycleCount++
			if inSlice(interestingCycles, cycleCount) {
				total += cycleCount * xregister
			}
		} else if instruction[0] == "noop" {
			cycleCount++
			if inSlice(interestingCycles, cycleCount) {
				total += cycleCount * xregister
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func inSlice(cycles []int, cycle int) bool {
	for _, c := range cycles {
		if cycle == c {
			return true
		}
	}

	return false
}
