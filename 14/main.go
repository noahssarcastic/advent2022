package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/noahssarcastic/advent2022/14/line"
	"github.com/noahssarcastic/advent2022/14/point"
)

var SAND_SPAWN = *point.New(500, 0)

func main() {
	input := flag.String("f", "input.txt", "Input file to run simulation on")
	v0 := flag.Bool("v", false, "Print sand state to console.")
	v1 := flag.Bool("vv", false, "Pause between each step.")
	v2 := flag.Bool("vvv", false, "Show sand paths.")
	flag.Parse()
	debugMode := -1
	if *v2 {
		debugMode = debugPaths
	} else if *v1 {
		debugMode = debugCheckpoint
	} else if *v0 {
		debugMode = debugStandard
	}

	f, err := os.Open(*input)
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

	sim := newSimulation(lines)
	db := newDebugger(sim, debugMode)
	for {
		err := move(sim)
		cantMove := errors.Is(err, ErrCantMove)
		if cantMove && point.Equal(sim.curr, SAND_SPAWN) {
			sim.sand = append(sim.sand, sim.curr)
			break
		} else if cantMove {
			db.placeSand(sim)
			sim.sand = append(sim.sand, sim.curr)
			sim.curr = SAND_SPAWN
		} else if err != nil {
			panic(err)
		} else {
			db.step(sim)
		}
	}
	db.final()
	fmt.Printf("There are %d grains of sand.\n", len(sim.sand))
}

type simulation struct {
	lines  []line.Line
	sand   []point.Point
	curr   point.Point
	ground int
}

func newSimulation(lines []line.Line) *simulation {
	sim := simulation{lines: lines, curr: SAND_SPAWN}
	for _, ln := range sim.lines {
		for _, pt := range ln.Endpoints() {
			if pt.Y()+2 > sim.ground {
				sim.ground = pt.Y() + 2
			}
		}
	}
	return &sim
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
		if !hitRocks(mv, sim.lines) &&
			!hitSand(mv, sim.sand) &&
			mv.Y() != sim.ground {
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
