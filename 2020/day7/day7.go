package day7

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day7 = 7

// Response _
func Response() string {

	ansPart1, ansPart2 := Day7(inputs.FetchInput(day7))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

type NBag struct {
	N     int
	Color string
}
type Rule struct {
	Src NBag
	Dst []NBag
}

var ruleRgx = regexp.MustCompile(`(?P<count>\d+ )?(?P<color>\w+ \w+) bags?`)
var countIdx = ruleRgx.SubexpIndex("count")
var coloridx = ruleRgx.SubexpIndex("color")

type bagMatch []string

func (bm bagMatch) NBag() NBag {
	return NBag{
		N:     bm.count(),
		Color: bm.color(),
	}
}

func (bm bagMatch) count() (n int) {
	fmt.Sscan(bm[countIdx], &n)
	return
}
func (bm bagMatch) color() string {
	return bm[coloridx]
}

// Day7 _
func Day7(r io.Reader) (ansPart1, ansPart2 int) {
	sc := bufio.NewScanner(r)

	rules := make(map[string]Rule)
	for sc.Scan() {
		line := sc.Text()
		matches := ruleRgx.FindAllStringSubmatch(line, -1)

		rule := Rule{}
		for i, match := range matches {
			bm := bagMatch(match)
			if i == 0 {
				rule.Src = bm.NBag()
				continue
			}
			rule.Dst = append(rule.Dst, bm.NBag())
		}
		rules[rule.Src.Color] = rule
	}

	for src := range rules {
		if src == "shiny gold" {
			continue
		}
		if canContainShinyGoldBag(src, rules) {
			ansPart1++
		}
	}
	ansPart2 = countBagCapacity("shiny gold", rules)
	return
}

func canContainShinyGoldBag(color string, rules map[string]Rule) bool {
	for _, nbag := range rules[color].Dst {
		if nbag.Color == "shiny gold" || canContainShinyGoldBag(nbag.Color, rules) {
			return true
		}
	}
	return false
}

func countBagCapacity(color string, rules map[string]Rule) int {
	sum := 0
	for _, nbag := range rules[color].Dst {
		sum += nbag.N * (countBagCapacity(nbag.Color, rules) + 1)
	}
	return sum
}
