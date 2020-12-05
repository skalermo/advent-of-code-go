package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func passwordFrom(line string) (int, int, byte, string) {
	x := strings.Split(line, " ")
	lowAndHigh, letter, password := strings.Split(x[0], "-"), x[1][0], x[2]
	low, _ := strconv.Atoi(lowAndHigh[0])
	high, _ := strconv.Atoi(lowAndHigh[1])

	return low, high, letter, password
}

func validate(low int, high int, letter byte, password string) bool {
	occurrences := strings.Count(password, string(letter))
	return low <= occurrences && occurrences <= high
}

func validate2(pos1 int, pos2 int, letter byte, password string) bool {
	pos1--
	pos2--
	if password[pos1] == letter && password[pos2] == letter {
		return false
	}
	if password[pos1] == letter || password[pos2] == letter {
		return true
	}
	return false
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	valid := 0
	for _, line := range lines {
		// if validate(passwordFrom(line)) {
		if validate2(passwordFrom(line)) {
			valid++
		}
	}

	println(valid)
}
