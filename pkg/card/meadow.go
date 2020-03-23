package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Meadow struct {
	Location location.Location
}

func (m Meadow) Value() int {
	panic("implement me")
}

func (m Meadow) Name() string {
	return "Meadow"
}

func (m Meadow) Symbol() string {
	return "M"
}

func (m Meadow) Place(l location.Location) grid.Card {
	m.Location = l
	return m
}

func (m Meadow) At() location.Location {
	return m.Location
}
