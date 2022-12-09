package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part 1

func onPerimeter(forrest [][]int, x, y int) bool {
	return x == 0 ||
		x == width(forrest)-1 ||
		y == 0 ||
		y == height(forrest)-1
}

func fromSouth(forrest [][]int, x, y int) bool {
	yOffsets := createRange(y+1, height(forrest))
	return all(yOffsets, func(i int) bool {
		return forrest[i][x] < forrest[y][x]
	})
}

func fromEast(forrest [][]int, x, y int) bool {
	xOffsets := createRange(x+1, width(forrest))
	return all(xOffsets, func(i int) bool {
		return forrest[y][i] < forrest[y][x]
	})
}

func fromNorth(forrest [][]int, x, y int) bool {
	yOffsets := reverse(createRange(0, y))
	return all(yOffsets, func(i int) bool {
		return forrest[i][x] < forrest[y][x]
	})
}

func fromWest(forrest [][]int, x, y int) bool {
	xOffsets := reverse(createRange(0, x))
	return all(xOffsets, func(i int) bool {
		return forrest[y][i] < forrest[y][x]
	})
}

func isVisible(forrest [][]int, x, y int) bool {
	return fromNorth(forrest, x, y) ||
		fromEast(forrest, x, y) ||
		fromSouth(forrest, x, y) ||
		fromWest(forrest, x, y)
}

func countVisibleTrees(forrest [][]int) (count int) {
	forEachTree(forrest, func(x, y int) {
		if onPerimeter(forrest, x, y) || isVisible(forrest, x, y) {
			count++
		}
	})
	return count
}

// Part 2

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
		// Check after increment; a blocking tree is visible
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

func calculateBestScore(forrest [][]int) (best int) {
	forEachTree(forrest, func(x, y int) {
		best = max(best, scenicScore(forrest, x, y))
	})
	return best
}

// Answers

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
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
	fmt.Println(countVisibleTrees(forrest))
	fmt.Println(calculateBestScore(forrest))
}
