package main

import (
	"fmt"

	"github.com/noahssarcastic/advent2022/debug"
)

const (
	debugStandard = iota
	debugCheckpoint
	debugPaths
)

type debugger struct {
	mode   int
	bounds debug.BBox
	main   *debug.Canvas
	path   *debug.Canvas
}

func newDebugger(sim *simulation, debugMode int) *debugger {
	d := debugger{mode: debugMode}
	if debugMode >= debugStandard {
		fmt.Printf("Ground set at %v.\n", sim.ground)

		d.bounds = *bounds(sim)
		d.main = debug.NewCanvas(600, 600)
		d.main.Draw(SAND_SPAWN.X(), SAND_SPAWN.Y(), '=')
		drawRocks(sim, d.main)
		fmt.Println("Initial state:")
		d.main.Print(d.bounds)
	}
	if debugMode >= debugPaths {
		d.path = d.main.Copy()
	}
	return &d
}

func bounds(sim *simulation) *debug.BBox {
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

func (db *debugger) step(sim *simulation) {
	if db.mode >= debugPaths {
		db.path.Draw(sim.curr.X(), sim.curr.Y(), '~')
		db.path.Print(db.bounds)
		debug.Pause()
	}
}

func (db *debugger) placeSand(sim *simulation) {
	if db.mode >= debugStandard {
		db.main.Draw(sim.curr.X(), sim.curr.Y(), 'o')
	}
	if db.mode >= debugCheckpoint {
		db.main.Print(db.bounds)
		debug.Pause()
	}
	if db.mode >= debugPaths {
		db.path = db.main.Copy()
	}
}

func (db *debugger) final() {
	if db.mode >= debugStandard {
		fmt.Println("Final state:")
		db.main.Print(db.bounds)
	}
}
