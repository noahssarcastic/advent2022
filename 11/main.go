package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	rounds = 10000
)

var superMod int = 1

type Monkey struct {
	items                        []int
	operation, opValue           string
	divisibleBy, ifTrue, ifFalse int
	inspections                  int
}

func newMonkey(
	items []int,
	operation, opValue string,
	divisibleBy, ifTrue, ifFalse int,
) *Monkey {
	return &Monkey{
		items:       items,
		operation:   operation,
		opValue:     opValue,
		divisibleBy: divisibleBy,
		ifTrue:      ifTrue,
		ifFalse:     ifFalse,
		inspections: 0,
	}
}

func (m *Monkey) inspect(worry int) int {
	var opValue int
	if m.opValue == "old" {
		opValue = worry
	} else {
		opValue, _ = strconv.Atoi(m.opValue)
	}

	if m.operation == "*" {
		return worry * opValue
	} else if m.operation == "+" {
		return worry + opValue
	}
	panic("invalid operation")
}

func (m *Monkey) throw() (item int, throwTo int) {
	item = m.dequeue()
	item = m.inspect(item)
	item %= superMod
	if item < 0 {
		item = 0
	}
	if item%m.divisibleBy == 0 {
		throwTo = m.ifTrue
	} else {
		throwTo = m.ifFalse
	}
	m.inspections++
	return item, throwTo
}

func (m *Monkey) dequeue() (item int) {
	item = m.items[0]
	m.items = m.items[1:]
	return item
}

func parseStartingItems(line string) []int {
	itemText := strings.Split(line, ":")[1]
	itemText = strings.ReplaceAll(itemText, " ", "")

	startingItems := make([]int, 0)
	for _, el := range strings.Split(itemText, ",") {
		item, _ := strconv.Atoi(el)
		startingItems = append(startingItems, item)
	}
	return startingItems
}

var operationRegex = regexp.MustCompile(`\s*Operation: new = old ([\*\+]) (\d+|old)`)

func parseOperation(line string) (operation string, opValue string) {
	opExpression := operationRegex.FindStringSubmatch(line)
	operation = opExpression[1]
	opValue = opExpression[2]
	return operation, opValue
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	testRegex, _ := regexp.Compile(`\s*Test: divisible by (\d+)`)
	trueRegex, _ := regexp.Compile(`\s*If true: throw to monkey (\d+)`)
	falseRegex, _ := regexp.Compile(`\s*If false: throw to monkey (\d+)`)
	ms := make([]Monkey, 0)
	for scanner.Scan() {
		lines := make([]string, 0, 6)
		for i := 0; i < 6; i++ {
			lines = append(lines, scanner.Text())
			scanner.Scan()
		}

		startingItems := parseStartingItems(lines[1])
		operation, opValue := parseOperation(lines[2])
		divisibleBy, _ := strconv.Atoi(testRegex.FindStringSubmatch(lines[3])[1])
		superMod *= divisibleBy
		ifTrue, _ := strconv.Atoi(trueRegex.FindStringSubmatch(lines[4])[1])
		ifFalse, _ := strconv.Atoi(falseRegex.FindStringSubmatch(lines[5])[1])
		ms = append(ms, *newMonkey(
			startingItems,
			operation,
			opValue,
			divisibleBy,
			ifTrue,
			ifFalse,
		))
	}

	for round := 0; round < rounds; round++ {
		for i := range ms {
			for len(ms[i].items) > 0 {
				item, throwTo := ms[i].throw()
				ms[throwTo].items = append(ms[throwTo].items, item)
			}
		}
	}

	inspections := make([]int, len(ms))
	for i := range ms {
		inspections[i] = ms[i].inspections
	}
	sort.Ints(inspections)

	monkeyBusiness := inspections[len(inspections)-2] * inspections[len(inspections)-1]
	fmt.Println(monkeyBusiness)
}
