package day11

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day11 = 11

// Response _
func Response() string {

	ansPart1, ansPart2 := Day11(inputs.FetchInput(day11))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

// Day11 _
func Day11(r io.Reader) (ansPart1, ansPart2 int) {
	scanner := bufio.NewScanner(r)
	grid := [][]byte{}
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	for _, row := range firstRulesTransform(grid) {
		ansPart1 += strings.Count(string(row), "#")
	}
	for _, row := range secondRulesTransform(grid) {
		ansPart2 += strings.Count(string(row), "#")
	}

	return
}

func firstRulesTransform(grid [][]byte) [][]byte {
	chaos := true
	for chaos {
		chaos = false
		newGrid := [][]byte{}
		for y, row := range grid {
			newRow := make([]byte, len(row))
			for x, b := range row {
				if b == 'L' || b == '#' {
					nOccupiedAdj := 0
					for i := 0; i < 3; i++ {
						for j := 0; j < 3; j++ {
							if i == j && j == 1 {
								continue
							}
							yAdj, xAdj := y+(i-1), x+(j-1)
							if yAdj < 0 || xAdj < 0 || yAdj == len(grid) || xAdj == len(row) {
								continue
							}
							if grid[yAdj][xAdj] == '#' {
								nOccupiedAdj++
							}
						}
					}
					if b == 'L' && nOccupiedAdj == 0 || b == '#' && nOccupiedAdj >= 4 {
						chaos = true
						if b == 'L' {
							b = '#'
						} else {
							b = 'L'
						}
					}
				}
				newRow[x] = b
			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid
	}
	return grid
}

func secondRulesTransform(grid [][]byte) [][]byte {
	chaos := true
	for chaos {
		chaos = false
		newGrid := [][]byte{}
		for y, row := range grid {
			newRow := make([]byte, len(row))
			for x, b := range row {
				if b == 'L' || b == '#' {
					nOccupiedAdj := 0
					for i := 0; i < 3; i++ {
						for j := 0; j < 3; j++ {
							if i == j && j == 1 {
								continue
							}
							dirY, dirX := i-1, j-1
							yAdj, xAdj := y+dirY, x+dirX
							for yAdj >= 0 && xAdj >= 0 && yAdj < len(grid) && xAdj < len(row) {
								bAdj := grid[yAdj][xAdj]
								if bAdj == '#' {
									nOccupiedAdj++
									break
								} else if bAdj == 'L' {
									break
								}
								yAdj, xAdj = yAdj+dirY, xAdj+dirX
							}
						}
					}
					if b == 'L' && nOccupiedAdj == 0 || b == '#' && nOccupiedAdj >= 5 {
						chaos = true
						if b == 'L' {
							b = '#'
						} else {
							b = 'L'
						}
					}
				}
				newRow[x] = b
			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid
	}
	return grid
}
