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

type point struct {
	x, y int
}

func day9part1() {
	relativePath := "Day9/input.txt"

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

	head := point{0, 0}
	tail := point{0, 0}
	visitedByTail := make(map[point]bool)
	visitedByTail[tail] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		direction := rune(line[0])
		moves, _ := strconv.Atoi(line[2:])

		for moves > 0 {
			switch direction {
			case 'U':
				head.y++
			case 'R':
				head.x++
			case 'D':
				head.y--
			case 'L':
				head.x--
			}
			moves--
			tail = adjustTail(head, tail)
			visitedByTail[tail] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("visited: %+v\n", visitedByTail)
	fmt.Println(len(visitedByTail))
}

func adjustTail(head, tail point) point {
	newTail := tail

	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, 1}, point{-1, 2}, point{0, 2}, point{1, 2}, point{2, 1}:
		newTail.y++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{1, 2}, point{2, 1}, point{2, 0}, point{2, -1}, point{1, -2}:
		newTail.x++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{2, -1}, point{1, -2}, point{0, -2}, point{-1, -2}, point{-2, -1}:
		newTail.y--
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-1, -2}, point{-2, -1}, point{-2, 0}, point{-2, 1}, point{-1, 2}:
		newTail.x--
	}

	return newTail
}
