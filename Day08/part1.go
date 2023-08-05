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

func day8part1() {
	relativePath := "Day08/input.txt"

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

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var row []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(invisibleCount(grid))
}

func transposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	tMatrix := make([][]int, cols)
	for i := range matrix {
		tMatrix[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tMatrix[j][i] = matrix[i][j]
		}
	}

	return tMatrix
}

func invisibleCount(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	tGrid := transposeMatrix(grid)

	var count int
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if !isInvisible(grid[j], tGrid[i], i, j) {
				count++
			}
		}
	}

	return 2*n + 2*(m-2) + count
}

func isInvisible(row, col []int, i, j int) bool {
	rowElem := row[i]
	rowLeft, rowRight := row[:i], row[i+1:]
	colElem := col[j]
	colTop, colDown := col[:j], col[j+1:]

	invisibility := [...]bool{false, false, false, false}
	for _, e := range rowLeft {
		if e >= rowElem {
			invisibility[0] = true
			break
		}
	}
	for _, e := range rowRight {
		if e >= rowElem {
			invisibility[1] = true
			break
		}
	}

	for _, e := range colTop {
		if e >= colElem {
			invisibility[2] = true
			break
		}
	}
	for _, e := range colDown {
		if e >= colElem {
			invisibility[3] = true
			break
		}
	}

	for _, i := range invisibility {
		if !i {
			return false
		}
	}

	return true
}
