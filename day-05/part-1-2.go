package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func binaryToInt(binaryStr string) int {
	i, _ := strconv.ParseInt(binaryStr, 2, 64)
	return int(i)
}

func decodeSeatID(encoded string) int {
	rowBinary := strings.ReplaceAll(encoded[:7], "F", "0")
	rowBinary = strings.ReplaceAll(rowBinary, "B", "1")

	colBinary := strings.ReplaceAll(encoded[7:], "L", "0")
	colBinary = strings.ReplaceAll(colBinary, "R", "1")
	row := binaryToInt(rowBinary)
	col := binaryToInt(colBinary)
	return row*8 + col

}

func part1(lines []string) {
	highestID := 0
	for _, line := range lines {
		if curID := decodeSeatID(line); highestID < curID {
			highestID = curID
		}
	}
	println(highestID)
}

func part2(lines []string) {
	lowestID := decodeSeatID(lines[0])
	highestID := lowestID
	xor := lowestID
	for _, line := range lines[1:] {
		curID := decodeSeatID(line)
		if highestID < curID {
			highestID = curID
		}
		if lowestID > curID {
			lowestID = curID
		}
		xor ^= curID
	}
	for i := lowestID; i <= highestID; i++ {
		xor ^= i
	}
	println(xor)
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
