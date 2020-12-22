package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code Day 9")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		preambleCount := 25
		preamble := []int{}

		for _, v := range input {
			i, _ := strconv.Atoi(v)

			if len(preamble) == preambleCount {
				if valid(preamble, i) {
					preamble = append(preamble[1:], i)
				} else {
					fmt.Printf("%d doesnt match preamble rule\n", i)
					return
				}
			} else {
				preamble = append(preamble, i)
			}
		}
	}
}

func valid(preamble []int, inp int) bool {
	for i := len(preamble) - 1; i >= 0; i-- {
		for j := len(preamble) - 2; j >= 1; j-- {
			if preamble[i]+preamble[j] == inp {
				return true
			}
		}
	}
	return false
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
