package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Bee struct {
	Location location.Location
}

func (b Bee) Value() int {
	panic("implement me")
}

func (b Bee) Name() string {
	return "Bee"
}

func (b Bee) Symbol() string {
	return "B"
}

func (b Bee) Place(l location.Location) grid.Card {
	b.Location = l
	return b
}

func (b Bee) At() location.Location {
	return b.Location
}
