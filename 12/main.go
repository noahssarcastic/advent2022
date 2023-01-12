package main

import (
	"fmt"
	"math"

	"github.com/noahssarcastic/advent2022/12/coord"
	"github.com/noahssarcastic/advent2022/12/hmap"
)

func getMoves(hm hmap.Heightmap, c *coord.Coord) (moves []coord.Coord) {
	if c.X() > 0 {
		moves = append(moves, *coord.Add(c, coord.New(-1, 0)))
	}
	if c.X() < hm.Width()-1 {
		moves = append(moves, *coord.Add(c, coord.New(1, 0)))
	}
	if c.Y() > 0 {
		moves = append(moves, *coord.Add(c, coord.New(0, -1)))
	}
	if c.Y() < hm.Height()-1 {
		moves = append(moves, *coord.Add(c, coord.New(0, 1)))
	}
	return moves
}

// Calculate the change in elevation between two coordinates
func getGrade(hm hmap.Heightmap, current, next *coord.Coord) int {
	return absInt(hm.Get(current) - hm.Get(next))
}

func isRepeat(c *coord.Coord, history []coord.Coord) bool {
	for _, h := range history {
		if coord.Equal(c, &h) {
			return true
		}
	}
	return false
}

func recurse(hm hmap.Heightmap, start, end, current *coord.Coord, history []coord.Coord, moveCount int) int {
	if coord.Equal(current, end) {
		return moveCount
	}
	quickest := math.MaxInt
	for _, m := range getMoves(hm, current) {
		if getGrade(hm, current, &m) > 1 || isRepeat(&m, history) {
			continue
		}
		pathLength := recurse(hm, start, end, &m, append(history, m), moveCount+1)
		if pathLength < quickest {
			quickest = pathLength
		}
	}
	return quickest
}

func main() {
	inputFile := parseArgs()
	hm, start, end := parseInput(inputFile)
	current := start
	history := []coord.Coord{*current}
	quickestPath := recurse(hm, start, end, current, history, 0)
	fmt.Printf("The quickest path takes %v steps.\n", quickestPath)
}
