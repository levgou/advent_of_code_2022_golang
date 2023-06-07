package day_10

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strconv"
	"strings"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_10/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_10/demo_input.txt"
)

const (
	NOOP = "noop"
	ADDX = "addx"
)

type Instruction struct {
	OP  string
	Arg int
}

func parseInstruction(line string) Instruction {
	if line == NOOP {
		return Instruction{OP: NOOP, Arg: 0}
	}
	opParts := strings.Split(line, " ")
	op := opParts[0]
	arg, err := strconv.Atoi(opParts[1])

	if err != nil {
		panic(err)
	}

	return Instruction{OP: op, Arg: arg}
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}

	tick := 0
	x := 1

	//  20th, 60th, 100th, 140th, 180th, and 220th
	cyclesOfInteres := []int{20, 60, 100, 140, 180, 220}

	sigStrength := 0
	pintX := func() {
		if slices.Index(cyclesOfInteres, tick) >= 0 {
			fmt.Println(fmt.Sprintf("tick: %d, x: %d", tick, x))
			sigStrength += x * tick
		}
	}

	for _, instruction := range instructions {
		if instruction.OP == NOOP {
			tick += 1
			pintX()
			continue
		}
		tick += 1
		pintX()

		tick += 1
		pintX()

		x += instruction.Arg
	}

	fmt.Println(sigStrength)

}

func Solution2() {

	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}

	x := 1
	instruction := 0
	adding := false

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			if x == col || x == col+1 || x == col-1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			op := instructions[instruction]

			if op.OP == ADDX {
				if adding {
					x += op.Arg
					adding = false
					instruction += 1
				} else {
					adding = true
				}
			} else {
				instruction += 1
			}
		}

		fmt.Println()
	}
}
