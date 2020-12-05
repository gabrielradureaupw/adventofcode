package day5

import (
	"bufio"
	"fmt"
	"io"
	"sort"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day5 = 5

// Response _
func Response() string {

	ansPart1, ansPart2 := Day5(inputs.FetchInput(day5))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

type seat string

func (s seat) ID() int {
	return s.Row()*8 + s.Column()
}

func (s seat) Row() int {
	return dichotomy(0, 127, string(s[:7]), 'B', 'F')
}
func (s seat) Column() int {
	return dichotomy(0, 7, string(s[7:]), 'R', 'L')
}

func dichotomy(low, high int, code string, up, down rune) int {
	diff := 1
	for i := 0; i < len(code)-1; i++ {
		diff = diff * 2
	}
	for _, c := range code {
		switch c {
		case up:
			low += diff
		case down:
			high -= diff
		}
		diff = diff / 2
	}
	return low
}

// Day5 _
func Day5(r io.Reader) (ansPart1, ansPart2 int) {
	sc := bufio.NewScanner(r)

	ids := []int{}
	for sc.Scan() {
		s := seat(sc.Text())
		id := s.ID()
		if id > ansPart1 {
			ansPart1 = id
		}
		ids = append(ids, id)
	}

	sort.IntSlice(ids).Sort()
	minSeatID := 1*8 + 0 // second row
	lastTakenSeat := minSeatID - 1
	for _, id := range ids {
		if id < minSeatID {
			continue
		}
		if id == lastTakenSeat+2 {
			ansPart2 = lastTakenSeat + 1
			break
		}
		lastTakenSeat = id
	}
	return
}
