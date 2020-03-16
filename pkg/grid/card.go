package grid

type Card interface {
	Value() int
	Name() string
	Symbol() string
}
