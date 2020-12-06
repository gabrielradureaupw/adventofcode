package day6_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day6"
)

func Test(t *testing.T) {

	ans1, ans2 := day6.Day6(strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b
`))
	if ans1 != 11 {
		t.Fatalf("expected 11 got %d", ans1)
	}
	if ans2 != 6 {
		t.Fatalf("expected 6 got %d", ans2)
	}
}
