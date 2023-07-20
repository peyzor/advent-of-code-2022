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

func day1Part1() {
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

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

	var maxCal int64
	for _, cal := range cals {
		if cal > maxCal {
			maxCal = cal
		}
	}

	fmt.Println(maxCal)

}
