package day4_test

import (
	"strings"
	"testing"

	"github.com/gabrielradureaupw/adventofcode/2020/day4"
)

func Test(t *testing.T) {
	input := strings.NewReader(`
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm
`)

	ans1, _ := day4.Day4(input)
	if ans1 != 2 {
		t.Fatalf("expected 2 got %d", ans1)
	}
}
