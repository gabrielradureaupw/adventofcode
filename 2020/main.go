package main

import (
	"fmt"

	"github.com/gabrielradureaupw/adventofcode/2020/day1"
	"github.com/gabrielradureaupw/adventofcode/2020/day2"
	"github.com/gabrielradureaupw/adventofcode/2020/day3"
	"github.com/gabrielradureaupw/adventofcode/2020/day4"
)

func main() {
	fmt.Printf(`
	Hello World!

	here are my responses:

	- day1: [%s],
	- day2: [%s],
	- day3: [%s],
	- day4: [%s],
	`,
		day1.Response(),
		day2.Response(),
		day3.Response(),
		day4.Response(),
	)
}
