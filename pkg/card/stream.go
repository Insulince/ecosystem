package card

type Stream struct {
}

func (s Stream) Value() int {
	panic("implement me")
}

func (s Stream) Name() string {
	return "Stream"
}

func (s Stream) Symbol() string {
	return "S"
}
