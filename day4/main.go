package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func removeSymbolRange(lower, upper int, symbol, part string) bool {

	removed := strings.ReplaceAll(part, symbol, "")
	p, _ := strconv.Atoi(removed)
	if p >= lower && p <= upper {
		return true
	}
	return false
}

func isInRangeDigits(digits, lower, upper int, part string) bool {
	if len(part) != digits {
		return false
	}
	p, _ := strconv.Atoi(part)
	if p >= lower && p <= upper {
		return true
	}
	return false
}

func check2(input string) bool {

	processed := strings.ReplaceAll(input, "\n", " ")
	tokens := strings.Split(processed, " ")
	count := 0

	for _, tkn := range tokens {
		parts := strings.Split(tkn, ":")

		switch parts[0] {
		case "byr":
			if isInRangeDigits(4, 1920, 2002, parts[1]) {
				count++
				continue
			}
		case "iyr":
			if isInRangeDigits(4, 2010, 2020, parts[1]) {
				count++
				continue
			}
		case "eyr":
			if isInRangeDigits(4, 2020, 2030, parts[1]) {
				count++
				continue
			}
		case "hgt":
			if match, err := regexp.Match("^\\d+cm$", []byte(parts[1])); err == nil && match {
				if removeSymbolRange(150, 193, "cm", parts[1]) {
					count++
					continue
				}
			}
			if match, err := regexp.Match("^\\d+in$", []byte(parts[1])); err == nil && match {
				if removeSymbolRange(59, 76, "in", parts[1]) {
					count++
					continue
				}
			}
		case "hcl":
			if match, err := regexp.Match("^#([0-9a-f]){6}$", []byte(parts[1])); err == nil && match {
				count++
				continue
			}
		case "ecl":
			if match, err := regexp.Match("^(amb|blu|brn|gry|grn|hzl|oth)$", []byte(parts[1])); err == nil && match {
				count++
				continue
			}
		case "pid":
			if match, err := regexp.Match("^\\d{9}$", []byte(parts[1])); err == nil && match {
				count++
				continue
			}
		}
	}

	return count == 7
}

func main() {

	check := func(input string, criteria ...string) int {
		for _, crit := range criteria {
			if !strings.Contains(input, crit) {
				return 0
			}
		}
		return 1
	}

	fmt.Println("Advent of Code Day 4")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")
		valid := 0
		for _, pp := range input {
			valid += check(pp, "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
		}

		fmt.Printf("Number of valid passports: %d\n", valid)

		fmt.Println("--- Part 2 ---")
		valid2 := 0
		for _, pp := range input {
			if check2(pp) {
				valid2++
			}
		}

		fmt.Printf("Number of valid passports: %d\n", valid2)
	}
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
