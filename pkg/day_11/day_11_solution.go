package day_11

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"strings"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_11/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_11/demo_input.txt"
)

type Monkey struct {
	Items   []int64
	Op      func(int64) int64
	Test    func(int64) bool
	Div     int64
	IfTrue  int64
	IfFalse int64
}

func parseOp(rhs string) func(int64) int64 {
	op := "+"
	if strings.Index(rhs, "*") >= 0 {
		op = "*"
	}

	operands := strings.Split(rhs, op)
	if strings.Index(operands[1], "old") >= 0 {
		return func(i int64) int64 {
			if op == "+" {
				return i + i
			} else {
				return i * i
			}
		}
	}

	numOperand := int64(shared.ParseInt(operands[1]))
	return func(i int64) int64 {
		if op == "+" {
			return i + numOperand
		} else {
			return i * numOperand
		}
	}
}

func parseMonkeys(lines []string) []Monkey {
	monkeys := make([]Monkey, 0)
	monkey := Monkey{}

	for _, line := range lines {
		if strings.Index(line, "Starting items") >= 0 {
			items := strings.Replace(line, "Starting items: ", "", 1)
			itemsSplit := strings.Split(items, ",")

			monkey.Items = make([]int64, len(itemsSplit))
			for i, item := range itemsSplit {
				monkey.Items[i] = int64(shared.ParseInt(strings.Trim(item, " ")))
			}
		} else if strings.Index(line, "Operation: new =") >= 0 {
			rhs := strings.Replace(line, "Operation: new =", "", 1)
			monkey.Op = parseOp(rhs)
		} else if strings.Index(line, "Test: divisible by") >= 0 {
			div := strings.Replace(line, "Test: divisible by", "", 1)
			divNum := int64(shared.ParseInt(strings.Trim(div, " ")))
			monkey.Test = func(i int64) bool {
				return i%divNum == 0
			}
			monkey.Div = divNum
		} else if strings.Index(line, " If true: throw to monkey") >= 0 {
			to := strings.Replace(line, " If true: throw to monkey", "", 1)
			toNum := shared.ParseInt(strings.Trim(to, " "))
			monkey.IfTrue = int64(toNum)
		} else if strings.Index(line, " If false: throw to monkey") >= 0 {
			to := strings.Replace(line, " If false: throw to monkey", "", 1)
			toNum := shared.ParseInt(strings.Trim(to, " "))
			monkey.IfFalse = int64(toNum)
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		}
	}

	return monkeys
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	monkeys := parseMonkeys(lines)
	fmt.Println(monkeys)

	itemsInspected := make(map[int]int64)
	for m := 0; m < len(monkeys); m++ {
		itemsInspected[m] = 0
	}

	coommonMod := int64(1)
	for m := 0; m < len(monkeys); m++ {
		coommonMod *= monkeys[m].Div
	}

	for i := 0; i < 10_000; i++ {
		//fmt.Println("Iteration", i)
		for m := 0; m < len(monkeys); m++ {
			//fmt.Println("Monkey", m)
			monkey := &monkeys[m]

			for _, item := range monkey.Items {
				itemsInspected[m]++

				worryLevel := monkey.Op(item)

				// Part1
				//worryLevel = worryLevel / 3
				// Part2:
				worryLevel %= coommonMod

				if monkey.Test(worryLevel) {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, worryLevel)
					//fmt.Println("Throwing", item, "to", monkey.IfTrue, "with worry level", worryLevel)
				} else {
					//fmt.Println("Throwing", item, "to", monkey.IfFalse, "with worry level", worryLevel)
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, worryLevel)
				}
			}

			monkey.Items = make([]int64, 0)
		}

		fmt.Println("\n")
		for m := 0; m < len(monkeys); m++ {
			fmt.Println("Monkey", m, "items:", monkeys[m].Items)
		}
	}

	fmt.Println(itemsInspected)
	max2, _ := shared.MaxN(maps.Values(itemsInspected), 2)
	fmt.Println(max2[0] * max2[1])
}
