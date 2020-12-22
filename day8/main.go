package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Operation string

const (
	acc Operation = "acc"
	jmp           = "jmp"
	nop           = "nop"
)

type Instruction struct {
	Operation Operation
	Argument  int
	ID        int
}

func (i Instruction) String() string {
	return fmt.Sprintf("Executing Instruction ID: %d Operation: %s Argument: %d", i.ID, i.Operation, i.Argument)
}

type Program struct {
	Instructions []Instruction
}

type Console struct {
	Accumulator int
	Current     int
	History     []int
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (c *Console) Run(p Program) error {

	c.Current = 0
	c.History = []int{}
	c.Accumulator = 0
	running := true

	for running {
		i := p.Instructions[c.Current]

		if contains(c.History, i.ID) {
			return errors.New("found endless loop")
		}

		c.Current += c.Process(i)

		if c.Current > len(p.Instructions) {
			running = false
		}
	}

	return nil
}

func (p Program) FindInstruction(id int) Instruction {

	for _, ins := range p.Instructions {
		if ins.ID == id {
			return ins
		}
	}
	return Instruction{}
}

func (c *Console) RunAutoRepair(p Program) error {

	c.Current = 0
	c.History = []int{}
	c.Accumulator = 0
	running := true

	for running {
		i := p.Instructions[c.Current]

		if contains(c.History, i.ID) {

			secondToLast := c.History[len(c.History)-3]

			i = p.Instructions[secondToLast]
			c.Current = secondToLast

			if i.Operation == jmp {
				i.Operation = nop
			} else if i.Operation == nop {
				i.Operation = jmp
			}

		}

		c.Current += c.Process(i)

		if c.Current >= len(p.Instructions) {
			running = false
		}
	}
	fmt.Printf("Value in the accumulator is %d\n", c.Accumulator)

	return nil
}

// returns nextInstruction offset
func (c *Console) Process(is Instruction) int {

	//	fmt.Println(is)

	nextInstruction := 1

	switch is.Operation {
	case acc:
		c.Accumulator += is.Argument
	case jmp:
		nextInstruction = is.Argument
	case nop:
	}
	c.History = append(c.History, is.ID)

	return nextInstruction
}

func main() {

	fmt.Println("Advent of Code Day 10")
	if input, err := readInput("input.txt"); err == nil {

		fmt.Println("--- Part 1 ---")

		var jolts []int

		for _, inp := range input {
			i, _ := strconv.Atoi(inp)
			jolts = append(jolts, i)
		}

		sort.Ints(jolts)

		part1(jolts)

		fmt.Println("--- Part 2 ---")

		part2(jolts)

	}
}
func part2(jolts []int) {

	c := make(chan int)

	collectCount := 0
	currentJolt := 0

	// for each position in the list of jolts we can fork goroutines at i+1 and i+2 and if they reach the end they count as a possible combination
	for i, _ := range jolts {
		j1 := jolts[i]
		if j1 <= currentJolt+3 {
			collectCount++
			go crawl (currentJolt+j1,jolts[i:], &collectCount, c)
		}

		if i+1 < len(jolts) {
			j2 := jolts[i+1]
			if j2 <= currentJolt+3 {
				collectCount++
				go crawl (currentJolt+j2,jolts[i+1:], &collectCount, c)
			}
		}
		if i+2 < len(jolts) {
			j3 := jolts[i+2]
			if j3 <= currentJolt+3 {
				collectCount++
				go crawl (currentJolt+j3,jolts[i+2:], &collectCount, c)
			}
		}
	}

	time.Sleep(time.Millisecond * 100)
	validresults := 0

	for i := 0; i < collectCount; i++ {
		switch result := <- c; result {
		case 0:
			continue
		case 1:
			validresults++
		}

	}

	fmt.Println ("Valid cases %d", validresults)


}

func crawl(currentJolt int, jolts []int, collectCount *int, c chan int) {

	// for each position in the list of jolts we can fork goroutines at i+1 and i+2 and if they reach the end they count as a possible combination
	for i, _ := range jolts {
		j1 := jolts[i]

		following := false

		if j1 <= currentJolt+3 {
			*collectCount++
			following = true
			go crawl (currentJolt+j1,jolts[i:], collectCount, c)
		}

		if i+1 < len(jolts) {
			j2 := jolts[i+1]
			if j2 <= currentJolt+3 {
				*collectCount++
				following = true
				go crawl (currentJolt+j2,jolts[i+1:], collectCount, c)
			}
		}
		if i+2 < len(jolts) {
			j3 := jolts[i+2]
			if j3 <= currentJolt+3 {
				*collectCount++
				following = true
				go crawl (currentJolt+j3,jolts[i+2:], collectCount, c)
			}
		}
		if !following {
			c <- 0
		}
	}

	c <- 1

}

func part1(jolts []int) {
	found := map[int]int{}
	currentJolt := 0

	for _, jolt := range jolts {

		if jolt > currentJolt+3 {
			log.Fatalf("Current jolt is  %d found next jolt at %d.", currentJolt, jolt)
		}

		found[jolt-currentJolt] ++
		currentJolt = jolt
	}

	currentJolt += 3
	found[3] ++

	for k, v := range found {
		log.Printf("%d differences of %d jolt\n", v, k)
	}
}

func readInput(file string) ([]string, error) {
	if data, err := ioutil.ReadFile(file); err == nil {
		input := string(data)
		return strings.Split(input, "\n"), nil
	}
	return nil, errors.New("Could not parse file")
}
