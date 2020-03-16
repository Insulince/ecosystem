package card

import "github.com/Insulince/ecosystem/pkg/grid"

func From(symbol string) grid.Card {
	switch symbol {
	case "b":
		return Bear{}
	case "B":
		return Bee{}
	case "D":
		return Deer{}
	case "d":
		return Dragonfly{}
	case "E":
		return Eagle{}
	case "F":
		return Fox{}
	case "M":
		return Meadow{}
	case "R":
		return Rabbit{}
	case "S":
		return Stream{}
	case "T":
		return Trout{}
	case "-":
		return Vacant{}
	case "W":
		return Wolf{}
	default:
		panic("cannot create card, invalid symbol provided: \"" + symbol + "\"")
	}
}
