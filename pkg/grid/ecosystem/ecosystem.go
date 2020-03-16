package ecosystem

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/card"
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
	"log"
	"strings"
)

const (
	rows    = 4
	columns = 5
)

type ecosystem struct {
	Cards [][]grid.Card

	BearScore      int
	BeeScore       int
	DeerScore      int
	DragonflyScore int
	EagleScore     int
	FoxScore       int
	MeadowScore    int
	RabbitScore    int
	StreamScore    int
	TroutScore     int
	WolfScore      int

	Gaps     int
	GapScore int

	Total int
}

func New() grid.Grid {
	e := ecosystem{}

	var cs [][]grid.Card
	for i := 0; i < rows; i++ {
		var row []grid.Card
		for j := 0; j < columns; j++ {
			vc := card.Vacant{}
			row = append(row, vc)
		}
		cs = append(cs, row)
	}

	e.Cards = cs

	return &e
}

func From(m string) grid.Grid {
	var cs [][]grid.Card

	rawRows := strings.Split(strings.Trim(m, "\n\t "), "\n")

	var rs []string
	for _, rawRow := range rawRows {
		if strings.Trim(rawRow, "\t ") != "" {
			rs = append(rs, rawRow)
		}
	}

	if len(rs) != rows {
		panic("invalid symbol map, incorrect number of rows")
	}

	for ri, r := range rs {
		rawSymbols := strings.Split(strings.Trim(r, "\t "), "")

		var symbols []string
		for _, rawSymbol := range rawSymbols {
			if strings.Trim(rawSymbol, "\t ") != "" {
				symbols = append(symbols, strings.Trim(rawSymbol, "\t "))
			}
		}

		if len(symbols) != columns {
			fmt.Println(symbols)
			panic("invalid symbol map, incorrect number of columns (" + fmt.Sprintf("%v", len(symbols)) + ") in row " + fmt.Sprintf("%v", ri))
		}

		var row []grid.Card
		for _, symbol := range symbols {
			row = append(row, card.From(symbol))
		}

		cs = append(cs, row)
	}

	e := ecosystem{
		Cards: cs,
	}

	return &e
}

func (e *ecosystem) Place(c grid.Card, l location.Location) {
	if _, ok := e.At(l).(card.Vacant); !ok {
		panic(grid.ErrCellAlreadyPopulated)
	}

	e.Cards[l.X][l.Y] = c
}

func (e *ecosystem) At(l location.Location) grid.Card {
	if l.X < 0 || l.X >= rows || l.Y < 0 || l.Y >= columns {
		return card.Vacant{}
	}
	return e.Cards[l.X][l.Y]
}

func (e *ecosystem) Adjacent(l location.Location) [4]grid.Card {
	return [4]grid.Card{
		e.At(l.Up()),
		e.At(l.Right()),
		e.At(l.Down()),
		e.At(l.Left()),
	}
}

func (e *ecosystem) DoubleAdjacent(l location.Location) [12]grid.Card {
	return [12]grid.Card{
		e.At(l.Up()),
		e.At(l.Right()),
		e.At(l.Down()),
		e.At(l.Left()),
		e.At(l.Up().Up()),
		e.At(l.Up().Right()),
		e.At(l.Up().Left()),
		e.At(l.Down().Down()),
		e.At(l.Down().Right()),
		e.At(l.Down().Left()),
		e.At(l.Right().Right()),
		e.At(l.Left().Left()),
	}
}

func (e *ecosystem) Calculate() (int, error) {
	// TODO: If needed, re enable this
	//for _, row := range e.Cards {
	//	for _, c := range row {
	//		switch c.(type) {
	//		case card.Vacant:
	//			return 0, grid.ErrPartiallyFilledGrid
	//		}
	//	}
	//}

	e.BearScore = e.calculateBear()
	e.BeeScore = e.calculateBee()
	e.DeerScore = e.calculateDeer()
	e.DragonflyScore = e.calculateDragonfly()
	e.EagleScore = e.calculateEagle()
	e.FoxScore = e.calculateFox()
	e.MeadowScore = e.calculateMeadow()
	e.RabbitScore = e.calculateRabbit()
	e.StreamScore = e.calculateStream()
	e.TroutScore = e.calculateTrout()
	e.WolfScore = e.calculateWolf()

	e.Gaps = e.calculateGaps()
	e.GapScore = e.calculateGapScore()

	e.Total = e.calculateTotal()

	return e.Total, nil
}

func (e *ecosystem) calculateBear() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Bear); ok {
				acs := e.Adjacent(location.New(x, y))

				for _, ac := range acs {
					switch ac.(type) {
					case card.Trout:
						v += 2
					case card.Bee:
						v += 2
					}
				}
			}
		}
	}

	return v
}

func (e *ecosystem) calculateBee() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Bee); ok {
				acs := e.Adjacent(location.New(x, y))

				for _, ac := range acs {
					switch ac.(type) {
					case card.Meadow:
						v += 3
					}
				}
			}
		}
	}

	return v
}

func (e *ecosystem) calculateDeer() int {
	v := 0

	rowFound := map[int]bool{}
	colFound := map[int]bool{}

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Deer); ok {
				rowFound[x] = true
				colFound[y] = true
			}
		}
	}

	for range rowFound {
		v += 2
	}

	for range colFound {
		v += 2
	}

	return v
}

func (e *ecosystem) calculateDragonfly() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Dragonfly); ok {
				l := location.New(x, y)

				cellChecked := map[string]bool{}

				ul := l.Up()
				uc := e.At(ul)
				if !cellChecked[ul.Coords()] {
					cellChecked[ul.Coords()] = true

					if _, ok := uc.(card.Stream); ok {
						v += e.traverseStream(ul, &cellChecked)
					}
				}

				rl := l.Right()
				rc := e.At(rl)
				if !cellChecked[rl.Coords()] {
					cellChecked[rl.Coords()] = true

					if _, ok := rc.(card.Stream); ok {
						v += e.traverseStream(rl, &cellChecked)
					}
				}

				dl := l.Down()
				bc := e.At(dl)
				if !cellChecked[dl.Coords()] {
					cellChecked[dl.Coords()] = true

					if _, ok := bc.(card.Stream); ok {
						v += e.traverseStream(dl, &cellChecked)
					}
				}

				ll := l.Left()
				lc := e.At(ll)
				if !cellChecked[ll.Coords()] {
					cellChecked[ll.Coords()] = true

					if _, ok := lc.(card.Stream); ok {
						v += e.traverseStream(ll, &cellChecked)
					}
				}
			}
		}
	}

	return v
}

func (e *ecosystem) traverseStream(l location.Location, cellChecked *map[string]bool) int {
	if _, ok := e.At(l).(card.Stream); !ok {
		panic("cannot traverseStream from a non-stream card!")
	}

	d := 1

	ul := l.Up()
	if !(*cellChecked)[ul.Coords()] {
		(*cellChecked)[ul.Coords()] = true

		tc := e.At(ul)
		if _, ok := tc.(card.Stream); ok {
			d += e.traverseStream(ul, cellChecked)
		}
	}

	rl := l.Right()
	if !(*cellChecked)[rl.Coords()] {
		(*cellChecked)[rl.Coords()] = true

		rc := e.At(rl)
		if _, ok := rc.(card.Stream); ok {
			d += e.traverseStream(rl, cellChecked)
		}
	}

	dl := l.Down()
	if !(*cellChecked)[dl.Coords()] {
		(*cellChecked)[dl.Coords()] = true

		bc := e.At(dl)
		if _, ok := bc.(card.Stream); ok {
			d += e.traverseStream(dl, cellChecked)
		}
	}

	ll := l.Left()
	if !(*cellChecked)[ll.Coords()] {
		(*cellChecked)[ll.Coords()] = true

		lc := e.At(ll)
		if _, ok := lc.(card.Stream); ok {
			d += e.traverseStream(ll, cellChecked)
		}
	}

	return d
}

func (e *ecosystem) calculateEagle() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Eagle); ok {
				dacs := e.DoubleAdjacent(location.New(x, y))

				for _, dac := range dacs {
					switch dac.(type) {
					case card.Trout:
						v += 2
					case card.Rabbit:
						v += 2
					}
				}
			}
		}
	}

	return v
}

func (e *ecosystem) calculateFox() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Fox); ok {
				acs := e.Adjacent(location.New(x, y))

				found := false
				for _, ac := range acs {
					switch ac.(type) {
					case card.Bear:
						found = true
					case card.Wolf:
						found = true
					}
					if found {
						break
					}
				}

				if !found {
					v += 3
				}
			}
		}
	}

	return v
}

func (e *ecosystem) calculateMeadow() int {
	v := 0

	cellChecked := map[string]bool{}

	for x, row := range e.Cards {
		for y, c := range row {
			l := location.New(x, y)
			if cellChecked[l.Coords()] {
				continue
			}
			cellChecked[l.Coords()] = true

			if _, ok := c.(card.Meadow); ok {
				ml := e.traverseMeadow(l, &cellChecked)
				switch ml {
				case 1:
					v += 0
				case 2:
					v += 3
				case 3:
					v += 6
				case 4:
					v += 10
				default:
					v += 15
				}
			}
		}
	}

	return v
}

func (e *ecosystem) traverseMeadow(l location.Location, cellChecked *map[string]bool) int {
	if _, ok := e.At(l).(card.Meadow); !ok {
		panic("cannot traverseMeadow from a non-meadow card!")
	}

	d := 1

	ul := l.Up()
	if !(*cellChecked)[ul.Coords()] {
		(*cellChecked)[ul.Coords()] = true

		tc := e.At(ul)
		if _, ok := tc.(card.Meadow); ok {
			d += e.traverseMeadow(ul, cellChecked)
		}
	}

	rl := l.Right()
	if !(*cellChecked)[rl.Coords()] {
		(*cellChecked)[rl.Coords()] = true

		rc := e.At(rl)
		if _, ok := rc.(card.Meadow); ok {
			d += e.traverseMeadow(rl, cellChecked)
		}
	}

	dl := l.Down()
	if !(*cellChecked)[dl.Coords()] {
		(*cellChecked)[dl.Coords()] = true

		bc := e.At(dl)
		if _, ok := bc.(card.Meadow); ok {
			d += e.traverseMeadow(dl, cellChecked)
		}
	}

	ll := l.Left()
	if !(*cellChecked)[ll.Coords()] {
		(*cellChecked)[ll.Coords()] = true

		lc := e.At(ll)
		if _, ok := lc.(card.Meadow); ok {
			d += e.traverseMeadow(ll, cellChecked)
		}
	}

	return d
}

func (e *ecosystem) calculateRabbit() int {
	v := 0

	for _, row := range e.Cards {
		for _, c := range row {
			if _, ok := c.(card.Rabbit); ok {
				v++
			}
		}
	}

	return v
}

func (e *ecosystem) calculateStream() int {
	found := false

	for _, row := range e.Cards {
		for _, c := range row {
			if _, ok := c.(card.Stream); ok {
				if !found {
					found = true
				} else {
					log.Println("WARNING: Inefficient grid. More than 1 stream found, but only 1 sufficient for maximum points (unless this is a dragonfly play).")
				}
			}
		}
	}

	if !found {
		log.Println("WARNING: Inefficient grid. Should include 1 stream to gain 8 points.")
		return 0
	}

	return 8
}

func (e *ecosystem) calculateTrout() int {
	v := 0

	for x, row := range e.Cards {
		for y, c := range row {
			if _, ok := c.(card.Trout); ok {
				acs := e.Adjacent(location.New(x, y))

				for _, ac := range acs {
					switch ac.(type) {
					case card.Dragonfly:
						v += 2
					case card.Stream:
						v += 2
					}
				}
			}
		}
	}

	return v
}

func (e *ecosystem) calculateWolf() int {
	found := false

	for _, row := range e.Cards {
		for _, c := range row {
			if _, ok := c.(card.Wolf); ok {
				if !found {
					found = true
				} else {
					log.Println("WARNING: Inefficient grid. More than 1 wolf found, but only 1 sufficient for maximum points.")
				}
			}
		}
	}

	if !found {
		log.Println("WARNING: Inefficient grid. Should include 1 wolf to gain 12 points.")
		return 0
	}

	return 12
}

func (e *ecosystem) Symbol() string {
	out := ""

	for _, row := range e.Cards {
		for _, c := range row {
			out += c.Symbol() + " "
		}
		out += "\n"
	}

	return out
}

func (e *ecosystem) calculateGaps() int {
	gaps := 0

	if e.BearScore == 0 {
		gaps++
	}
	if e.BeeScore == 0 {
		gaps++
	}
	if e.DeerScore == 0 {
		gaps++
	}
	if e.DragonflyScore == 0 {
		gaps++
	}
	if e.EagleScore == 0 {
		gaps++
	}
	if e.FoxScore == 0 {
		gaps++
	}
	if e.MeadowScore == 0 {
		gaps++
	}
	if e.RabbitScore == 0 {
		gaps++
	}
	if e.StreamScore == 0 {
		gaps++
	}
	if e.TroutScore == 0 {
		gaps++
	}
	if e.WolfScore == 0 {
		gaps++
	}

	return gaps
}

func (e *ecosystem) calculateGapScore() int {
	switch e.Gaps {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		return 12
	case 3:
		return 7
	case 4:
		return 3
	case 5:
		return 0
	default:
		return -5
	}
}

func (e *ecosystem) calculateTotal() int {
	total := 0

	total += e.BearScore
	total += e.BeeScore
	total += e.DeerScore
	total += e.DragonflyScore
	total += e.EagleScore
	total += e.FoxScore
	total += e.MeadowScore
	total += e.RabbitScore
	total += e.StreamScore
	total += e.TroutScore
	total += e.WolfScore

	total += e.GapScore

	return total
}

func (e *ecosystem) DumpScores() string {
	out := ""

	out += fmt.Sprintf("BearScore: %v\n", e.BearScore)
	out += fmt.Sprintf("BeeScore: %v\n", e.BeeScore)
	out += fmt.Sprintf("DeerScore: %v\n", e.DeerScore)
	out += fmt.Sprintf("DragonflyScore: %v\n", e.DragonflyScore)
	out += fmt.Sprintf("EagleScore: %v\n", e.EagleScore)
	out += fmt.Sprintf("FoxScore: %v\n", e.FoxScore)
	out += fmt.Sprintf("MeadowScore: %v\n", e.MeadowScore)
	out += fmt.Sprintf("RabbitScore: %v\n", e.RabbitScore)
	out += fmt.Sprintf("StreamScore: %v\n", e.StreamScore)
	out += fmt.Sprintf("TroutScore: %v\n", e.TroutScore)
	out += fmt.Sprintf("WolfScore: %v\n", e.WolfScore)
	out += "\n"
	out += fmt.Sprintf("Gaps: %v\n", e.Gaps)
	out += fmt.Sprintf("GapScore: %v\n", e.GapScore)
	out += "\n"
	out += fmt.Sprintf("Total: %v\n", e.Total)

	return out
}
