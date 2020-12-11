package main

import (
	"fmt"

	"github.com/gabrielradureaupw/adventofcode/2020/day1"
	"github.com/gabrielradureaupw/adventofcode/2020/day10"
	"github.com/gabrielradureaupw/adventofcode/2020/day11"
	"github.com/gabrielradureaupw/adventofcode/2020/day2"
	"github.com/gabrielradureaupw/adventofcode/2020/day3"
	"github.com/gabrielradureaupw/adventofcode/2020/day4"
	"github.com/gabrielradureaupw/adventofcode/2020/day5"
	"github.com/gabrielradureaupw/adventofcode/2020/day6"
	"github.com/gabrielradureaupw/adventofcode/2020/day7"
	"github.com/gabrielradureaupw/adventofcode/2020/day8"
	"github.com/gabrielradureaupw/adventofcode/2020/day9"
)

func main() {
	fmt.Printf(`
	Hello World!

	here are my responses:

	- day1: [%s],
	- day2: [%s],
	- day3: [%s],
	- day4: [%s],
	- day5: [%s],
	- day6: [%s],
	- day7: [%s],
	- day8: [%s],
	- day9: [%s],
	- day10: [%s],
	- day11: [%s],
	`,
		day1.Response(),
		day2.Response(),
		day3.Response(),
		day4.Response(),
		day5.Response(),
		day6.Response(),
		day7.Response(),
		day8.Response(),
		day9.Response(),
		day10.Response(),
		day11.Response(),
	)
}
