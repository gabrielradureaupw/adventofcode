package day3

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
	it "github.com/yanatan16/itertools"
)

const day3 = 3

// Response _
func Response() string {
	var grid []string
	if b, err := ioutil.ReadAll(inputs.FetchInput(day3)); err != nil {
		panic(err)
	} else {
		grid = strings.Split(string(b), "\n")
		grid = grid[:len(grid)-1] // ignore empty line
	}
	patternSize := len(grid[0])

	type Slope struct{ right, down int }

	slideWithSlope := func(slope Slope) (treesEncountered int) {
		for x, y := 0, 0; y < len(grid); x, y = x+slope.right, y+slope.down {
			if grid[y][x%patternSize] == '#' {
				treesEncountered++
			}
		}
		return
	}

	ansPart1 := slideWithSlope(Slope{3, 1})
	ansPart2 := it.Reduce(
		it.Map(
			func(i interface{}) interface{} {
				return slideWithSlope(i.(Slope))
			},
			it.New(Slope{1, 1}, Slope{3, 1}, Slope{5, 1}, Slope{7, 1}, Slope{1, 2}),
		),
		func(memo, element interface{}) interface{} { return memo.(int) * element.(int) },
		1,
	).(int)

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}
