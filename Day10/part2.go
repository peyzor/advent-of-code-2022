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

func day10part2() {
	relativePath := "Day10/input.txt"

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

	registerX, cycleCount := 1, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		instruction := strings.Fields(line)
		cycleControl(&cycleCount, &registerX)
		if instruction[0] == "addx" {
			value, _ := strconv.Atoi(instruction[1])
			cycleControl(&cycleCount, &registerX)
			registerX += value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func cycleControl(cycleNumber, registerX *int) {
	if (*cycleNumber)%40 == 0 && *cycleNumber <= 220 {
		fmt.Println()
	}
	if *registerX-1 == *cycleNumber%40 || *registerX == *cycleNumber%40 || *registerX+1 == *cycleNumber%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*cycleNumber++
}
