package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func onPerimeter(forrest [][]int, x, y int) bool {
	return x == 0 ||
		x == width(forrest)-1 ||
		y == 0 ||
		y == height(forrest)-1
}

func fromSouth(forrest [][]int, x, y int) bool {
	return every(
		createRange(y+1, height(forrest)),
		func(i int) bool {
			return forrest[i][x] < forrest[y][x]
		})
}

func fromEast(forrest [][]int, x, y int) bool {
	return every(
		createRange(x+1, width(forrest)),
		func(i int) bool {
			return forrest[y][i] < forrest[y][x]
		})
}

func fromNorth(forrest [][]int, x, y int) bool {
	return every(
		reverse(createRange(0, y)),
		func(i int) bool {
			return forrest[i][x] < forrest[y][x]
		})
}

func fromWest(forrest [][]int, x, y int) bool {
	return every(
		reverse(createRange(0, x)),
		func(i int) bool {
			return forrest[y][i] < forrest[y][x]
		})
}

func isVisible(forrest [][]int, x, y int) bool {
	return fromNorth(forrest, x, y) ||
		fromEast(forrest, x, y) ||
		fromSouth(forrest, x, y) ||
		fromWest(forrest, x, y)
}

func scoreNorth(forrest [][]int, x, y int) (score int) {
	for i := y - 1; i >= 0; i-- {
		score++
		if forrest[i][x] >= forrest[y][x] {
			break
		}
	}
	return score
}

func scoreEast(forrest [][]int, x, y int) (score int) {
	for i := x + 1; i < width(forrest); i++ {
		score++
		if forrest[y][i] >= forrest[y][x] {
			break
		}
	}
	return score
}

func scoreSouth(forrest [][]int, x, y int) (score int) {
	for i := y + 1; i < height(forrest); i++ {
		score++
		if forrest[i][x] >= forrest[y][x] {
			break
		}
	}
	return score
}

func scoreWest(forrest [][]int, x, y int) (score int) {
	for i := x - 1; i >= 0; i-- {
		score++
		if forrest[y][i] >= forrest[y][x] {
			break
		}
	}
	return score
}

func scenicScore(forrest [][]int, x, y int) int {
	return scoreNorth(forrest, x, y) *
		scoreEast(forrest, x, y) *
		scoreSouth(forrest, x, y) *
		scoreWest(forrest, x, y)
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Parse
	forrest := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, h := range line {
			treeHeight, _ := strconv.Atoi(string(h))
			row = append(row, treeHeight)
		}
		forrest = append(forrest, row)
	}

	// Count trees visible from outside forrest
	visible := 0
	forEachTree(forrest, func(x, y int) {
		if onPerimeter(forrest, x, y) || isVisible(forrest, x, y) {
			visible++
		}
	})
	fmt.Println(visible == 1700)

	// Calculate scenic score
	bestScore := 0
	forEachTree(forrest, func(x, y int) {
		bestScore = max(bestScore, scenicScore(forrest, x, y))
	})
	fmt.Println(bestScore == 470596)
}
