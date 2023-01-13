package main

import (
	"fmt"
	"math"

	"github.com/noahssarcastic/advent2022/12/coord"
	"github.com/noahssarcastic/advent2022/12/hmap"
)

type Map struct {
	hm         hmap.Heightmap
	start, end *coord.Coord
}

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
	return hm.Get(next) - hm.Get(current)
}

func recurse(m *Map, current *coord.Coord, history []coord.Coord, moveCount int) int {
	if coord.Equal(current, m.end) {
		return moveCount
	}
	quickest := math.MaxInt
	for _, move := range getMoves(m.hm, current) {
		isRepeat := coord.Any(history, &move)
		if getGrade(m.hm, current, &move) > 1 || isRepeat {
			continue
		}
		pathLength := recurse(m, &move, append(history, move), moveCount+1)
		if pathLength < quickest {
			quickest = pathLength
		}
	}
	return quickest
}

func main() {
	inputFile := parseArgs()
	m := parseInput(inputFile)
	current := m.start
	history := []coord.Coord{*current}
	quickestPath := recurse(m, current, history, 0)
	fmt.Printf("The quickest path takes %v steps.\n", quickestPath)
}
