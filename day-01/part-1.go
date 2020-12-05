package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	numStrings := strings.Split(string(input), "\n")
	var nums = []int{}
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

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
