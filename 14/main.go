package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/noahssarcastic/advent2022/14/line"
	"github.com/noahssarcastic/advent2022/14/point"
	"github.com/noahssarcastic/advent2022/debug"
)

var SAND_SPAWN = *point.New(500, 0)

type simulation struct {
	lines []line.Line
	sand  []point.Point
	curr  point.Point
}

func main() {
	debugMode := flag.Bool("debug", false, "a bool")
	flag.Parse()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []line.Line
	for scanner.Scan() {
		lines = append(lines, parse(scanner.Text())...)
	}
	if scanner.Err() != nil {
		panic(err)
	}
	sim := simulation{lines: lines, curr: SAND_SPAWN}

	var bb *debug.BBox
	var canv *debug.Canvas
	var pathCanv *debug.Canvas
	if *debugMode {
		bb = extents(&sim)
		canv = debug.NewCanvas(600, 600)
		canv.Draw(SAND_SPAWN.X(), SAND_SPAWN.Y(), '=')
		drawRocks(&sim, canv)
		pathCanv = canv.Copy()
		canv.Print(*bb)
		debug.Pause()
	}

	for tick := 0; tick < 10000; tick++ {
		if err := move(&sim); errors.Is(err, ErrCantMove) {
			if *debugMode {
				canv.Draw(sim.curr.X(), sim.curr.Y(), 'o')
				pathCanv = canv.Copy()
				canv.Print(*bb)
				debug.Pause()
			}
			sim.sand = append(sim.sand, sim.curr)
			sim.curr = SAND_SPAWN
			tick = 0
		} else if err != nil {
			panic(err)
		}
		if *debugMode {
			pathCanv.Draw(sim.curr.X(), sim.curr.Y(), '~')
			pathCanv.Print(*bb)
			debug.Pause()
		}
	}
	fmt.Printf("There are %d grains of sand.\n", len(sim.sand))
}

var ErrCantMove = errors.New("sand blocked; can't move")

func move(sim *simulation) error {
	pt := sim.curr
	mvs := []point.Point{
		*point.New(pt.X(), pt.Y()+1),
		*point.New(pt.X()-1, pt.Y()+1),
		*point.New(pt.X()+1, pt.Y()+1),
	}
	for _, mv := range mvs {
		if !hitRocks(mv, sim.lines) && !hitSand(mv, sim.sand) {
			sim.curr = mv
			return nil
		}
	}
	return ErrCantMove
}

func hitRocks(pos point.Point, lines []line.Line) bool {
	for _, rock := range lines {
		collided := line.On(rock, pos)
		if collided {
			return true
		}
	}
	return false
}

func hitSand(pos point.Point, sand []point.Point) bool {
	for _, grain := range sand {
		if point.Equal(pos, grain) {
			return true
		}
	}
	return false
}

func hitGround(pos point.Point) bool {
	return false
}

func extents(sim *simulation) *debug.BBox {
	bb := debug.Bounds(
		SAND_SPAWN.X(), SAND_SPAWN.Y(),
		SAND_SPAWN.X(), SAND_SPAWN.Y(),
	)
	for _, ln := range sim.lines {
		for _, pt := range ln.Endpoints() {
			bb.Expand(pt.X(), pt.Y())
		}
	}
	return bb
}

func drawRocks(sim *simulation, canv *debug.Canvas) {
	for _, ln := range sim.lines {
		for _, pt := range ln.Along() {
			canv.Draw(pt.X(), pt.Y(), '#')
		}
	}
}
