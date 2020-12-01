package day1

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

// Response _
func Response() string {

	inputSlice := parseIntSlice(fetchInput())
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

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

func fetchInput() io.Reader {
	const inputURL = "https://adventofcode.com/2020/day/1/input"
	const session = "53616c7465645f5f75b77582051a3867a4aac617492d96e39cf78b7f19f6dc846df755d5472c952504abcf4ae859d161" // expires in nov 2030
	req, _ := http.NewRequest(http.MethodGet, inputURL, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return resp.Body
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
