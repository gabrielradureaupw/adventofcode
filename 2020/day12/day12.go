package day12

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day12 = 12

// Response _
func Response() string {

	ansPart1, ansPart2 := Day12(inputs.FetchInput(day12))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

// Day12 _
func Day12(r io.Reader) (ansPart1, ansPart2 int) {
	scanner := bufio.NewScanner(r)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	ansPart1 = withoutWayPoint(instructions)
	ansPart2 = withWayPoint(instructions)
	return
}

func withoutWayPoint(instructions []string) (manhattan int) {
	posNorth, posEast := 0, 0
	forwardNorth, forwardEast := 0, 1
	deg := 0

	for _, instruction := range instructions {
		var action rune
		var value int
		fmt.Sscanf(instruction, "%c%d", &action, &value)
		switch action {
		case 'N':
			posNorth += value
		case 'S':
			posNorth -= value
		case 'E':
			posEast += value
		case 'W':
			posEast -= value
		case 'L':
			deg += 2 * value
			fallthrough
		case 'R':
			deg -= value
			switch ((deg%360 + 360) % 360) / 90 {
			case 0:
				forwardNorth, forwardEast = 0, 1
			case 1:
				forwardNorth, forwardEast = 1, 0
			case 2:
				forwardNorth, forwardEast = 0, -1
			case 3:
				forwardNorth, forwardEast = -1, 0
			default:
				panic(fmt.Errorf("pivoted %c %d => %d", action, value, ((deg%360+360)%360)/90))
			}
		case 'F':
			posNorth += forwardNorth * value
			posEast += forwardEast * value
		}
	}
	posSouth, posWest := 0, 0
	if posNorth < 0 {
		posSouth -= posNorth
		posNorth = 0
	}
	if posEast < 0 {
		posWest -= posEast
		posEast = 0
	}
	return posEast + posNorth + posSouth + posWest
}

func withWayPoint(instructions []string) (manhattan int) {
	posNorth, posEast := 0, 0
	wpForwardNorth, wpForwardEast := 1, 10

	for _, instruction := range instructions {
		var action rune
		var value int
		fmt.Sscanf(instruction, "%c%d", &action, &value)
		switch action {
		case 'N':
			wpForwardNorth += value
		case 'S':
			wpForwardNorth -= value
		case 'E':
			wpForwardEast += value
		case 'W':
			wpForwardEast -= value
		case 'R':
			value *= -1
			fallthrough
		case 'L':
			switch ((value%360 + 360) % 360) / 90 {
			case 0:
				// wpForwardNorth, wpForwardEast = wpForwardNorth, wpForwardEast
			case 1:
				wpForwardNorth, wpForwardEast = wpForwardEast, -1*wpForwardNorth
			case 2:
				wpForwardNorth, wpForwardEast = -1*wpForwardNorth, -1*wpForwardEast
			case 3:
				wpForwardNorth, wpForwardEast = -1*wpForwardEast, wpForwardNorth
			default:
				panic(fmt.Errorf("pivoted %c %d => %d", action, value, ((value%360+360)%360)/90))
			}
		case 'F':
			posNorth += wpForwardNorth * value
			posEast += wpForwardEast * value
		}
	}
	posSouth, posWest := 0, 0
	if posNorth < 0 {
		posSouth -= posNorth
		posNorth = 0
	}
	if posEast < 0 {
		posWest -= posEast
		posEast = 0
	}
	return posEast + posNorth + posSouth + posWest
}
