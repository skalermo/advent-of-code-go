package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ':' || r == ' ' || r == '\n'
}

func validatePassport(passport []string) bool {
	fieldsChecked := 0

	for i := 0; i < len(passport); i += 2 {
		key := passport[i]

		switch key {
		case "byr":
			fieldsChecked++
			break
		case "iyr":
			fieldsChecked++
			break
		case "eyr":
			fieldsChecked++
			break
		case "hgt":
			fieldsChecked++
			break
		case "hcl":
			fieldsChecked++
			break
		case "ecl":
			fieldsChecked++
			break
		case "pid":
			fieldsChecked++
			break
		case "cid":
			break
		default:
			return false
		}
	}
	return fieldsChecked == 7
}
func validatePassport2(passport []string) bool {
	fieldsChecked := 0

	for i := 0; i < len(passport); i += 2 {
		key := passport[i]
		val := passport[i+1]

		switch key {
		case "byr":
			if validateByr(val) {
				fieldsChecked++
			}
			break
		case "iyr":
			if validateIyr(val) {
				fieldsChecked++
			}
			break
		case "eyr":
			if validateEyr(val) {
				fieldsChecked++
			}
			break
		case "hgt":
			if validateHgt(val) {
				fieldsChecked++
			}
			break
		case "hcl":
			if validateHcl(val) {
				fieldsChecked++
			}
			break
		case "ecl":
			validEcls := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, validEcl := range validEcls {
				if val == validEcl {
					fieldsChecked++
					break
				}
			}
			break
		case "pid":
			if validatePid(val) {
				fieldsChecked++
			}
			break
		case "cid":
			break
		default:
			return false
		}
	}
	return fieldsChecked == 7
}

func validateByr(str string) bool {
	num, _ := strconv.Atoi(str)
	return 1920 <= num && num <= 2002
}

func validateIyr(str string) bool {
	num, _ := strconv.Atoi(str)
	return 2010 <= num && num <= 2020
}

func validateEyr(str string) bool {
	num, _ := strconv.Atoi(str)
	return 2020 <= num && num <= 2030
}

func validateHgt(str string) bool {
	if len(str) < 3 {
		return false
	}
	val, err := strconv.Atoi(str[:len(str)-2])
	if err != nil {
		return false
	}
	suffix := str[len(str)-2:]
	switch suffix {
	case "cm":
		return 150 <= val && val <= 193
	case "in":
		return 59 <= val && val <= 76
	default:
		return false
	}
}

func validateHcl(str string) bool {
	var validHcl = regexp.MustCompile("^#[a-f0-9]{6}$")
	return bool(validHcl.MatchString(str))
}

func validatePid(str string) bool {
	if len(str) != 9 {
		return false
	}
	_, err := strconv.Atoi(str[:len(str)-2])
	return err == nil
}

func part1(chunks []string) {
	validPassports := 0
	for _, c := range chunks {
		keysAndVals := strings.FieldsFunc(c, split)
		if validatePassport(keysAndVals) {
			validPassports++
		}
	}
	println(validPassports)
}

func part2(chunks []string) {
	validPassports := 0
	for _, c := range chunks {
		keysAndVals := strings.FieldsFunc(c, split)
		if validatePassport2(keysAndVals) {
			validPassports++
		}
	}
	println(validPassports)
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	chunks := strings.Split(string(input), "\n\n")
	part1(chunks)
	part2(chunks)
}
