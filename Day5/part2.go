package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day5part2() {
	relativePath := "Day5/input.txt"

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

			popped := make([]rune, 0)
			for i := 0; i < moveCount; i++ {
				popped = append(popped, stacks[from-1].Pop())
			}
			for i := len(popped) - 1; i >= 0; i-- {
				stacks[to-1].Push(popped[i])
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
