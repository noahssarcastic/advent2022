package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSectionRange(s string) (int, int) {
	splitString := strings.Split(s, "-")
	start, _ := strconv.Atoi(splitString[0])
	end, _ := strconv.Atoi(splitString[1])
	return start, end
}

func doesContain(s string) bool {
	splitString := strings.Split(s, ",")
	firstStart, firstEnd := getSectionRange(splitString[0])
	secondStart, secondEnd := getSectionRange(splitString[1])
	return (firstStart <= secondStart && secondEnd <= firstEnd) ||
		(secondStart <= firstStart && firstEnd <= secondEnd)
}

func doesOverlap(s string) bool {
	splitString := strings.Split(s, ",")
	firstStart, firstEnd := getSectionRange(splitString[0])
	secondStart, secondEnd := getSectionRange(splitString[1])
	for i := firstStart; i <= firstEnd; i++ {
		for j := secondStart; j <= secondEnd; j++ {
			if i == j {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("input2.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	containCount := 0
	overlapCount := 0
	for scanner.Scan() {
		if doesContain(scanner.Text()) {
			containCount++
		}
		if doesOverlap(scanner.Text()) {
			overlapCount++
		}
	}
	check(scanner.Err())
	fmt.Printf("Teams where a section is fully contained in another: %v\n", containCount)
	fmt.Printf("Teams where sections overlap: %v\n", overlapCount)
}
