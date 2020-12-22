package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type AllergenCandidate struct {
	Name        string
	Ingredients []string
}

type IngredientCounter map[string]int

func (a AllergenCandidate) ContainsIngredient (s string) bool {
	for _, i := range a.Ingredients {
		if i == s {
			return true
		}
	}
	return false
}

func (a AllergenCandidate) RemoveIngredient(s string) {
	for i, ing := range a.Ingredients {
		if ing == s {
			a.Ingredients = append(a.Ingredients[:i], a.Ingredients[i:]...)
			return
		}
	}
}

func main() {

	fmt.Println("Advent of Code Day 21")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		allergens := map[string]AllergenCandidate{}
		counter := IngredientCounter{}

		for _, in := range input {

			ing, allg := processInput(in)

			for _, i := range ing {
				if _, ok := counter[i]; ok {
					counter[i]++
				} else {
					counter[i] = 1
				}
			}

			for _, a := range allg {
				if allergen, ok := allergens[a]; ok {

					// remove existing ingredients from allergen
					for _, oing := range allergen.Ingredients {
						found := false
						for _, ing2 := range ing {
							if ing2 == oing {
								found = true
							}
						}
						if !found {
							allergen.RemoveIngredient(oing)
						}
					}
				} else {
					allergens[a] = AllergenCandidate{
						Name:        a,
						Ingredients: ing,
					}
				}
			}
		}

		fmt.Printf(" -- Part 2 --")

	}
}

func processInput(in string) ([]string, []string) {

	splitted := strings.Split(in, " (contains ")

	return strings.Split(splitted[0], " "), strings.Split(strings.ReplaceAll(splitted[1], ")", ""), ",")
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
