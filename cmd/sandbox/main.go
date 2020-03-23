package main

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/card"
	"github.com/Insulince/ecosystem/pkg/grid/ecosystem"
	"github.com/Insulince/ecosystem/pkg/location"
	"log"
)

func main() {
	ecoMap := `
	W M M B M
	D M B M M
	F M M B M
	E S T b M
`

	eco := ecosystem.FromMap(ecoMap)

	fmt.Println(eco.Map())

	_, err := eco.Score()
	if err != nil {
		log.Fatalf("%v while calculating ecosystem score\n", err)
	}

	var b card.Bear = card.Bear{}

	b = b.Place(location.New(0, 0)).(card.Bear)

	fmt.Println(eco.DumpScores())
}
