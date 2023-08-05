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

type Range struct {
	start int64
	stop  int64
}

func day4part1() {
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
		if oneContains(r1, r2) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func oneContains(r1, r2 Range) bool {
	if r1.start <= r2.start && r1.stop >= r2.stop {
		return true
	} else if r2.start <= r1.start && r2.stop >= r1.stop {
		return true
	}

	return false
}

func parseRange(r string) (Range, error) {
	rg := strings.Split(r, "-")
	x, err := strconv.ParseInt(rg[0], 10, 64)
	if err != nil {
		return Range{}, err
	}
	y, err := strconv.ParseInt(rg[1], 10, 64)
	if err != nil {
		return Range{}, err
	}

	return Range{start: x, stop: y}, nil
}
