package card

type Bee struct {
}

func (b Bee) Value() int {
	panic("implement me")
}

func (b Bee) Name() string {
	return "Bee"
}

func (b Bee) Symbol() string {
	return "B"
}
