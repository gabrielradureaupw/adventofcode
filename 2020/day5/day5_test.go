package day5_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day5"
)

func Test(t *testing.T) {
	input := strings.NewReader(`FBFBBFFRLR`)

	ans1, _ := day5.Day5(input)
	if ans1 != 357 {
		t.Fatalf("expected 357 got %d", ans1)
	}
}
