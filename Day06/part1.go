package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day6part1() {
	relativePath := "Day6/input.txt"

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

	var track []rune
	tracker := make(map[rune]bool)
	scanner := bufio.NewScanner(file)

	var sb strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sb.WriteString(line)
	}

	for i, c := range sb.String() {
		_, ok := tracker[c]
		if ok {
			index := IndexOf(track, c)
			for i := 0; i <= index; i++ {
				delete(tracker, track[i])
			}
			track = track[index+1:]
		}
		tracker[c] = true
		track = append(track, c)
		if len(tracker) == 4 {
			fmt.Printf("idx: %d\n", i)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
