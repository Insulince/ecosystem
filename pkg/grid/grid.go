package grid

import (
	"errors"
	"github.com/Insulince/ecosystem/pkg/location"
)

var (
	ErrCellAlreadyPopulated = errors.New("cell is already populated")
	ErrPartiallyFilledGrid  = errors.New("grid is only partially filled")
)

type Grid interface {
	Place(c Card, l location.Location)
	At(l location.Location) Card
	Adjacent(l location.Location) [4]Card
	DoubleAdjacent(l location.Location) [12]Card
	Calculate() (int, error)
	Symbol() string
	DumpScores() string
}
