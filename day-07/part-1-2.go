package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type containeeInfo struct {
	quantity int
	color    string
}

type containee struct {
	bag      *bag
	quantity int
}

type bag struct {
	containers    []*bag
	containees    []containee
	containeesSum int
	color         string
}

func parseLine(line string) (string, int, []containeeInfo) {
	words := strings.Split(line, " ")
	containerColor := strings.Join(words[0:2], " ")
	if words[4] == "no" {
		return containerColor, 0, nil
	}
	containeesInfo := make([]containeeInfo, 0)
	containeesSum := 0
	for i := 4; i < len(words); i += 4 {
		quantity, _ := strconv.Atoi(words[i])
		containeesSum += quantity
		color := strings.Join(words[i+1:i+3], " ")
		containeesInfo = append(containeesInfo, containeeInfo{quantity, color})
	}
	return containerColor, containeesSum, containeesInfo
}

func createContainGraph(lines []string) *bag {
	colorBagMap := make(map[string]*bag)
	for _, line := range lines {
		containerColor, containeesSum, containeesInfo := parseLine(line)

		container, ok := colorBagMap[containerColor]
		if !ok {
			container = &bag{nil, nil, containeesSum, containerColor}
			colorBagMap[containerColor] = container
		} else {
			container.containeesSum = containeesSum
		}

		containees := make([]containee, len(containeesInfo))
		for i, cInfo := range containeesInfo {

			b, ok := colorBagMap[cInfo.color]
			if !ok {
				b = &bag{[]*bag{container}, nil, 0, cInfo.color}
				colorBagMap[cInfo.color] = b
			} else {
				b.containers = append(b.containers, container)
			}
			containees[i] = containee{b, cInfo.quantity}
		}
		container.containees = append(container.containees, containees...)
	}
	// return pointer to shiny gold bag
	return colorBagMap["shiny gold"]
}

func part1(lines []string) {
	shinyGoldBag := createContainGraph(lines)
	closedBags := make(map[string]bool)
	bagStack := make([]*bag, 0)
	for _, c := range shinyGoldBag.containers {
		bagStack = append(bagStack, c)
	}
	for i := 0; i < len(bagStack); i++ {
		closedBags[bagStack[i].color] = true
		for _, c := range bagStack[i].containers {
			if _, exists := closedBags[c.color]; !exists {
				bagStack = append(bagStack, c)
			}
		}
	}
	println(len(closedBags))
}

func part2(lines []string) {
	shinyGoldBag := createContainGraph(lines)
	bagStack := make([]containee, 0)

	bagsTotal := shinyGoldBag.containeesSum
	for _, c := range shinyGoldBag.containees {
		bagStack = append(bagStack, c)
	}
	for i := 0; i < len(bagStack); i++ {
		bagsTotal += bagStack[i].quantity * bagStack[i].bag.containeesSum
		for _, c := range bagStack[i].bag.containees {
			c.quantity *= bagStack[i].quantity
			bagStack = append(bagStack, c)
		}
	}
	println(bagsTotal)
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
