package day6

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day6 = 6

// Response _
func Response() string {

	ansPart1, ansPart2 := Day6(inputs.FetchInput(day6))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

const form = "abcdefghijklmnopqrstuvwxyz"

// Day6 _
func Day6(r io.Reader) (ansPart1, ansPart2 int) {
	sc := bufio.NewScanner(r)

	set := make(map[rune]struct{})
	groupAnswers := form
	for sc.Scan() {
		if line := sc.Text(); line == "" || line == "\n" {
			ansPart1 += len(set)
			ansPart2 += len(groupAnswers)
			set = make(map[rune]struct{})
			groupAnswers = form
		} else {
			commonAnswers := strings.Builder{}
			for _, b := range line {
				set[b] = struct{}{}
				if strings.IndexRune(groupAnswers, b) != -1 {
					commonAnswers.WriteRune(b)
				}
			}
			groupAnswers = commonAnswers.String()
		}
	}
	ansPart1 += len(set)
	ansPart2 += len(groupAnswers)
	return
}
