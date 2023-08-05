package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		var r rune
		return r
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack) Peek() rune {
	if len(s.items) == 0 {
		var r rune
		return r
	}

	item := s.items[len(s.items)-1]
	return item
}

func (s *Stack) addToBottom(item rune) {
	s.items = append([]rune{item}, s.items...)
}

func day5part1() {
	relativePath := "Day05/input.txt"

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

	stacks := make([]Stack, 9)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for !strings.HasPrefix(scanner.Text(), " 1 ") {
		for i, c := range scanner.Text() {
			if c != ' ' && c != '[' && c != ']' {
				stacks[i/4].addToBottom(c)
			}
		}
		scanner.Scan()
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "move") {
			var moveCount, from, to int
			_, err := fmt.Sscanf(line, "move %d from %d to %d", &moveCount, &from, &to)
			if err != nil {
				return
			}
			for i := 0; i < moveCount; i++ {
				stacks[to-1].Push(stacks[from-1].Pop())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var res string
	for _, s := range stacks {
		res += string(s.Peek())
	}

	fmt.Println(res)
}
