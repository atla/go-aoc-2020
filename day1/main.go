package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code Day 1")
	fmt.Println("--- Part 1 ---")

	input := readInput("input.txt")

	if a, b, err := findEntriesMatchingSum(input, 2020); err == nil {
		fmt.Printf("Found matching entries %d %d\n", a, b)
		fmt.Println("Result is: ", a*b)
	}

	fmt.Println("--- Part 2 ---")

	if a, b, c, err := findMatching(input, 2020); err == nil {
		fmt.Printf("Found matching entries %d %d %d\n", a, b, c)
		fmt.Println("Result is: ", a*b*c)
	}
}

func readInput(file string) []int {
	var numbers []int

	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		splitted := strings.Split(input, "\n")

		for _, number := range splitted {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func findEntriesMatchingSum(input []int, target int) (int, int, error) {

	if len(input) <= 1 {
		return -1, -1, errors.New("Could not find matching sum")
	}

	for _, value := range input[1:] {
		if input[0]+value == target {
			return input[0], value, nil
		}
	}
	return findEntriesMatchingSum(input[1:], target)
}

func findMatching(input []int, target int) (int, int, int, error) {

	for i, e1 := range input {
		for e, e2 := range input {
			for k, e3 := range input {
				if i != e && e != k && i != k {
					if e1+e2+e3 == target {
						return e1, e2, e3, nil
					}
				}
			}
		}
	}
	return -1, -1, -1, errors.New("No matching result")

}
