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
	db := debugger{mode: debugMode}
	if debugMode >= debugStandard {
		fmt.Printf("Ground set at %v.\n", sim.ground)
		db.bounds = *bounds(sim)
		// TODO: don't hard code
		db.main = debug.NewCanvas(1000, 1000)
		db.main.Draw(SAND_SPAWN.X(), SAND_SPAWN.Y(), '=')
		db.drawRocks(sim)
		db.drawGround(sim)
		fmt.Println("Initial state:")
		db.main.Print(db.bounds)
	}
	if debugMode >= debugCheckpoint {
		debug.Pause()
	}
	if debugMode >= debugPaths {
		db.path = db.main.Copy()
	}
	return &db
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
	bb.Expand(bb.XMax(), bb.YMax()+2)
	// bb.Expand(bb.XMax()+(bb.YMax()-bb.YMin()), bb.YMax())
	// bb.Expand(bb.XMin()-(bb.YMax()-bb.YMin()), bb.YMax())
	return bb
}

func (db *debugger) drawRocks(sim *simulation) {
	for _, ln := range sim.lines {
		for _, pt := range ln.Along() {
			db.main.Draw(pt.X(), pt.Y(), '#')
		}
	}
}

func (db *debugger) drawGround(sim *simulation) {
	for i := db.bounds.XMin(); i <= db.bounds.XMax(); i++ {
		db.main.Draw(i, sim.ground, '#')
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
		db.bounds.Expand(sim.curr.X(), sim.curr.Y())
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
