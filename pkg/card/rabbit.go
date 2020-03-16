package card

type Rabbit struct {
}

func (r Rabbit) Value() int {
	panic("implement me")
}

func (r Rabbit) Name() string {
	return "Rabbit"
}

func (r Rabbit) Symbol() string {
	return "R"
}
