package card

type Dragonfly struct {
}

func (d Dragonfly) Value() int {
	panic("implement me")
}

func (d Dragonfly) Name() string {
	return "Dragonfly"
}

func (d Dragonfly) Symbol() string {
	return "d"
}
