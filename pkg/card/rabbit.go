package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Rabbit struct {
	Location location.Location
}

func (r Rabbit) Value() int {
	panic("implement me")
}

func (r Rabbit) Name() string {
	return "Rabbit"
}

func (r Rabbit) Symbol() string {
	return "R"
}

func (r Rabbit) Place(l location.Location) grid.Card {
	r.Location = l
	return r
}

func (r Rabbit) At() location.Location {
	return r.Location
}
