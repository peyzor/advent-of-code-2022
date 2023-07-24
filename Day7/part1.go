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

type Node struct {
	Name     string
	Size     int
	IsFile   bool
	Parent   *Node
	Children map[string]*Node
}

func day7part1() {
	relativePath := "Day7/input.txt"

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

	var totalSize int
	for _, dir := range dirs {
		size := calcSize(*dir)
		if size <= 100_000 {
			totalSize += size
		}
	}

	fmt.Println(totalSize)
}

func calcSize(node Node) int {
	var size int
	if node.IsFile {
		return node.Size
	}

	for _, d := range node.Children {
		size += calcSize(*d)
	}

	return size
}
