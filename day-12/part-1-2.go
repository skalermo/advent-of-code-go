package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	east = iota
	south
	west
	north
)

func parseLine(line string) (byte, int) {
	navInstruction := line[0]
	value, _ := strconv.Atoi(line[1:])
	return navInstruction, value
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1(lines []string) {
	// for east, south, west, north respectively
	navPos := make([]int, 4)

	curDir := east

	for _, line := range lines {
		nav, val := parseLine(line)
		switch nav {
		case 'F':
			navPos[curDir] += val
			break
		case 'L':
			curDir = (curDir - val/90 + 4) % 4
			break
		case 'R':
			curDir = (curDir + val/90) % 4
			break
		case 'E':
			navPos[east] += val
			break
		case 'S':
			navPos[south] += val
			break
		case 'W':
			navPos[west] += val
			break
		case 'N':
			navPos[north] += val
			break
		}
	}
	manhattanDist := abs(navPos[east]-navPos[west]) + abs(navPos[north]-navPos[south])
	println(manhattanDist)
}

func rotateClockwise(x int, y int, times int) (int, int) {
	// rotates waypoint pos around ship clockwise by 90 degrees
	newX := x
	newY := y
	for i := 0; i < times; i++ {
		newX, newY = -newY, newX
	}
	return newX, newY
}

func part2(lines []string) {
	shipPosX := 0
	shipPosY := 0
	wayPointX := 10
	wayPointY := -1

	for _, line := range lines {
		nav, val := parseLine(line)
		switch nav {
		case 'F':
			shipPosX += wayPointX * val
			shipPosY += wayPointY * val
			break
		case 'L':
			times := (4 - val/90) % 4
			wayPointX, wayPointY = rotateClockwise(wayPointX, wayPointY, times)
			break
		case 'R':
			times := (val / 90) % 4
			wayPointX, wayPointY = rotateClockwise(wayPointX, wayPointY, times)
			break
		case 'E':
			wayPointX += val
			break
		case 'S':
			wayPointY += val
			break
		case 'W':
			wayPointX -= val
			break
		case 'N':
			wayPointY -= val
			break
		}
	}
	manhattanDist := abs(shipPosX) + abs(shipPosY)
	println(manhattanDist)
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
