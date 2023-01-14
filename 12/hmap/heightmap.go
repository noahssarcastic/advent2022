package hmap

import "github.com/noahssarcastic/advent2022/12/coord"

type Heightmap [][]byte

func New() Heightmap {
	return make(Heightmap, 0)
}

func (hm Heightmap) Width() int {
	return len(hm[0])
}

func (hm Heightmap) Height() int {
	return len(hm)
}

func (hm Heightmap) Get(c coord.Coord) int {
	return int(hm[c.Y()][c.X()])
}

func (hm Heightmap) InBounds(c coord.Coord) bool {
	return (c.X() >= 0 &&
		c.X() < hm.Width() &&
		c.Y() >= 0 &&
		c.Y() < hm.Height())
}

func (hm Heightmap) Grade(current, next coord.Coord) int {
	return hm.Get(next) - hm.Get(current)
}

func (hm Heightmap) IsTraversable(current, next coord.Coord) bool {
	return hm.Grade(current, next) <= 1
}

func (hm Heightmap) Adjacent(c coord.Coord) []coord.Coord {
	return []coord.Coord{
		coord.Add(c, coord.New(-1, 0)),
		coord.Add(c, coord.New(1, 0)),
		coord.Add(c, coord.New(0, -1)),
		coord.Add(c, coord.New(0, 1)),
	}
}
