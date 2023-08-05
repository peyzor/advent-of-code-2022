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

func day8part2() {
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

	fmt.Println(highestScenicScore(grid))
}

func highestScenicScore(grid [][]int) int {
	tGrid := transposeMatrix(grid)

	var scenicScores []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			s := calcScenicScore(grid[j], tGrid[i], i, j)
			scenicScores = append(scenicScores, s)
		}
	}

	var maxScore int
	for _, s := range scenicScores {
		if s > maxScore {
			maxScore = s
		}
	}

	return maxScore
}

func calcScenicScore(row, col []int, i, j int) int {
	rowElem := row[i]
	rowLeft, rowRight := row[:i], row[i+1:]
	colElem := col[j]
	colTop, colDown := col[:j], col[j+1:]
	visibility := [...]int{1, 1, 1, 1}

	var leftV int
	for k := len(rowLeft) - 1; k >= 0; k-- {
		leftV++
		if rowLeft[k] >= rowElem {
			break
		}

	}
	visibility[0] = leftV

	var rightV int
	for _, e := range rowRight {
		rightV++
		if e >= rowElem {
			break
		}
	}
	visibility[1] = rightV

	var topV int
	for k := len(colTop) - 1; k >= 0; k-- {
		topV++
		if colTop[k] >= colElem {
			break
		}
	}
	visibility[2] = topV

	var downV int
	for _, e := range colDown {
		downV++
		if e >= colElem {
			break
		}
	}
	visibility[3] = downV

	scenicScore := 1
	for _, v := range visibility {
		scenicScore *= v
	}

	return scenicScore
}
