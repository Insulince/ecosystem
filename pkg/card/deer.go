package card

type Deer struct {
}

func (d Deer) Value() int {
	panic("implement me")
}

func (d Deer) Name() string {
	return "Deer"
}

func (d Deer) Symbol() string {
	return "D"
}
