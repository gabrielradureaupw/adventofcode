package day12_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day12"
)

func Test(t *testing.T) {

	ans1, ans2 := day12.Day12(strings.NewReader(`F10
N3
F7
R90
F11
`))
	if ans1 != 25 {
		t.Fatalf("expected 25 got %d", ans1)
	}
	if ans2 != 286 {
		t.Fatalf("expected 286 got %d", ans2)
	}
}
