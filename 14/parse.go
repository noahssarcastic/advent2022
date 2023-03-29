package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noahssarcastic/advent2022/14/line"
	"github.com/noahssarcastic/advent2022/14/point"
)

func parse(s string) (lines []line.Line) {
	pts := strings.Split(strings.ReplaceAll(s, " ", ""), "->")
	var prev *point.Point = nil
	for i := range pts {
		curr := parsePoint(pts[i])
		if prev != nil {
			lines = append(lines, *line.New(*prev, *curr))
		}
		prev = curr
	}
	return lines
}

func parsePoint(s string) *point.Point {
	pair := strings.Split(s, ",")
	first, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(fmt.Errorf("failed to parse point.Point from '%s': %w", s, err))
	}
	second, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(fmt.Errorf("failed to parse point.Point from '%s': %w", s, err))
	}
	return point.New(first, second)
}
