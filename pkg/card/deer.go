package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Deer struct {
	Location location.Location
}

func (d Deer) Value() int {
	panic("implement me")
}

func (d Deer) Name() string {
	return "Deer"
}

func (d Deer) Symbol() string {
	return "D"
}

func (d Deer) Place(l location.Location) grid.Card {
	d.Location = l
	return d
}

func (d Deer) At() location.Location {
	return d.Location
}
