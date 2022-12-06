package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Helper functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func moveOneByOne(num int, from, to *[]byte) {
	bs := (*from)[len((*from))-num:]
	*from = (*from)[:len((*from))-num]
	for i := len(bs) - 1; i >= 0; i-- {
		*to = append(*to, bs[i])
	}
}

func moveAllAtOnce(num int, from, to *[]byte) {
	bs := (*from)[len((*from))-num:]
	*from = (*from)[:len((*from))-num]
	*to = append(*to, bs...)
}

// Parse

func generateStacks(numStacks int, startState []string) [][]byte {
	stacks := make([][]byte, numStacks)
	for line := len(startState) - 1; line >= 0; line-- {
		for i := 0; i < numStacks; i++ {
			crate := startState[line][4*i+1]
			if crate != ' ' {
				stacks[i] = append(stacks[i], crate)
			}
		}
	}
	return stacks
}

func parseState(scanner *bufio.Scanner) [][]byte {
	matchLabel := regexp.MustCompile(`^\s*(?:\d+\s+)*(\d+)`)
	startState := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if matchLabel.MatchString(line) {
			match := matchLabel.FindStringSubmatch(line)
			numStacks := parseInt(match[1])
			// Skip blank line
			scanner.Scan()
			return generateStacks(numStacks, startState)
		} else {
			startState = append(startState, line)
		}
	}
	return nil
}

func parseProcedure(scanner *bufio.Scanner) (procedure []string) {
	for scanner.Scan() {
		procedure = append(procedure, scanner.Text())
	}
	return procedure
}

func parseInput(scanner *bufio.Scanner) ([][]byte, []string) {
	state := parseState(scanner)
	procedure := parseProcedure(scanner)
	check(scanner.Err())
	return state, procedure
}

func parseInstruction(s string) (int, int, int) {
	r, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	check(err)
	match := r.FindStringSubmatch(s)
	return parseInt(match[1]), parseInt(match[2]), parseInt(match[3])
}

// Run

func runProcedure(stacks [][]byte, procedure []string) {
	for _, line := range procedure {
		num, from, to := parseInstruction(line)
		moveAllAtOnce(num, &stacks[from-1], &stacks[to-1])
	}
}

func main() {
	f, err := os.Open("input2.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	stacks, procedure := parseInput(scanner)

	runProcedure(stacks, procedure)

	out := make([]byte, 0)
	for _, stack := range stacks {
		out = append(out, stack[len(stack)-1])
	}
	fmt.Printf("The crates on top of each stack are \"%v\"", string(out))
}
