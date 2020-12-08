package day8

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day8 = 8

// Response _
func Response() string {

	ansPart1, ansPart2 := Day8(inputs.FetchInput(day8))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

const (
	nop = "nop"
	acc = "acc"
	jmp = "jmp"
)

// Day8 _
func Day8(r io.Reader) (ansPart1, ansPart2 int) {
	b, _ := ioutil.ReadAll(r)
	program := strings.Split(string(b), "\n")

	runProgram := func(program []string) (res int, looped bool) {
		pc := 0
		visited := []int{}
		for pc < len(program) {
			instruction := program[pc]
			visited = append(visited, pc)
			op, arg := "", 0
			fmt.Sscanf(instruction, "%s %d", &op, &arg)
			switch op {
			case jmp:
				pc += arg
				for _, pvPc := range visited {
					if pc == pvPc {
						return res, true
					}
				}
			case acc:
				res += arg
				fallthrough
			default:
				pc++
			}
		}
		return
	}

	ansPart1, _ = runProgram(program)

	looped := true
	for i, instruction := range program {
		op, arg := "", 0
		fmt.Sscanf(instruction, "%s %d", &op, &arg)
		if op == jmp || op == nop {
			copyPgrm := make([]string, len(program))
			copy(copyPgrm, program)
			if op == jmp {
				copyPgrm[i] = strings.Replace(copyPgrm[i], jmp, nop, 1)
			} else {
				copyPgrm[i] = strings.Replace(copyPgrm[i], nop, jmp, 1)
			}
			ansPart2, looped = runProgram(copyPgrm)
			if !looped {
				break
			}
		}
	}

	return
}
