package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	fmt.Println("Advent of Code Day 6")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		count := 0

		for _, groupInput := range input {
			count += countGroup(groupInput)
		}

		fmt.Printf("Sum of questions is %d: \n", count)

		fmt.Println("--- Part 2 ---")

		count2 := 0

		for _, groupInput := range input {
			count2 += countGroup2(groupInput)
		}

		fmt.Printf("Sum of questions is %d: \n", count2)
	}
}

func countGroup2(input string) int {

	m := map[string]int{}
	pInGroup := 1 + strings.Count(input, "\n")
	input = strings.ReplaceAll(input, "\n", "")

	for _, i := range input {

		id := string(i)
		if _, contains := m[id]; contains {
			m[id] ++
		} else {
			m[id] = 1
		}

	}
	count := 0
	for _, v := range m {
		if v == pInGroup {
			count++
		}
	}

	return count
}

func countGroup(input string) int {

	m := map[string]int{}

	input = strings.ReplaceAll(input, "\n", "")

	for _, i := range input {
		m[string(i)] = 1
	}

	return len(m)
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
