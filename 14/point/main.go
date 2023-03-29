package point

type Point struct {
	x, y int
}

func (pt *Point) X() int {
	return pt.x
}

func (pt *Point) Y() int {
	return pt.y
}

func New(x, y int) *Point {
	return &Point{x, y}
}

func Copy(pt Point) *Point {
	newPt := pt
	return &newPt
}
func Equal(a, b Point) bool {
	return a.x == b.x && a.y == b.y
}
