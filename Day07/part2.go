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

func day7part2() {
	relativePath := "Day07/input.txt"

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

	var rootDir *Node
	var currDir *Node
	var dirs []*Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input := strings.Fields(line)

		if len(input) > 2 {
			var destination string
			_, err := fmt.Sscanf(line, "$ cd %s", &destination)
			if err != nil {
				log.Fatal(err)
			}
			if destination == ".." {
				currDir = currDir.Parent

			} else if destination == "/" {
				currDir = &Node{
					Name:     "/",
					Size:     0,
					IsFile:   false,
					Parent:   nil,
					Children: make(map[string]*Node),
				}
				rootDir = currDir

			} else {
				currDir = currDir.Children[destination]
			}

		} else if input[0] == "dir" {
			dirName := input[1]

			currDir.Children[dirName] = &Node{
				Name:     dirName,
				Size:     0,
				IsFile:   false,
				Parent:   currDir,
				Children: make(map[string]*Node),
			}
			dirs = append(dirs, currDir.Children[dirName])

		} else if input[0] != "$" {
			size, _ := strconv.Atoi(input[0])
			filename := input[1]

			currDir.Children[filename] = &Node{
				Name:     filename,
				Size:     size,
				IsFile:   true,
				Parent:   currDir,
				Children: nil,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rootSize := calcSize(*rootDir)
	totalAvailable := 70_000_000
	needed := 30_000_000
	availableNow := totalAvailable - rootSize
	needToDelete := needed - availableNow

	var candidates []*Node
	for _, dir := range dirs {
		size := calcSize(*dir)
		dir.Size = size
		if size >= needToDelete {
			candidates = append(candidates, dir)
		}
	}

	smallestCandidate := candidates[0]
	for _, c := range candidates {
		if c.Size < smallestCandidate.Size {
			smallestCandidate = c
		}
	}

	fmt.Printf("%d\n", smallestCandidate.Size)
}
