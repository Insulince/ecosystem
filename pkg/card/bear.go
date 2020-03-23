package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Bear struct {
	Location location.Location
}

func (b Bear) Value() int {
	panic("implement me")
}

func (b Bear) Name() string {
	return "Bear"
}

func (b Bear) Symbol() string {
	return "b"
}

func (b Bear) Place(l location.Location) grid.Card {
	b.Location = l
	return b
}

func (b Bear) At() location.Location {
	return b.Location
}
