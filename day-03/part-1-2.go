package main

import (
	"io/ioutil"
	"strings"
)

func countTrees(lines []string, right int, down int) int {
	// assuming all lines have same length
	lineLength := len(lines[0])

	treesEncountered := 0
	for lineNo, pos := down, right; lineNo < len(lines); lineNo, pos = lineNo+down, (pos+right)%lineLength {
		if lines[lineNo][pos] == '#' {
			treesEncountered++
		}
	}
	return treesEncountered
}

func part1(lines []string) {
	println(countTrees(lines, 3, 1))
}

func part2(lines []string) {
	inputs := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	multiplied := 1
	for _, i := range inputs {
		multiplied *= countTrees(lines, i[0], i[1])
	}
	println(multiplied)
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	part1(lines)
	part2(lines)
}
