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

func (hm Heightmap) Get(c *coord.Coord) int {
	return int(hm[c.Y()][c.X()])
}
