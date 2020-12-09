package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(numbers []int) int {
	n := 25
	lastNumbers := make(map[int]bool)
	for i := 0; i < n; i++ {
		lastNumbers[numbers[i]] = true
	}

outer:
	for i := n; i < len(numbers); i++ {
		curNum := numbers[i]
		for j := i - n; j < i; j++ {
			if _, exists := lastNumbers[curNum-numbers[j]]; exists {
				delete(lastNumbers, numbers[i-n])
				lastNumbers[curNum] = true
				continue outer
			}
		}
		return curNum
	}
	return 0
}

func part2(numbers []int, part1Answer int) {
	startOfWindow := 0
	endOfWindow := 2
	var contiguousSet []int

outer:
	for endOfWindow < len(numbers) {
		curSumOfWindow := 0
		for _, num := range numbers[startOfWindow:endOfWindow] {
			curSumOfWindow += num
		}
		if curSumOfWindow > part1Answer {
			startOfWindow++
		} else if curSumOfWindow < part1Answer {
			endOfWindow++
		} else {
			contiguousSet = numbers[startOfWindow:endOfWindow]
			break outer
		}
	}

	min := contiguousSet[0]
	max := contiguousSet[0]
	for _, num := range contiguousSet {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	println(min + max)
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	numbers := make([]int, len(lines))
	for i := range lines {
		numbers[i], _ = strconv.Atoi(lines[i])
	}

	part1Answer := part1(numbers)
	println(part1Answer)
	part2(numbers, part1Answer)
}
