package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bag struct {
	ID       string
	Contains []BagAmount
}

func (b Bag) ContainsBagsAmount(bags []Bag) int {
	count := 1
	for _, ba := range b.Contains {

		bb := findBag(bags, ba.Bag.ID)

		count += ba.Amount * bb.ContainsBagsAmount(bags)
	}
	return count
}

type BagAmount struct {
	Bag    Bag
	Amount int
}
type BagHolder struct {
	BH map[string]bool
}

func findBag(bags []Bag, id string) Bag {

	for _, b := range bags {
		if b.ID == id {
			return b
		}
	}

	return Bag{
		ID: "",
	}
}

func (b Bag) String() string {
	var sb strings.Builder

	sb.WriteString("Bag [ID= ")
	sb.WriteString(b.ID)
	sb.WriteString("]: \n")

	if len(b.Contains) > 0 {
		for _, b := range b.Contains {
			sb.WriteString(fmt.Sprintf(" - %dx %s\n", b.Amount, b.Bag.ID))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {

	fmt.Println("Advent of Code Day 7")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		var bags []Bag
		bagHolder := BagHolder{
			BH: make(map[string]bool),
		}
		for _, in := range input {
			bags = append(bags, parseBag(in))
		}

		for _, b := range bags {
			fmt.Println(b)
		}

		findBagHolders("shiny gold", bags, &bagHolder)

		for k, _ := range bagHolder.BH {
			fmt.Printf("A [%s] bag can hold a [%s] bag\n", k, "shiny gold")
		}

		fmt.Printf("Overall there are %d bags that can hold the bag\n", len(bagHolder.BH))

		fmt.Println("--- Part 2 ---")

		for _, b := range bags {
			if b.ID == "shiny gold" {
				fmt.Printf("A single shiny gold bag contains %d bags\n", b.ContainsBagsAmount(bags)-1)
			}
		}

	}
}

func findBagHolders(s string, bags []Bag, bh *BagHolder) {

	for _, b := range bags {

		for _, c := range b.Contains {
			if c.Bag.ID == s {
				bh.BH[b.ID] = true
				findBagHolders(b.ID, bags, bh)
			}

		}

	}

}

func getBagFrom(in string) (Bag, int) {

	b := Bag{}
	amount := -1

	if strings.Contains(in, "bags") || strings.Contains(in, "bag") {
		split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(in, " "), " "), " ")

		if len(split) == 4 {
			// is first part a number?
			if i, err := strconv.Atoi(split[0]); err == nil {
				amount = i
				b.ID = fmt.Sprintf("%s %s", split[1], split[2])
			}
		}
		if len(split) == 3 {
			b.ID = fmt.Sprintf("%s %s", split[0], split[1])
		} else {
			fmt.Errorf("Split is not three parts but %d", len(split))
		}
	}

	return b, amount
}

func parseBag(in string) Bag {

	splitted := strings.Split(in, "contain")
	outerBag, _ := getBagFrom(splitted[0])

	if strings.Contains(splitted[1], "no other bags") {

	} else {
		second := strings.ReplaceAll(splitted[1], ".", "")
		second = strings.ReplaceAll(second, "\n", ",")
		bagsInput := strings.Split(second, ",")

		for _, bi := range bagsInput {
			b, amount := getBagFrom(bi)
			outerBag.Contains = append(outerBag.Contains, BagAmount{b, amount})
		}
	}

	return outerBag
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
