package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func priority(item rune) int {
	if unicode.IsLower(item) {
		return int(item-'a') + 1
	} else {
		return int(item-'A') + 27
	}
}

func findError(pack string) int {
	first := pack[:len(pack)/2]
	second := pack[len(pack)/2:]
	for _, r := range first {
		if strings.ContainsRune(second, r) {
			return priority(r)
		}
	}
	return -1
}

func findBadge(first, second, third string) int {
	for _, r := range first {
		if strings.ContainsRune(second, r) &&
			strings.ContainsRune(third, r) {
			return priority(r)
		}
	}
	return -1
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()
		total += findBadge(first, second, third)
	}
	check(scanner.Err())
	fmt.Printf("The sum of the priorities of the team badges is %v.", total)
}
