package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	instrType string
	argument  int
	line      int
}

func parseInstruction(line string, lineNo int) instruction {
	words := strings.Split(line, " ")
	arg, _ := strconv.Atoi(words[1])
	return instruction{words[0], arg, lineNo}
}

func executeInstruction(instr instruction, acc *int, nextInstr *int) {
	if instr.instrType == "acc" {
		*acc += instr.argument
		*nextInstr++
	} else if instr.instrType == "jmp" {
		*nextInstr += instr.argument
	} else if instr.instrType == "nop" {
		*nextInstr++
	}
}

func changeInstruction(instr *instruction) {
	if instr.instrType == "jmp" {
		instr.instrType = "nop"
	} else if instr.instrType == "nop" {
		instr.instrType = "jmp"
	}
}

func part1(lines []string) {
	instructions := make([]instruction, len(lines))
	for i := range lines {
		instructions[i] = parseInstruction(lines[i], i)
	}

	acc, _, _ := execute(instructions, -1, false)
	println(acc)
}

func part2(lines []string) {
	instructions := make([]instruction, len(lines))
	for i := range lines {
		instructions[i] = parseInstruction(lines[i], i)
	}

	_, _, jmpsAndNops := execute(instructions, -1, true)
	for _, instr := range jmpsAndNops {
		if acc, terminatedOk, _ := execute(instructions, instr.line, false); terminatedOk {
			println(acc)
			break
		}
	}
}

func execute(instructions []instruction, iToChange int, debug bool) (int, bool, []instruction) {
	acc := 0
	nextInstr := 0
	jmpsAndNops := make([]instruction, 0)
	canaries := make([]bool, len(instructions))
	for i := 0; i < len(canaries); i++ {
		canaries[i] = false
	}

	for true {
		if canaries[nextInstr] {
			break
		} else {
			canaries[nextInstr] = true
		}
		instr := instructions[nextInstr]
		if debug && (instr.instrType == "jmp" || instr.instrType == "nop") {
			jmpsAndNops = append(jmpsAndNops, instr)
		}

		if iToChange == nextInstr {
			changeInstruction(&instr)
		}
		executeInstruction(instr, &acc, &nextInstr)
		if nextInstr == len(instructions) {
			return acc, true, nil
		}
	}
	return acc, false, jmpsAndNops
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
