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

	eng := engine.New(ecosystem.Candidates)

	t := time.Now()
	for i := 0; i < 1000000000; i++ {
		if i%10000000 == 0 {
			fmt.Printf("--- %v - %v (%v)\n", i, eng.BestScore, time.Since(t))
		} else if i%1000000 == 0 {
			fmt.Printf("------ %v - %v (%v)\n", i, eng.BestScore, time.Since(t))
		} else if i%100000 == 0 {
			fmt.Printf("--------- %v - %v (%v)\n", i, eng.BestScore, time.Since(t))
		} else if i%25000 == 0 {
			fmt.Printf("------------ %v - %v\n", i, eng.BestScore)
		}

		eng.NextEcosystem()

		e := ecosystem.From(eng.Current)

		v, err := e.Calculate()
		if err != nil {
			panic(err) // TODO: Handle.
		}

		if v >= eng.BestScore {
			if v > eng.BestScore {
				fmt.Println("=== NEW BEST ===")
			}
			eng.Best = e
			eng.BestScore = v
			fmt.Printf("=== %v === (@ %v)\n", eng.BestScore, i)
			fmt.Println(eng.Best.Symbol())
			fmt.Println()
			fmt.Println(eng.Best.DumpScores())
			fmt.Printf("==========\n")
		}
	}
	fmt.Println(time.Since(t))

	fmt.Println(eng.BestScore)
	fmt.Println()
	fmt.Println(eng.Best.Symbol())
	fmt.Println()
	fmt.Println(eng.Best.DumpScores())
}
