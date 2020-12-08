package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Bag struct {
	ID       string
	Contains []Bag
}

func NewBag(id string, bags... Bag) Bag {
	return Bag{
		ID:       id,
		Contains: []Bag{bags},
	}
}

func main() {

	fmt.Println("Advent of Code Day 6")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")
		bags := []Bag{}

	}
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
