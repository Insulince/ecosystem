package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Eagle struct {
	Location location.Location
}

func (e Eagle) Value() int {
	panic("implement me")
}

func (e Eagle) Name() string {
	return "Eagle"
}

func (e Eagle) Symbol() string {
	return "E"
}

func (e Eagle) Place(l location.Location) grid.Card {
	e.Location = l
	return e
}

func (e Eagle) At() location.Location {
	return e.Location
}
