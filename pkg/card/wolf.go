package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Wolf struct {
	Location location.Location
}

func (w Wolf) Value() int {
	panic("implement me")
}

func (w Wolf) Name() string {
	return "Wolf"
}

func (w Wolf) Symbol() string {
	return "W"
}

func (w Wolf) Place(l location.Location) grid.Card {
	w.Location = l
	return w
}

func (w Wolf) At() location.Location {
	return w.Location
}
