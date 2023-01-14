package coord

type Coord [2]int

func New(x, y int) Coord {
	return Coord{x, y}
}

func Equal(a, b Coord) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

func (c Coord) X() int { return c[0] }

func (c Coord) Y() int { return c[1] }

func Add(a, b Coord) Coord {
	return New(a.X()+b.X(), a.Y()+b.Y())
}
