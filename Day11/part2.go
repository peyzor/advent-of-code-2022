package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func day11part2() {
	relativePath := "Day11/input.txt"

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

	var monkeys []*monkey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "Monkey") {
			var id int
			_, err := fmt.Sscanf(line, "Monkey %d:", &id)
			if err != nil {
				log.Fatal(err)
			}

			m := monkey{id: id}
			for line != "" {
				scanner.Scan()
				line = strings.TrimSpace(scanner.Text())
				if strings.HasPrefix(line, "Starting items: ") {
					m.items = extractItems(line)
				} else if strings.HasPrefix(line, "Operation") {
					m.operation = extractOperation(line)
				} else if strings.HasPrefix(line, "Test") {
					m.test = extractTest(line)
				} else if strings.HasPrefix(line, "If true") {
					m.test.trueTarget = extractIfTrue(line)
				} else if strings.HasPrefix(line, "If false") {
					m.test.falseTarget = extractIfFalse(line)
				}
			}

			monkeys = append(monkeys, &m)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	playMonkeys(monkeys, 10000, 1)

	var inspectionCounts []int
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.inspectionCount)
	}

	max1, max2 := findTwoLargest(inspectionCounts)
	fmt.Println(max1 * max2)
}
