package util

type Point struct {
	X, Y int
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

type Direction struct {
	Dx, Dy int
}

var UP Direction = Direction{0, -1}
var DOWN Direction = Direction{0, 1}
var LEFT Direction = Direction{-1, 0}
var RIGHT Direction = Direction{1, 0}

type PointWithDirection struct {
	Point Point
	Dir   Direction
}

func (pd PointWithDirection) RotateRight() PointWithDirection {
	d := pd.Dir
	switch pd.Dir {
	case UP:
		d = RIGHT
	case DOWN:
		d = LEFT
	case LEFT:
		d = UP
	case RIGHT:
		d = DOWN
	}
	return PointWithDirection{pd.Point, d}
}

func (pd PointWithDirection) MoveInDirection() PointWithDirection {
	return PointWithDirection{Point{pd.Point.X + pd.Dir.Dx, pd.Point.Y + pd.Dir.Dy}, pd.Dir}
}
