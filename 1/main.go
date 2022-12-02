package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Cals float64 // The number of total calories in an elf's inventory.
}

// Create a new Elf with an empty inventory.
func NewElf() Elf {
	return Elf{0}
}

// Implement sort.Interface for []Elf.
type ByCals []Elf

func (a ByCals) Len() int           { return len(a) }
func (a ByCals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCals) Less(i, j int) bool { return a[i].Cals < a[j].Cals }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scan(scanner *bufio.Scanner) []Elf {
	elves := make([]Elf, 1)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, NewElf())
			i++
		} else {
			cal, err := strconv.ParseFloat(line, 64)
			check(err)
			elves[i].Cals += cal
		}
	}
	check(scanner.Err())
	return elves
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	elves := scan(scanner)

	sort.Sort(ByCals(elves))

	fmt.Printf("The elf with the most calories has %v kcal.\n", elves[len(elves)-1].Cals)

	topThree := elves[len(elves)-1].Cals +
		elves[len(elves)-2].Cals +
		elves[len(elves)-3].Cals
	fmt.Printf("The top three elves with the most calories have %v kcal total.\n", topThree)
}
