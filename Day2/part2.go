package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type choice struct {
	opp string
	res string
}

func day2Part2() {
	relativePath := "Day2/input.txt"

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
	}
	roundScores := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	var totalScore int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		entry := strings.Split(line, " ")
		c := choice{
			entry[0],
			entry[1],
		}
		p := getPick(c)
		cs, ok := choiceScores[p]
		if !ok {
			log.Fatal("entry not valid")
		}
		rs, ok := roundScores[c.res]
		if !ok {
			log.Fatal("entry not valid")
		}
		totalScore += cs + rs
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalScore)
}

func getPick(c choice) string {
	var pick string

	if c.opp == "A" && c.res == "X" {
		pick = "C"
	} else if c.opp == "A" && c.res == "Y" {
		pick = "A"
	} else if c.opp == "A" && c.res == "Z" {
		pick = "B"
	} else if c.opp == "B" && c.res == "X" {
		pick = "A"
	} else if c.opp == "B" && c.res == "Y" {
		pick = "B"
	} else if c.opp == "B" && c.res == "Z" {
		pick = "C"
	} else if c.opp == "C" && c.res == "X" {
		pick = "B"
	} else if c.opp == "C" && c.res == "Y" {
		pick = "C"
	} else if c.opp == "C" && c.res == "Z" {
		pick = "A"
	}

	return pick
}
