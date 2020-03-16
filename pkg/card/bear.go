package card

type Bear struct {
}

func (b Bear) Value() int {
	panic("implement me")
}

func (b Bear) Name() string {
	return "Bear"
}

func (b Bear) Symbol() string {
	return "b"
}
