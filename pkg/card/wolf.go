package card

type Wolf struct {
}

func (w Wolf) Value() int {
	panic("implement me")
}

func (w Wolf) Name() string {
	return "Wolf"
}

func (w Wolf) Symbol() string {
	return "W"
}
