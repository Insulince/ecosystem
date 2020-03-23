package grid

import "github.com/Insulince/ecosystem/pkg/location"

type Card interface {
	Value() int
	Name() string
	Symbol() string
	Place(location.Location) Card
	At() location.Location
}
