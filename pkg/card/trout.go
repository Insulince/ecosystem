package card

type Trout struct {
}

func (t Trout) Value() int {
	panic("implement me")
}

func (t Trout) Name() string {
	return "Trout"
}

func (t Trout) Symbol() string {
	return "T"
}
