package day10_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day10"
)

func Test(t *testing.T) {

	ans1, ans2 := day10.Day10(strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`))
	if ans1 != 220 {
		t.Fatalf("expected 220 got %d", ans1)
	}
	if ans2 != 19208 {
		t.Fatalf("expected 19208 got %d", ans2)
	}
}
