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

func day1Part2() {
	relativePath := "Day1/input.txt"

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	inputPath := filepath.Join(currentDir, relativePath)

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cals []int64
	var elfCal int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			cals = append(cals, elfCal)
			elfCal = 0
			continue
		}

		cal, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		elfCal += cal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	top3Elves := make([]int64, 3)
	for i := 0; i < 3; i += 1 {
		max, maxIdx := findMax(cals)
		top3Elves = append(top3Elves, max)
		cals = append(cals[:maxIdx], cals[maxIdx+1:]...)
	}

	var top3ElvesSum int64
	for _, cal := range top3Elves {
		top3ElvesSum += cal
	}

	fmt.Println(top3ElvesSum)
}

func findMax(elements []int64) (int64, int) {
	var max int64
	var maxIdx int
	for idx, e := range elements {
		if e > max {
			max = e
			maxIdx = idx
		}
	}

	return max, maxIdx
}
