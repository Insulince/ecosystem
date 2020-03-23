package engine

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/grid/ecosystem"
	"math/rand"
	"strings"
	"time"
)

type Engine struct {
	Candidates   string
	GridsVisited []string
	Current      string
	Best         *ecosystem.Ecosystem
	BestScore    int

	verbose bool
}

type Config struct {
	Candidates string
	Verbose    bool
}

func New(c Config) Engine {
	candidates := c.Candidates

	lines := strings.Split(strings.Trim(candidates, "\n\t "), "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\t ")
	}

	candidates = strings.Join(lines, "")

	e := Engine{
		Candidates: candidates,
		verbose:    c.Verbose,
	}

	e.NextEcosystem()

	return e
}

func (e *Engine) NextEcosystem() {
	var g string

	unique := false
	u := 0
	for !unique {
		if u > 1000000 {
			d := e.Best.Map()
			d += "\n\n" + e.Best.DumpScores()
			panic("too many non unique, here is dump:\n" + d)
		}

		g = ""
		candidates := `WMMMMMMMMMMBBS`
		candidates += buildRemainingCandidates(e.Candidates, len(candidates))

		for ri := 0; ri < 4; ri++ {
			for ci := 0; ci < 5; ci++ {
				var c string

				c, candidates = randGetCut(candidates)

				g += c
			}
			g += "\n"
		}

		if !efficient(g) {
			continue
		}

		unique = true
		for _, og := range e.GridsVisited {
			if g == og {
				fmt.Println("not unique!")
				u++
				unique = false
				break
			}
		}
	}

	e.Current = g
}

func efficient(g string) bool {
	if strings.Count(g, "W") != 1 { // Exactly 1 wolf
		return false
	}

	if strings.Count(g, "M") != 10 { // Exactly 10 meadows
		return false
	}

	if strings.Count(g, "B") < 2 { // 2 or more Bees
		return false
	}

	if strings.Index(g, "S") == -1 { // 1 or more streams
		return false
	}

	if strings.Index(g, "d") != -1 { // 0 dragonflys
		return false
	}

	return true
}

func randGetCut(s string) (string, string) {
	if len(s) == 1 {
		return s, ""
	}
	return getCut(s, rand.Intn(len(s)))
}

func getCut(s string, i int) (string, string) {
	return string(s[i]), s[:i] + s[i+1:]
}

func buildRemainingCandidates(can string, usedSpaces int) string {
	set := ""
	for i := 0; i < 20-usedSpaces; i++ {
		var c string
		c, can = randGetCut(can)
		set += c
	}
	return set
}

func (e *Engine) Run() {
	t := time.Now()
	for i := 0; i < 1000000000; i++ {
		if i%10000000 == 0 {
			fmt.Printf("--- %v - %v (%v)\n", i, e.BestScore, time.Since(t))
		} else if i%1000000 == 0 {
			fmt.Printf("------ %v - %v (%v)\n", i, e.BestScore, time.Since(t))
		} else if i%100000 == 0 {
			fmt.Printf("--------- %v - %v (%v)\n", i, e.BestScore, time.Since(t))
		} else if i%25000 == 0 {
			fmt.Printf("------------ %v - %v\n", i, e.BestScore)
		}

		e.NextEcosystem()

		eco := ecosystem.FromMap(e.Current)

		v, err := eco.Score()
		if err != nil {
			panic(err) // TODO: Handle.
		}

		if v >= e.BestScore {
			if v > e.BestScore {
				fmt.Println("=== NEW BEST ===")
			}
			e.Best = eco
			e.BestScore = v
			fmt.Printf("=== %v === (@ %v)\n", e.BestScore, i)
			fmt.Println(e.Best.Map())
			fmt.Println()
			fmt.Println(e.Best.DumpScores())
			fmt.Printf("==========\n")
		}
	}
	fmt.Println(time.Since(t))
}
