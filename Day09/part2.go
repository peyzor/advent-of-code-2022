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

func day9part2() {
	relativePath := "Day09/input.txt"

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

	knots := make([]point, 10)
	visitedByTail := make(map[point]bool)
	visitedByTail[knots[9]] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		direction := rune(line[0])
		moves, _ := strconv.Atoi(line[2:])

		for moves > 0 {
			switch direction {
			case 'U':
				knots[0].y++
			case 'R':
				knots[0].x++
			case 'D':
				knots[0].y--
			case 'L':
				knots[0].x--
			}
			// each knot is the head of the successive knot
			for i := 0; i < len(knots)-1; i++ {
				knots[i+1] = adjustTailExtended(knots[i+1], knots[i])
			}
			moves--
			visitedByTail[knots[9]] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(visitedByTail))
}

func adjustTailExtended(tail point, head point) (newTail point) {
	newTail = tail
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, 1}, point{-1, 2}, point{0, 2}, point{1, 2}, point{2, 1}, point{2, 2}, point{-2, 2}:
		newTail.y++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{1, 2}, point{2, 1}, point{2, 0}, point{2, -1}, point{1, -2}, point{2, 2}, point{2, -2}:
		newTail.x++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, -2}, point{2, -1}, point{1, -2}, point{0, -2}, point{-1, -2}, point{-2, -1}, point{2, -2}:
		newTail.y--
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, -2}, point{-1, -2}, point{-2, -1}, point{-2, -0}, point{-2, 1}, point{-1, 2}, point{-2, 2}:
		newTail.x--
	}
	return
}
