package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Dragonfly struct {
	Location location.Location
}

func (d Dragonfly) Value() int {
	panic("implement me")
}

func (d Dragonfly) Name() string {
	return "Dragonfly"
}

func (d Dragonfly) Symbol() string {
	return "d"
}

func (d Dragonfly) Place(l location.Location) grid.Card {
	d.Location = l
	return d
}

func (d Dragonfly) At() location.Location {
	return d.Location
}
