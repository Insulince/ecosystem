package card

type Eagle struct {
}

func (e Eagle) Value() int {
	panic("implement me")
}

func (e Eagle) Name() string {
	return "Eagle"
}

func (e Eagle) Symbol() string {
	return "E"
}
