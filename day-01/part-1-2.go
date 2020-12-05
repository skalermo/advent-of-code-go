package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(nums []int) {
outer:
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j && num1+num2 == 2020 {
				println(num1 * num2)
				break outer
			}

		}
	}
}

func part2(nums []int) {
outer:
	for i, num1 := range nums {
		for j, num2 := range nums {
			for k, num3 := range nums {
				if i != j && j != k && num1+num2+num3 == 2020 {
					println(num1 * num2 * num3)
					break outer
				}
			}
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	numStrings := lines[:len(lines)-1]
	var nums = []int{}
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	part1(nums)
	part2(nums)
}
