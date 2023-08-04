package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type monkey struct {
	id              int
	items           []item
	operation       operation
	test            test
	inspectionCount int
}

type item struct {
	level int
}

type operation struct {
	operator string
	operand  int
}

type test struct {
	number      int
	trueTarget  int
	falseTarget int
}

func day11part1() {
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

	playMonkeys(monkeys, 20, 3)

	var inspectionCounts []int
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.inspectionCount)
	}

	max1, max2 := findTwoLargest(inspectionCounts)
	fmt.Println(max1 * max2)
}

func extractItems(str string) []item {
	var items []item

	str = strings.TrimPrefix(str, "Starting items: ")
	numbers := strings.Split(str, ",")
	for _, n := range numbers {
		level, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item{level: level})
	}

	return items
}

func extractOperation(str string) operation {
	str = strings.TrimPrefix(str, "Operation: new = old ")
	operatorOperand := strings.Split(str, " ")

	var operand int
	operator := operatorOperand[0]
	operandStr := operatorOperand[1]
	if operandStr == "old" {
		operator = "**"
		operand = 2

	} else {
		var err error
		operand, err = strconv.Atoi(operandStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	return operation{
		operator: operator,
		operand:  operand,
	}
}

func extractTest(str string) test {
	var t test

	var number int
	_, err := fmt.Sscanf(str, "Test: divisible by %d", &number)
	if err != nil {
		log.Fatal(err)
	}

	t.number = number
	return t
}

func extractIfTrue(str string) int {
	var target int

	_, err := fmt.Sscanf(str, "If true: throw to monkey %d", &target)
	if err != nil {
		log.Fatal(err)
	}

	return target
}

func extractIfFalse(str string) int {
	var target int

	_, err := fmt.Sscanf(str, "If false: throw to monkey %d", &target)
	if err != nil {
		log.Fatal(err)
	}

	return target
}

func findTwoLargest(numbers []int) (int, int) {
	if len(numbers) < 2 {
		return 0, 0
	}

	var largest, secondLargest int

	if numbers[0] > numbers[1] {
		largest = numbers[0]
		secondLargest = numbers[1]
	} else {
		largest = numbers[1]
		secondLargest = numbers[0]
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > largest {
			secondLargest = largest
			largest = numbers[i]
		} else if numbers[i] > secondLargest {
			secondLargest = numbers[i]
		}
	}

	return largest, secondLargest
}

func playMonkeys(monkeys []*monkey, rounds, divide int) {
	bigLimit := 1
	for _, m := range monkeys {
		bigLimit *= m.test.number
	}

	round := 1
	for round <= rounds {
		for midx, m := range monkeys {
			var indicesToDelete []int

			for idx, i := range m.items {
				var newLevel int
				switch m.operation.operator {
				case "*":
					newLevel = (i.level * m.operation.operand) / divide
				case "+":
					newLevel = (i.level + m.operation.operand) / divide
				case "**":
					newLevel = int(math.Pow(float64(i.level), float64(m.operation.operand))) / divide
				}

				newLevel = newLevel % bigLimit
				if newLevel%m.test.number == 0 {
					items := monkeys[m.test.trueTarget].items
					items = append(items, item{level: newLevel})
					monkeys[m.test.trueTarget].items = items
				} else {
					items := monkeys[m.test.falseTarget].items
					items = append(items, item{level: newLevel})
					monkeys[m.test.falseTarget].items = items
				}

				indicesToDelete = append(indicesToDelete, idx)
				m.inspectionCount++
			}

			monkeys[midx].items = []item{}
		}

		round++
	}
}

func deleteMultipleItems(slice []item, indicesToDelete []int) []item {
	// Create a map to store the indices of the items to delete
	toDelete := make(map[int]bool)
	for _, idx := range indicesToDelete {
		toDelete[idx] = true
	}

	// Create a new slice that will contain the remaining elements
	newSlice := make([]item, 0)

	// Iterate over the original slice, and only append the elements that are not marked for deletion
	for idx, item := range slice {
		if !toDelete[idx] {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
