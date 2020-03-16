package location

import "fmt"

type Location struct {
	X int
	Y int
}

func New(x int, y int) Location {
	return Location{
		X: x,
		Y: y,
	}
}

func (l Location) Coords() string {
	return fmt.Sprintf("%v,%v", l.X, l.Y)
}

func (l Location) Up() Location {
	return Location{
		X: l.X,
		Y: l.Y - 1,
	}
}

func (l Location) Right() Location {
	return Location{
		X: l.X + 1,
		Y: l.Y,
	}
}

func (l Location) Down() Location {
	return Location{
		X: l.X,
		Y: l.Y + 1,
	}
}

func (l Location) Left() Location {
	return Location{
		X: l.X - 1,
		Y: l.Y,
	}
}
