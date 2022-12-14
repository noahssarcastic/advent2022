package main

import (
	"fmt"
	"strconv"
)

func printBoardState(knots []Knot) {
	visits := map[string]int{}
	for i, k := range knots {
		visits[getKey(&k)] = i
	}
	for r := 15; r >= 0; r-- {
		row := make([]byte, 0)
		for c := 0; c < 28; c++ {
			if c-15 == 0 && r-5 == 0 {
				row = append(row, 'S')
			} else if i, check := visits[getKey(&Knot{float64(c - 15), float64(r - 5)})]; check {
				id := strconv.Itoa(i)
				row = append(row, id[0])
			} else {
				row = append(row, '.')
			}
		}
		fmt.Println(string(row))
	}
}

func printPositions(head, tail *Knot) {
	fmt.Printf("H(%v,%v) T(%v,%v)\n", head.x, head.y, tail.x, tail.y)
}
