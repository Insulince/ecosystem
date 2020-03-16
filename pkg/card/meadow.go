package card

type Meadow struct {
}

func (m Meadow) Value() int {
	panic("implement me")
}

func (m Meadow) Name() string {
	return "Meadow"
}

func (m Meadow) Symbol() string {
	return "M"
}
