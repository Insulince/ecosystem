package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Fox struct {
	Location location.Location
}

func (f Fox) Value() int {
	panic("implement me")
}

func (f Fox) Name() string {
	return "Fox"
}

func (f Fox) Symbol() string {
	return "F"
}

func (f Fox) Place(l location.Location) grid.Card {
	f.Location = l
	return f
}

func (f Fox) At() location.Location {
	return f.Location
}
