package main

import "sort"

// Create a sequential int range between start (inclusive) and end (exclusive).
func createRange(start, end int) []int {
	newRange := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		newRange = append(newRange, i)
	}
	return newRange
}

func reverse(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return newSlice
}

func width(matrix [][]int) int {
	return len(matrix[0])
}

func height(matrix [][]int) int {
	return len(matrix)
}

func every(slice []int, check func(i int) bool) bool {
	for _, i := range slice {
		if !check(i) {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func forEachTree(forrest [][]int, do func(x, y int)) {
	for y, row := range forrest {
		for x := range row {
			do(x, y)
		}
	}
}
