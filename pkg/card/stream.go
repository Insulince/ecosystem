package card

import (
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
)

type Stream struct {
	Location location.Location
}

func (s Stream) Value() int {
	panic("implement me")
}

func (s Stream) Name() string {
	return "Stream"
}

func (s Stream) Symbol() string {
	return "S"
}

func (s Stream) Place(l location.Location) grid.Card {
	s.Location = l
	return s
}

func (s Stream) At() location.Location {
	return s.Location
}
