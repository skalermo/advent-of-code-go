package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) {
	timestamp, _ := strconv.Atoi(lines[0])
	busIDs := strings.Split(lines[1], ",")
	toWait := timestamp
	var chosenBus int
	for _, busID := range busIDs {
		if id, err := strconv.Atoi(busID); err == nil {
			if timeToWait := id - timestamp%id; toWait > timeToWait {
				toWait = timeToWait
				chosenBus = id
			}
		}
	}
	println(toWait * chosenBus)
}

func part2(lines []string) {
	input := strings.Split(lines[1], ",")
	rs := []int{}
	ns := []int{}

	for i, n := range input {
		if n, err := strconv.Atoi(n); err == nil {
			r := n - i
			for r < 0 {
				r += n
			}
			rs = append(rs, r)
			ns = append(ns, n)
		}
	}

	// To find solution used search by sieving without sorting ns
	// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Search_by_sieving

	multipliedTotal := ns[0]
	s := rs[0]
	for i := 1; i < len(ns); i++ {
		for (s % ns[i]) != rs[i] {
			s += multipliedTotal
		}
		multipliedTotal *= ns[i]
	}
	println(s)
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	part1(lines)
	part2(lines)
}
