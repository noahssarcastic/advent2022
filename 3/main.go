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

func priority(pack string) int {
	first := pack[:len(pack)/2]
	second := pack[len(pack)/2:]
	for _, c := range first {
		if strings.ContainsRune(second, c) {
			if unicode.IsLower(c) {
				return int(c-'a') + 1
			} else {
				return int(c-'A') + 27
			}
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
		line := scanner.Text()
		total += priority(line)
	}
	check(scanner.Err())
	fmt.Printf("The sum of the priorities of the misplaced items is %v.", total)
}
