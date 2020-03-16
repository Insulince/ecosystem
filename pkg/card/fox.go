package card

type Fox struct {
}

func (f Fox) Value() int {
	panic("implement me")
}

func (f Fox) Name() string {
	return "Fox"
}

func (f Fox) Symbol() string {
	return "F"
}
