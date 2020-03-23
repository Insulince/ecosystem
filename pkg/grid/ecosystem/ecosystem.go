package ecosystem

import (
	"fmt"
	"github.com/Insulince/ecosystem/pkg/card"
	"github.com/Insulince/ecosystem/pkg/grid"
	"github.com/Insulince/ecosystem/pkg/location"
	"strings"
)

const (
	rows = 4
	cols = 5

	Candidates = `
	bbbbbbbbbbbb
	BBBBBB
	TTTTTTTTTT
	EEEEEEEE
	RRRRRRRR
	FFFFFFFFFFFF
	SSSSSSSSSSSSSSSSSSS
	DDDDDDDDDDDD
	`
	//	BB
	// 	MMMMMMMMMMMMMMMMMMMM
	// 	dddddddd
	// 	S
	// 	WWWWWWWWWWWW
)

type Config struct {
}

type Ecosystem struct {
	Cards [rows][cols]grid.Card `json:"-"`

	Gaps int `json:"gaps"`

	Scores Scores `json:"scores"`
}

type Scores struct {
	Bear      int `json:"bear"`
	Bee       int `json:"bee"`
	Deer      int `json:"deer"`
	Dragonfly int `json:"dragonfly"`
	Eagle     int `json:"eagle"`
	Fox       int `json:"fox"`
	Meadow    int `json:"meadow"`
	Rabbit    int `json:"rabbit"`
	Stream    int `json:"stream"`
	Trout     int `json:"trout"`
	Wolf      int `json:"wolf"`

	Gaps int `json:"gaps"`
}

func (s Scores) Total() int {
	total := 0

	total += s.Bear
	total += s.Bee
	total += s.Deer
	total += s.Dragonfly
	total += s.Eagle
	total += s.Fox
	total += s.Meadow
	total += s.Rabbit
	total += s.Stream
	total += s.Trout
	total += s.Wolf

	total += s.Gaps

	return total
}

func New(c Config) *Ecosystem {
	eco := &Ecosystem{}

	return eco
}

func FromMap(m string) *Ecosystem {
	eco := &Ecosystem{}

	m = cleanMap(m)

	eco.Cards = mapToCards(m)

	return eco
}

func (eco *Ecosystem) Place(c grid.Card, l location.Location) {
	if eco.At(l) != nil {
		panic("already populated with a card") // TODO
	}

	c = c.Place(l)

	eco.Cards[l.X][l.Y] = c
}

func (eco *Ecosystem) At(l location.Location) grid.Card {
	if l.X < 0 || l.X >= rows || l.Y < 0 || l.Y >= cols {
		return nil
	}
	return eco.Cards[l.X][l.Y]
}

func (eco *Ecosystem) Adjacent(l location.Location) [4]grid.Card {
	return [4]grid.Card{
		eco.At(l.Up()),
		eco.At(l.Right()),
		eco.At(l.Down()),
		eco.At(l.Left()),
	}
}

func (eco *Ecosystem) DoubleAdjacent(l location.Location) [12]grid.Card {
	return [12]grid.Card{
		eco.At(l.Up()),
		eco.At(l.Right()),
		eco.At(l.Down()),
		eco.At(l.Left()),
		eco.At(l.Up().Up()),
		eco.At(l.Up().Right()),
		eco.At(l.Up().Left()),
		eco.At(l.Down().Down()),
		eco.At(l.Down().Right()),
		eco.At(l.Down().Left()),
		eco.At(l.Right().Right()),
		eco.At(l.Left().Left()),
	}
}

func (eco *Ecosystem) Score() (int, error) {
	eco.Scores.Bear = eco.calculateBear()
	eco.Scores.Bee = eco.calculateBee()
	eco.Scores.Deer = eco.calculateDeer()
	eco.Scores.Dragonfly = eco.calculateDragonfly()
	eco.Scores.Eagle = eco.calculateEagle()
	eco.Scores.Fox = eco.calculateFox()
	eco.Scores.Meadow = eco.calculateMeadow()
	eco.Scores.Rabbit = eco.calculateRabbit()
	eco.Scores.Stream = eco.calculateStream()
	eco.Scores.Trout = eco.calculateTrout()
	eco.Scores.Wolf = eco.calculateWolf()

	eco.Gaps = eco.calculateGaps()
	eco.Scores.Gaps = eco.calculateGapScore()

	total := eco.Scores.Total()

	return total, nil
}

func (eco *Ecosystem) calculateBear() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Bear); ok {
				acs := eco.Adjacent(location.New(x, y))

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

func (eco *Ecosystem) calculateBee() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Bee); ok {
				acs := eco.Adjacent(location.New(x, y))

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

func (eco *Ecosystem) calculateDeer() int {
	v := 0

	rowFound := map[int]bool{}
	colFound := map[int]bool{}

	for x, row := range eco.Cards {
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

func (eco *Ecosystem) calculateDragonfly() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Dragonfly); ok {
				l := location.New(x, y)

				cellChecked := map[string]bool{}

				ul := l.Up()
				uc := eco.At(ul)
				if !cellChecked[ul.Coords()] {
					cellChecked[ul.Coords()] = true

					if _, ok := uc.(card.Stream); ok {
						v += eco.traverseStream(ul, &cellChecked)
					}
				}

				rl := l.Right()
				rc := eco.At(rl)
				if !cellChecked[rl.Coords()] {
					cellChecked[rl.Coords()] = true

					if _, ok := rc.(card.Stream); ok {
						v += eco.traverseStream(rl, &cellChecked)
					}
				}

				dl := l.Down()
				bc := eco.At(dl)
				if !cellChecked[dl.Coords()] {
					cellChecked[dl.Coords()] = true

					if _, ok := bc.(card.Stream); ok {
						v += eco.traverseStream(dl, &cellChecked)
					}
				}

				ll := l.Left()
				lc := eco.At(ll)
				if !cellChecked[ll.Coords()] {
					cellChecked[ll.Coords()] = true

					if _, ok := lc.(card.Stream); ok {
						v += eco.traverseStream(ll, &cellChecked)
					}
				}
			}
		}
	}

	return v
}

func (eco *Ecosystem) traverseStream(l location.Location, cellChecked *map[string]bool) int {
	if _, ok := eco.At(l).(card.Stream); !ok {
		panic("cannot traverseStream from a non-stream card!")
	}

	d := 1

	ul := l.Up()
	if !(*cellChecked)[ul.Coords()] {
		(*cellChecked)[ul.Coords()] = true

		tc := eco.At(ul)
		if _, ok := tc.(card.Stream); ok {
			d += eco.traverseStream(ul, cellChecked)
		}
	}

	rl := l.Right()
	if !(*cellChecked)[rl.Coords()] {
		(*cellChecked)[rl.Coords()] = true

		rc := eco.At(rl)
		if _, ok := rc.(card.Stream); ok {
			d += eco.traverseStream(rl, cellChecked)
		}
	}

	dl := l.Down()
	if !(*cellChecked)[dl.Coords()] {
		(*cellChecked)[dl.Coords()] = true

		bc := eco.At(dl)
		if _, ok := bc.(card.Stream); ok {
			d += eco.traverseStream(dl, cellChecked)
		}
	}

	ll := l.Left()
	if !(*cellChecked)[ll.Coords()] {
		(*cellChecked)[ll.Coords()] = true

		lc := eco.At(ll)
		if _, ok := lc.(card.Stream); ok {
			d += eco.traverseStream(ll, cellChecked)
		}
	}

	return d
}

func (eco *Ecosystem) calculateEagle() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Eagle); ok {
				dacs := eco.DoubleAdjacent(location.New(x, y))

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

func (eco *Ecosystem) calculateFox() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Fox); ok {
				acs := eco.Adjacent(location.New(x, y))

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

func (eco *Ecosystem) calculateMeadow() int {
	v := 0

	cellChecked := map[string]bool{}

	for x, row := range eco.Cards {
		for y, c := range row {
			l := location.New(x, y)
			if cellChecked[l.Coords()] {
				continue
			}
			cellChecked[l.Coords()] = true

			if _, ok := c.(card.Meadow); ok {
				ml := eco.traverseMeadow(l, &cellChecked)
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

func (eco *Ecosystem) traverseMeadow(l location.Location, cellChecked *map[string]bool) int {
	if _, ok := eco.At(l).(card.Meadow); !ok {
		panic("cannot traverseMeadow from a non-meadow card!")
	}

	d := 1

	ul := l.Up()
	if !(*cellChecked)[ul.Coords()] {
		(*cellChecked)[ul.Coords()] = true

		tc := eco.At(ul)
		if _, ok := tc.(card.Meadow); ok {
			d += eco.traverseMeadow(ul, cellChecked)
		}
	}

	rl := l.Right()
	if !(*cellChecked)[rl.Coords()] {
		(*cellChecked)[rl.Coords()] = true

		rc := eco.At(rl)
		if _, ok := rc.(card.Meadow); ok {
			d += eco.traverseMeadow(rl, cellChecked)
		}
	}

	dl := l.Down()
	if !(*cellChecked)[dl.Coords()] {
		(*cellChecked)[dl.Coords()] = true

		bc := eco.At(dl)
		if _, ok := bc.(card.Meadow); ok {
			d += eco.traverseMeadow(dl, cellChecked)
		}
	}

	ll := l.Left()
	if !(*cellChecked)[ll.Coords()] {
		(*cellChecked)[ll.Coords()] = true

		lc := eco.At(ll)
		if _, ok := lc.(card.Meadow); ok {
			d += eco.traverseMeadow(ll, cellChecked)
		}
	}

	return d
}

func (eco *Ecosystem) calculateRabbit() int {
	v := 0

	for _, row := range eco.Cards {
		for _, c := range row {
			if _, ok := c.(card.Rabbit); ok {
				v++
			}
		}
	}

	return v
}

func (eco *Ecosystem) calculateStream() int {
	found := false

	for _, row := range eco.Cards {
		for _, c := range row {
			if _, ok := c.(card.Stream); ok {
				if !found {
					found = true
				} else {
					//log.Println("WARNING: Inefficient grid. More than 1 stream found, but only 1 sufficient for maximum points (unless this is a dragonfly play).")
				}
			}
		}
	}

	if !found {
		//log.Println("WARNING: Inefficient grid. Should include 1 stream to gain 8 points.")
		return 0
	}

	return 8
}

func (eco *Ecosystem) calculateTrout() int {
	v := 0

	for x, row := range eco.Cards {
		for y, c := range row {
			if _, ok := c.(card.Trout); ok {
				acs := eco.Adjacent(location.New(x, y))

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

func (eco *Ecosystem) calculateWolf() int {
	found := false

	for _, row := range eco.Cards {
		for _, c := range row {
			if _, ok := c.(card.Wolf); ok {
				if !found {
					found = true
				} else {
					//log.Println("WARNING: Inefficient grid. More than 1 wolf found, but only 1 sufficient for maximum points.")
				}
			}
		}
	}

	if !found {
		//log.Println("WARNING: Inefficient grid. Should include 1 wolf to gain 12 points.")
		return 0
	}

	return 12
}

func (eco *Ecosystem) Map() string {
	out := ""

	for _, row := range eco.Cards {
		for _, c := range row {
			out += c.Symbol() + " "
		}
		out += "\n"
	}

	out = out[:len(out)-1]

	return out
}

func (eco *Ecosystem) DumpScores() string {
	// TODO
	fmt.Println("aint done yet nerd")
	return ""
}

func (eco *Ecosystem) calculateGaps() int {
	gaps := 0

	if eco.Scores.Bear == 0 {
		gaps++
	}
	if eco.Scores.Bee == 0 {
		gaps++
	}
	if eco.Scores.Deer == 0 {
		gaps++
	}
	if eco.Scores.Dragonfly == 0 {
		gaps++
	}
	if eco.Scores.Eagle == 0 {
		gaps++
	}
	if eco.Scores.Fox == 0 {
		gaps++
	}
	if eco.Scores.Meadow == 0 {
		gaps++
	}
	if eco.Scores.Rabbit == 0 {
		gaps++
	}
	if eco.Scores.Stream == 0 {
		gaps++
	}
	if eco.Scores.Trout == 0 {
		gaps++
	}
	if eco.Scores.Wolf == 0 {
		gaps++
	}

	return gaps
}

func (eco *Ecosystem) calculateGapScore() int {
	switch eco.Gaps {
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

func cleanMap(m string) string {
	m = strings.Trim(m, "\n\t ")
	lines := strings.Split(m, "\n")

	for li := 0; li < len(lines); li++ {
		lines[li] = strings.Trim(lines[li], "\t ")

		if lines[li] == "" {
			lines = append(lines[:li], lines[li+1:]...)
			li--
			continue
		}

		items := strings.Split(lines[li], "")
		for ii := 0; ii < len(items); ii++ {
			if items[ii] == " " {
				items = append(items[:ii], items[ii+1:]...)
				ii--
				continue
			}
		}

		if len(items) != cols {
			panic("invalid symbol map, incorrect number of columns in row") // TODO
		}

		lines[li] = strings.Join(items, " ")
	}

	if len(lines) != rows {
		panic("invalid symbol map, incorrect number of rows") // TODO
	}

	m = strings.Join(lines, "\n")

	return m
}

func mapToCards(m string) [rows][cols]grid.Card {
	cards := [rows][cols]grid.Card{}

	lines := strings.Split(m, "\n")

	for ri, line := range lines {
		symbols := strings.Split(line, " ")

		for ci, symbol := range symbols {
			cards[ri][ci] = card.From(symbol)
		}
	}

	return cards
}
