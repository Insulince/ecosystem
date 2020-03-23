package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Trout struct {
	Location location.Location
}

func (t Trout) Value() int {
	panic("implement me")
}

func (t Trout) Name() string {
	return "Trout"
}

func (t Trout) Symbol() string {
	return "T"
}

func (t Trout) Place(l location.Location) grid.Card {
	t.Location = l
	return t
}

func (t Trout) At() location.Location {
	return t.Location
}
