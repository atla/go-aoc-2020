package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	fmt.Println("Advent of Code Day 5")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		seats := []Seat{}
		maxID := -1
		for _, code := range input {
			row := findRow(code[:7])
			col := findColumn(code[7:])
			id := row*8 + col
			seat := Seat{
				Code: code,
				Row:  row,
				Col:  col,
				ID:   id,
			}
			seats = append(seats, seat)
			maxID = max(maxID, id)

			fmt.Println(seat)
		}

		fmt.Printf("The seat with the highest ID is: %d\n", maxID)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func crawlRow(in string, lower, upper, rows int) int {

	if len(in) == 1 {
		if in == "F" {
			return lower
		}
		return upper
	}
	rows /= 2

	if string(in[0]) == "F" {
		upper -= rows
	} else {
		lower += rows
	}

	return crawlRow(in[1:], lower, upper, rows)
}

func findRow(in string) int {
	return crawlRow(in, 0, 127, 128)
}

func findColumn(in string) int {
	return crawlColumn(in, 0, 7, 8)
}

func crawlColumn(in string, lower, upper, seats int) int {

	if len(in) == 1 {
		if in == "L" {
			return lower
		}
		return upper
	}
	seats /= 2

	if string(in[0]) == "L" {
		upper -= seats
	} else {
		lower += seats
	}

	return crawlColumn(in[1:], lower, upper, seats)
}

type Seat struct {
	Code     string
	Row, Col int
	ID       int
}

func (s Seat) String() string {
	return fmt.Sprintf("%s: row %d, column %d, seat ID %d.", s.Code, s.Row, s.Col, s.ID)
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
