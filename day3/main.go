package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	data          []string
	width, height int
}

type Toboggan struct {
	x, y    int
	slope   func(int, int) (int, int)
	counter func(string) int
}

func (m *Map) pos(x, y int) string {
	return string(m.data[y][x%m.width])
}

func part1(input []string) {
	m := Map{
		data:   input,
		width:  len(input[0]),
		height: len(input),
	}
	t := Toboggan{
		x: 0,
		y: 0,
		slope: func(x, y int) (int, int) {
			return x + 3, y + 1
		},
		counter: func(pos string) int {
			if pos == "#" {
				return 1
			}
			return 0
		},
	}

	trees := 0
	for t.y < m.height {
		trees += t.counter(m.pos(t.x, t.y))
		t.x, t.y = t.slope(t.x, t.y)
	}
	fmt.Printf("Encountered trees %d\n", trees)
}

type Slope struct {
	x, y int
}

func part2(input []string) {
	m := Map{
		data:   input,
		width:  len(input[0]),
		height: len(input),
	}

	counter := func(pos string) int {
		if pos == "#" {
			return 1
		}
		return 0
	}

	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	trees := []int{0, 0, 0, 0, 0}

	for i, s := range slopes {

		t := Toboggan{
			x: 0,
			y: 0,
			slope: func(x, y int) (int, int) {
				return x + s.x, y + s.y
			},
			counter: counter,
		}

		for t.y < m.height {
			trees[i] += t.counter(m.pos(t.x, t.y))
			t.x, t.y = t.slope(t.x, t.y)
		}
	}

	product := trees[0]

	for _, v := range trees[1:] {
		product *= v
	}

	fmt.Printf("Product of encountered trees %d\n", product)
}

func main() {
	fmt.Println("Advent of Code Day 3")
	if input, err := readInput("input.txt"); err == nil {
		fmt.Println("--- Part 1 ---")
		part1(input)

		fmt.Println("--- Part 2 ---")
		part2(input)
	}
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
