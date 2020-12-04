package day4

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day4 = 4

const passportTag = "passport"

// Passport _
type Passport struct {
	BirthYear      string `passport:"byr" valid:"required,range(1920|2002)"`
	IssueYear      string `passport:"iyr" valid:"required,range(2010|2020)"`
	ExpirationYear string `passport:"eyr" valid:"required,range(2020|2030)"`
	Height         string `passport:"hgt" valid:"required,matches(^[0-9]+(cm|in)$)"`
	HairColor      string `passport:"hcl" valid:"required,matches(^#[0-9a-f]{6}$)"`
	EyeColor       string `passport:"ecl" valid:"required,in(amb|blu|brn|gry|grn|hzl|oth)"`
	PassportID     string `passport:"pid" valid:"required,matches(^[0-9]{9}$)"`
	CountryID      string `passport:"cid" valid:"optional"`
}

func newPassport() (new Passport) {
	return Passport{CountryID: "xmas"}
}

// Set struct field with matching passport tag value
func (p *Passport) Set(key, value string) error {
	v := reflect.ValueOf(p).Elem()
	t := reflect.TypeOf(p).Elem()
	for i := 0; i < v.NumField(); i++ {
		if key == t.Field(i).Tag.Get(passportTag) {
			v.Field(i).SetString(value)
			return nil
		}
	}
	return fmt.Errorf("no struct field matching '%s' in struct %+v", key, *p)
}

// Validate _
func (p Passport) Validate() (valid1 bool, valid2 bool) {
	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" && reflect.TypeOf(p).Field(i).Name != "CountryID" {
			return false, false
		}
	}
	valid1 = true

	_, err := govalidator.ValidateStruct(p)
	valid2 = err == nil
	if valid2 {
		size, unit := 0, ""
		fmt.Fscanf(strings.NewReader(p.Height), "%d%s", &size, &unit)
		switch unit {
		case "cm":
			valid2 = size >= 150 && size <= 193
		case "in":
			valid2 = size >= 59 && size <= 76
		}
	}
	return
}

// Response _
func Response() string {

	ansPart1, ansPart2 := Day4(inputs.FetchInput(day4))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

// Day4 _
func Day4(r io.Reader) (ansPart1 int, ansPart2 int) {
	scanner := bufio.NewScanner(r)
	document := newPassport()
	checkDocument := func() {
		if valid1, valid2 := document.Validate(); valid2 {
			ansPart1++
			ansPart2++
		} else if valid1 {
			ansPart1++
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 2 {
			checkDocument()
			document = newPassport()
			continue
		}
		tokenizer := bufio.NewScanner(strings.NewReader(line))
		tokenizer.Split(bufio.ScanWords)
		for tokenizer.Scan() {
			tokenParts := strings.Split(tokenizer.Text(), ":")
			document.Set(tokenParts[0], tokenParts[1])
		}
	}
	checkDocument()
	return
}
