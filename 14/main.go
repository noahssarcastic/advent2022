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

var SandSpawn = *point.New(500, 0)

var (
	input = flag.String("f", "input.txt", "run simulation on given input file")
	v0    = flag.Bool("v", false, "print sand state to console")
	v1    = flag.Bool("vv", false, "pause between each step")
	v2    = flag.Bool("vvv", false, "show sand paths")
)

func main() {
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	var lines []line.Line
	for scanner.Scan() {
		lines = append(lines, parse(scanner.Text())...)
	}
	if scanner.Err() != nil {
		panic(err)
	}

	sim := newSimulation(lines)
	db := newDebugger(sim)
	for {
		err := move(sim)
		cantMove := errors.Is(err, ErrCantMove)
		if cantMove && point.Equal(sim.curr, SandSpawn) {
			sim.sand = append(sim.sand, sim.curr)
			break
		} else if cantMove {
			db.placeSand(sim)
			sim.sand = append(sim.sand, sim.curr)
			sim.setHit(sim.curr)
			sim.curr = SandSpawn
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
	hitMap map[point.Point]struct{}
}

func (sim *simulation) hit(pt point.Point) bool {
	_, found := sim.hitMap[pt]
	return found || pt.Y() == sim.ground
}

func (sim *simulation) setHit(pt point.Point) {
	sim.hitMap[pt] = struct{}{}
}

func newSimulation(lines []line.Line) *simulation {
	sim := simulation{
		lines:  lines,
		curr:   SandSpawn,
		hitMap: make(map[point.Point]struct{}),
	}
	for _, ln := range sim.lines {
		for _, pt := range ln.Endpoints() {
			if pt.Y()+2 > sim.ground {
				sim.ground = pt.Y() + 2
			}
		}
		for _, pt := range ln.Along() {
			sim.setHit(pt)
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
		if !sim.hit(mv) {
			sim.curr = mv
			return nil
		}
	}
	return ErrCantMove
}
