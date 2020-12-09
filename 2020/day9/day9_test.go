package day9_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day9"
)

func Test(t *testing.T) {

	ans1, ans2 := day9.Day9(strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`), 5)
	if ans1 != 127 {
		t.Fatalf("expected 127 got %d", ans1)
	}
	if ans2 != 62 {
		t.Fatalf("expected 62 got %d", ans2)
	}
}
