package day7_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day7"
)

func Test(t *testing.T) {

	ans1, _ := day7.Day7(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
	dark orange bags contain 3 bright white bags, 4 muted yellow bags.
	bright white bags contain 1 shiny gold bag.
	muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
	shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
	dark olive bags contain 3 faded blue bags, 4 dotted black bags.
	vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
	faded blue bags contain no other bags.
	dotted black bags contain no other bags.
	`))
	if ans1 != 4 {
		t.Fatalf("expected 4 got %d", ans1)
	}
	_, ans2 := day7.Day7(strings.NewReader(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`))
	if ans2 != 126 {
		t.Fatalf("expected 126 got %d", ans2)
	}
}
