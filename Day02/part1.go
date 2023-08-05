package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type input struct {
	opp string
	me  string
}

func day2Part1() {
	relativePath := "Day02/input.txt"

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

	choiceScores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var totalScore int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		entry := strings.Split(line, " ")
		if len(entry) != 2 {
			log.Fatal("entry not valid")
		}
		i := input{
			opp: entry[0],
			me:  entry[1],
		}
		cs, ok := choiceScores[i.me]
		if !ok {
			log.Fatal("choice not valid")
		}
		totalScore += getRoundScore(i) + cs
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalScore)
}

func getRoundScore(inp input) int {
	var score int

	if inp.opp == "A" && inp.me == "X" {
		score = 3
	} else if inp.opp == "A" && inp.me == "Y" {
		score = 6
	} else if inp.opp == "A" && inp.me == "Z" {
		score = 0
	} else if inp.opp == "B" && inp.me == "X" {
		score = 0
	} else if inp.opp == "B" && inp.me == "Y" {
		score = 3
	} else if inp.opp == "B" && inp.me == "Z" {
		score = 6
	} else if inp.opp == "C" && inp.me == "X" {
		score = 6
	} else if inp.opp == "C" && inp.me == "Y" {
		score = 0
	} else if inp.opp == "C" && inp.me == "Z" {
		score = 3
	}

	return score
}
