package main

import (
	"io/ioutil"
	"strings"
)

func part1(lines []string) {
	answersCount := 0
	answeredQuestions := make([]int, 26)
	for i := range answeredQuestions {
		answeredQuestions[i] = 0
	}
	for _, line := range lines {
		if len(line) == 0 {
			for i := range answeredQuestions {
				answersCount += answeredQuestions[i]
				answeredQuestions[i] = 0
			}

		} else {
			// assuming r is from range [a-z]
			for _, r := range line {
				answeredQuestions[r-'a'] = 1
			}
		}
	}
	println(answersCount)
}

func part2(lines []string) {
	answersCount := 0
	answeredQuestions := make([]int, 26)
	for i := range answeredQuestions {
		answeredQuestions[i] = 0
	}

	peopleInGroup := 0
	for _, line := range lines {
		if len(line) == 0 {
			for i := range answeredQuestions {
				if answeredQuestions[i] == peopleInGroup {
					answersCount++
				}
				answeredQuestions[i] = 0
			}
			peopleInGroup = 0
		} else {
			peopleInGroup++

			// assuming r is from range [a-z]
			for _, r := range line {
				answeredQuestions[r-'a']++
			}
		}
	}
	println(answersCount)
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
