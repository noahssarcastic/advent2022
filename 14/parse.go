package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(s string) (lines []Line) {
	pts := strings.Split(strings.ReplaceAll(s, " ", ""), "->")
	var prev *Point = nil
	for i := range pts {
		curr := parsePoint(pts[i])
		if prev != nil {
			lines = append(lines, Line{*prev, *curr})
		}
		prev = curr
	}
	return lines
}

func parsePoint(s string) *Point {
	pair := strings.Split(s, ",")
	first, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(fmt.Errorf("failed to parse Point from '%s': %w", s, err))
	}
	second, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(fmt.Errorf("failed to parse Point from '%s': %w", s, err))
	}
	return &Point{first, second}
}
