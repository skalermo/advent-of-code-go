package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part1(numbers []int) {
	joltDiffs := make(map[int]int)
	joltDiffs[1] = 0
	joltDiffs[3] = 0

	joltDiffs[numbers[0]-0]++
	for i := 0; i < len(numbers)-1; i++ {
		joltDiffs[numbers[i+1]-numbers[i]]++
	}
	joltDiffs[3]++
	println(joltDiffs[1] * joltDiffs[3])

}

func part2(numbers []int) {
	numbers = append([]int{0}, numbers...)
	totalWays := make([]int, len(numbers))

	totalWays[0] = 1
	for i := 0; i < len(numbers); i++ {
		if i+3 < len(numbers) && numbers[i+3]-numbers[i] <= 3 {
			totalWays[i+3] += totalWays[i]
		}
		if i+2 < len(numbers) && numbers[i+2]-numbers[i] <= 3 {
			totalWays[i+2] += totalWays[i]
		}
		if i+1 < len(numbers) {
			totalWays[i+1] += totalWays[i]
		}
	}
	println(totalWays[len(totalWays)-1])
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

	sort.Ints(numbers)

	part1(numbers)
	part2(numbers)
}
