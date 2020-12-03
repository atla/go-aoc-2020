package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type PasswordCheck struct {
	lowerBound int
	upperBound int
	character  string
	password   string
}

type Matcher func(check *PasswordCheck) bool

func parsePasswordCheck(input string) (*PasswordCheck, error) {
	splitted := strings.Split(input, ":")
	if len(splitted) != 2 {
		return nil, errors.New("Could not parse input")
	}

	criteria := strings.Split(splitted[0], " ")
	bounds := strings.Split(criteria[0], "-")
	lowerBound, _ := strconv.Atoi(bounds[0])
	upperBound, _ := strconv.Atoi(bounds[1])

	return &PasswordCheck{
		lowerBound,
		upperBound,
		criteria[1],
		strings.ReplaceAll(splitted[1], " ", ""),
	}, nil
}

func matchesCriteria1(pwc *PasswordCheck) bool {
	count := strings.Count(pwc.password, pwc.character)
	return count >= pwc.lowerBound && count <= pwc.upperBound
}

func matchesCriteria2(pwc *PasswordCheck) bool {
	matched := 0
	for _, pos := range []int{pwc.lowerBound, pwc.upperBound} {
		if string(pwc.password[pos-1]) == pwc.character {
			matched++
		}
	}
	return matched == 1
}

func checkPassword(matcher Matcher) {
	validPasswords := 0

	if input, err := readInput("input.txt"); err == nil {
		for _, line := range input {
			if pwc, err := parsePasswordCheck(line); err == nil {
				if matcher(pwc) {
					validPasswords++
				}
			}
		}
	}
	fmt.Printf("Number of valid passwords %d\n", validPasswords)
}

func main() {
	fmt.Println("Advent of Code Day 2")

	fmt.Println("--- Part 1 ---")
	checkPassword(matchesCriteria1)

	fmt.Println("--- Part 2 ---")
	checkPassword(matchesCriteria2)

}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
