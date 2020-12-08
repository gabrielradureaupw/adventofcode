package day8_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day8"
)

func Test(t *testing.T) {

	ans1, ans2 := day8.Day8(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`))
	if ans1 != 5 {
		t.Fatalf("expected 5 got %d", ans1)
	}
	if ans2 != 8 {
		t.Fatalf("expected 8 got %d", ans1)
	}
}
