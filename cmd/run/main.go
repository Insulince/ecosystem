package main

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/engine"
	"github.com/Insulince/ecosystem/pkg/grid/ecosystem"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	e := engine.New(engine.Config{
		Candidates: ecosystem.Candidates,
		Verbose:    true,
	})

	e.Run()

	fmt.Println(e.BestScore)
	fmt.Println()
	fmt.Println(e.Best.Map())
	fmt.Println()
	fmt.Println(e.Best.DumpScores())
}
