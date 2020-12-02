package day1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day1 = 1

// Response _
func Response() string {

	inputSlice := parseIntSlice(inputs.FetchInput(day1))
	sort.IntSlice(inputSlice).Sort()

	size := len(inputSlice)

	ansPart1 := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if xi, xj := inputSlice[i], inputSlice[j]; xi+xj == 2020 {
				ansPart1 = xi * xj
				break
			}
		}
	}

	ansPart2 := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if xi, xj, xk := inputSlice[i], inputSlice[j], inputSlice[k]; xi+xj+xk == 2020 {
					ansPart2 = xi * xj * xk
					break
				}
			}
		}
	}

	ansRecPart1 := recursionResponse(inputSlice, 2)
	ansRecPart2 := recursionResponse(inputSlice, 3)

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
		"\t and with recursion ",
		"part1: ", ansRecPart1, "\tpart2: ", ansRecPart2,
	)
}

func recursionResponse(input []int, n int) (prod int) {
	_, prod = recursion(input, 0, n, -1, 0, 1)
	return
}

func recursion(input []int, depth, maxDepth, pvIndex, pvSum, pvProd int) (sum, prod int) {
	if depth >= maxDepth || maxDepth <= 0 {
		return pvSum, pvProd
	}
	for index := pvIndex + 1; index < len(input); index++ {
		x := input[index]
		sum, prod = recursion(input, depth+1, maxDepth, index, pvSum+x, pvProd*x)
		if sum == 2020 {
			return sum, prod
		}
	}
	return 0, 0
}

func parseIntSlice(input io.Reader) (slice []int) {
	if ic, ok := input.(io.ReadCloser); ok {
		defer ic.Close()
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		slice = append(slice, n)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return slice
}
