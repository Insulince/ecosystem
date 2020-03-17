package main

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/grid/ecosystem"
)

func main() {
	m := `
	W M M B M
	D M B M M
	F M M B M
	E S T b M
`
	e := ecosystem.From(m)

	fmt.Println(e.Symbol())

	_, err := e.Calculate()
	if err != nil {
		panic(err)
	}

	fmt.Println(e.DumpScores())
}
