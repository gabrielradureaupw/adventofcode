package day9

import (
	"fmt"
	"io"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day9 = 9

// Response _
func Response() string {

	ansPart1, ansPart2 := Day9(inputs.FetchInput(day9), 25)

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

// Day9 _
func Day9(r io.Reader, preambleLength int) (ansPart1, ansPart2 int) {
	readInt := func() (n int) {
		fmt.Fscanf(r, "%d\n", &n)
		return
	}
	preamble := make([]int, preambleLength)
	sums := make(map[int][]int, preambleLength-1)
	for i := 0; i < preambleLength; i++ {
		n := readInt()
		for _, pv := range preamble[:i] {
			sums[n] = append(sums[n], n+pv)
		}
		preamble[i] = n
	}

	numbers := make([]int, preambleLength)
	copy(numbers, preamble)
	n := readInt()
	for n != 0 {
		valid := false
		for i := range sums {
			for _, sum := range sums[i] {
				if n == sum {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid && ansPart1 == 0 {
			ansPart1 = n
			break
		}

		numbers = append(numbers, n)
		for _, pv := range preamble {
			sums[n] = append(sums[n], n+pv)
		}

		delete(sums, preamble[0])
		preamble = append(preamble[1:], n)
		n = readInt()
	}
	for low := 0; low < len(numbers)-1; low++ {
		sum := numbers[low]
		min, max := sum, sum
		for high := low + 1; high < len(numbers); high++ {
			sum += numbers[high]
			if min > numbers[high] {
				min = numbers[high]
			} else if max < numbers[high] {
				max = numbers[high]
			}
			if sum == ansPart1 {
				ansPart2 = min + max
				break
			}
		}
		if ansPart2 != 0 {
			break
		}
	}
	return
}
