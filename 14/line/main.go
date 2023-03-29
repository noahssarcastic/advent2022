package line

import (
	"fmt"

	"github.com/noahssarcastic/advent2022/14/point"
)

type Line struct {
	fst, snd point.Point
}

func (ln *Line) Endpoints() []point.Point {
	return []point.Point{ln.fst, ln.snd}
}

func New(fst, snd point.Point) *Line {
	return &Line{fst, snd}
}

type slopeError struct {
	ln Line
}

func (e *slopeError) Error() string {
	return fmt.Sprintf("line %v is not horizontal or vertical", e.ln)
}

// Calculate all discrete Points along the given Line.
func (ln *Line) Along() (pts []point.Point) {
	if ln.fst.X() == ln.snd.X() {
		for i := min(ln.fst.Y(), ln.snd.Y()); i <= max(ln.fst.Y(), ln.snd.Y()); i++ {
			pts = append(pts, *point.New(ln.fst.X(), i))
		}
	} else if ln.fst.Y() == ln.snd.Y() {
		for i := min(ln.fst.X(), ln.snd.X()); i <= max(ln.fst.X(), ln.snd.X()); i++ {
			pts = append(pts, *point.New(i, ln.fst.Y()))
		}
	} else {
		panic(slopeError{*ln})
	}
	return pts
}

// Return true if the given Point is on the given Line.
func On(ln Line, pt point.Point) bool {
	a := ln.fst
	b := ln.snd
	if a.X() == b.X() {
		return pt.X() == a.X() &&
			min(a.Y(), b.Y()) <= pt.Y() &&
			pt.Y() <= max(a.Y(), b.Y())
	} else if a.Y() == b.Y() {
		return pt.Y() == a.Y() &&
			min(a.X(), b.X()) <= pt.X() &&
			pt.X() <= max(a.X(), b.X())
	} else {
		panic(slopeError{ln})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
