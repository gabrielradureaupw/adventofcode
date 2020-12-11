package day11_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day11"
)

func Test(t *testing.T) {

	ans1, ans2 := day11.Day11(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`))
	if ans1 != 37 {
		t.Fatalf("expected 37 got %d", ans1)
	}
	if ans2 != 26 {
		t.Fatalf("expected 26 got %d", ans2)
	}
}
