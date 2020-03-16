package card

type Vacant struct {
}

func (v Vacant) Value() int {
	panic("implement me")
}

func (v Vacant) Name() string {
	return "Vacant"
}

func (v Vacant) Symbol() string {
	return "-"
}
