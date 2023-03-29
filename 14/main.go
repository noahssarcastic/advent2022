package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func CopyPoint(pt Point) *Point {
	newPt := pt
	return &newPt
}

var SAND_SPAWN = Point{500, 0}

func (a *Point) Equal(b *Point) bool {
	return a.X == b.X && a.Y == b.Y
}

type Line struct {
	start, end Point
}

func (ln *Line) PointsAlong() (pts []Point) {
	if ln.start.X == ln.end.X {
		for i := min(ln.start.Y, ln.end.Y); i <= max(ln.start.Y, ln.end.Y); i++ {
			pts = append(pts, Point{ln.start.X, i})
		}
	} else if ln.start.Y == ln.end.Y {
		for i := min(ln.start.X, ln.end.X); i <= max(ln.start.X, ln.end.X); i++ {
			pts = append(pts, Point{i, ln.start.Y})
		}
	} else {
		panic(fmt.Errorf("line %v is not horizontal or vertical", ln))
	}
	return pts
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

func collision(ln *Line, pt *Point) bool {
	a := ln.start
	b := ln.end
	if a.X == b.X {
		return pt.X == a.X &&
			min(a.Y, b.Y) <= pt.Y &&
			pt.Y <= max(a.Y, b.Y)
	} else if a.Y == b.Y {
		return pt.Y == a.Y &&
			min(a.X, b.X) <= pt.X &&
			pt.X <= max(a.X, b.X)
	} else {
		panic(fmt.Errorf("line %v is not horizontal or vertical", ln))
	}
}

type Simulation struct {
	Lines []Line
	Sand  []Point
	Curr  Point
}

func NewSimulation(lines []Line) *Simulation {
	sim := Simulation{}
	sim.Lines = lines
	sim.ResetCurr()
	return &sim
}

func (s *Simulation) ResetCurr() {
	s.Curr = SAND_SPAWN
}

func extents(sim *Simulation) *BBox {
	bb := BBox{SAND_SPAWN.X, SAND_SPAWN.Y, SAND_SPAWN.X, SAND_SPAWN.Y}
	for _, line := range sim.Lines {
		for _, pt := range []Point{line.start, line.end} {
			if pt.X < bb.x0 {
				bb.x0 = pt.X
			}
			if pt.X > bb.x1 {
				bb.x1 = pt.X
			}
			if pt.Y < bb.y0 {
				bb.y0 = pt.Y
			}
			if pt.Y > bb.y1 {
				bb.y1 = pt.Y
			}
		}
	}
	return &bb
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []Line
	for scanner.Scan() {
		lines = append(lines, parse(scanner.Text())...)
	}
	if scanner.Err() != nil {
		panic(err)
	}

	sim := NewSimulation(lines)
	bb := extents(sim)
	canv := NewCanvas(600, 600)
	drawRocks(sim, canv)
	canv.Print(*bb)
	Pause()
	for tick := 0; tick < 10000; tick++ {
		if err := move(sim); errors.Is(err, ErrCantMove) {
			tick = 0
			sim.Sand = append(sim.Sand, sim.Curr)
			canv.Draw(sim.Curr.X, sim.Curr.Y, 'o')
			canv.Print(*bb)
			sim.ResetCurr()
			Pause()
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Println(len(sim.Sand))
}

func drawRocks(sim *Simulation, canv *Canvas) {
	for _, ln := range sim.Lines {
		for _, pt := range ln.PointsAlong() {
			canv.Draw(pt.X, pt.Y, '#')
		}
	}
}

var ErrCantMove = errors.New("sand blocked; can't move")

func move(sim *Simulation) error {
	pt := sim.Curr
	mvs := []Point{
		{pt.X, pt.Y + 1},
		{pt.X - 1, pt.Y + 1},
		{pt.X + 1, pt.Y + 1},
	}
	for _, mv := range mvs {
		if !hitRocks(&mv, sim.Lines) && !hitSand(&mv, sim.Sand) {
			sim.Curr = mv
			return nil
		}
	}
	return ErrCantMove
}

func hitRocks(pos *Point, lines []Line) bool {
	for _, rock := range lines {
		collided := collision(&rock, pos)
		if collided {
			return true
		}
	}
	return false
}

func hitSand(pos *Point, sand []Point) bool {
	for _, grain := range sand {
		if pos.Equal(&grain) {
			return true
		}
	}
	return false
}
