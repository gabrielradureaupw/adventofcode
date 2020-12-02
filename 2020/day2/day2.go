package day2

import (
	"fmt"
	"strings"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day2 = 2

// Response _
func Response() string {

	r := inputs.FetchInput(day2)

	ansPart1 := 0
	ansPart2 := 0

	var n1, n2 int
	var c byte
	var s string
	for {
		if n, _ := fmt.Fscanf(r, "%d-%d %c: %s", &n1, &n2, &c, &s); n <= 0 {
			break
		}
		count := strings.Count(s, string(c))
		if count >= n1 && count <= n2 {
			ansPart1++
		}
		if c1, c2 := s[n1-1], s[n2-1]; c1 != c2 && (c1 == c || c2 == c) {
			ansPart2++
		}
	}

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}
