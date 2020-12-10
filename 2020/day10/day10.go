package day10

import (
	"bufio"
	"fmt"
	"io"
	"sort"

	"github.com/gabrielradureaupw/adventofcode/2020/inputs"
)

const day10 = 10

// Response _
func Response() string {

	ansPart1, ansPart2 := Day10(inputs.FetchInput(day10))

	return fmt.Sprint(
		"part1: ", ansPart1, "\tpart2: ", ansPart2,
	)
}

// Day10 _
func Day10(r io.Reader) (ansPart1, ansPart2 int) {
	scanner := bufio.NewScanner(r)
	ints := []int{}
	for scanner.Scan() {
		n := 0
		fmt.Sscanf(scanner.Text(), "%d", &n)
		ints = append(ints, n)
	}
	ints = append(ints, 0)
	sort.IntSlice(ints).Sort()

	diff1, diff3 := 0, 1 // +1 diff3 for the difference between the device and the last adapter
	last := 0            // from the charging outlet
	for _, n := range ints {
		diff := n - last
		switch diff {
		case 1:
			diff1++
		case 3:
			diff3++
		}
		last = n
	}
	ansPart1 = diff1 * diff3

	ansPart2 = build(0, ints, map[int]treeNode{}).npath(map[int]int{})
	return
}

type treeNode struct {
	value    int
	children []treeNode
}

func build(i int, ints []int, nodes map[int]treeNode) (tn treeNode) {
	if node, ok := nodes[ints[i]]; ok {
		return node
	}
	tn.value = ints[i]
	for j := i + 1; j <= i+3 && j < len(ints); j++ {
		next := ints[j]
		if next-tn.value > 3 {
			break
		}
		tn.children = append(tn.children, build(j, ints, nodes))
	}
	nodes[tn.value] = tn
	return
}

func (tn treeNode) npath(npath map[int]int) (n int) {
	if len(tn.children) == 0 {
		return 1
	}
	if n, ok := npath[tn.value]; ok {
		return n
	}
	for _, child := range tn.children {
		n += child.npath(npath)
	}
	npath[tn.value] = n
	return n
}
